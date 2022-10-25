package sdk

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	ethchain "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	curve "github.com/zecrey-labs/zecrey-crypto/ecc/ztwistededwards/tebn254"
	"github.com/zecrey-labs/zecrey-crypto/util/ecdsaHelper"
	"github.com/zecrey-labs/zecrey-crypto/util/eddsaHelper"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	zecreyLegendRpc "github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey-legend"
	zecreyLegendUtil "github.com/zecrey-labs/zecrey-legend/common/util"
	"github.com/zeromicro/go-zero/core/logx"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	DefaultGasLimit = 5000000
	NameSuffix      = ".zec"
)

type client struct {
	nftMarketURL   string
	legendURL      string
	providerClient *_rpc.ProviderClient
	keyManager     KeyManager
}

func (c *client) SetKeyManager(keyManager KeyManager) {
	c.keyManager = keyManager
}

/*=========================== account =======================*/

func (c *client) CreateL1Account() (l1Addr, privateKeyStr, l2pk, seed string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logx.Errorf("[] GenerateKey err: %s", err)
		return "", "", "", "", err
	}
	privateKeyStr = hex.EncodeToString(crypto.FromECDSA(privateKey))
	l1Addr, err = ecdsaHelper.GenerateL1Address(privateKey)
	fmt.Println("l1Addr", l1Addr)
	if err != nil {
		logx.Errorf("[] GenerateL1Address err: %s", err)
		return "", "", "", "", err
	}
	seed, err = eddsaHelper.GetEddsaSeed(privateKey)
	if err != nil {
		logx.Errorf("[] GetEddsaSeed err: %s", err)
		return "", "", "", "", err
	}
	fmt.Println("seed", seed)
	l2pk = eddsaHelper.GetEddsaPublicKey(seed[2:])
	fmt.Println("pk", l2pk)
	return
}

func (c *client) RegisterAccountWithPrivateKey(accountName, l1Addr, l2pk, privateKey, ZecreyLegendContract, ZnsPriceOracle string) (txHash string, err error) {
	var chainId *big.Int
	chainId, err = c.providerClient.ChainID(context.Background())
	if err != nil {
		return "", err
	}
	authCli, err := _rpc.NewAuthClient(c.providerClient, privateKey, chainId)
	if err != nil {
		return "", err
	}
	px, py, err := zecreyLegendUtil.PubKeyStrToPxAndPy(l2pk)
	if err != nil {
		return "", err
	}

	gasPrice, err := c.providerClient.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return "", err
	}
	priceOracleInstance, err := zecreyLegendRpc.LoadStablePriceOracleInstance(c.providerClient, ZnsPriceOracle)
	if err != nil {
		return "", err
	}
	txHash, err = zecreyLegendRpc.RegisterZNS(c.providerClient, authCli,
		zecreyInstance, priceOracleInstance,
		gasPrice, DefaultGasLimit, accountName,
		common.HexToAddress(l1Addr), px, py)
	return txHash, err
}

func (c *client) GetAccountByAccountName(accountName, ZecreyLegendContract string) (address string, err error) {
	res, err := zecreyLegendUtil.ComputeAccountNameHashInBytes(accountName + NameSuffix)
	if err != nil {
		logx.Error(err)
		return "", err
	}
	resBytes := zecreyLegendUtil.SetFixed32Bytes(res)
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return "", err
	}
	// fetch by accountNameHash
	addr, err := zecreyInstance.GetAddressByAccountNameHash(zecreyLegendRpc.EmptyCallOpts(), resBytes)
	if err != nil {
		logx.Error(err)
		return "", err
	}
	return addr.String(), nil
}

func (c *client) ApplyRegisterHost(
	accountName string, l2Pk string, OwnerAddr string) (*RespApplyRegisterHost, error) {
	resp, err := http.PostForm(c.legendURL+"/api/v1/register/applyRegisterHost",
		url.Values{
			"account_name": {accountName},
			"l2_pk":        {l2Pk},
			"owner_addr":   {OwnerAddr}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespApplyRegisterHost{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

/*=========================== collection =======================*/

func (c *client) CreateCollection(
	accountName string, ShortName string, CategoryId string, CollectionUrl string,
	ExternalLink string, TwitterLink string, InstagramLink string, TelegramLink string, DiscordLink string, LogoImage string,
	FeaturedImage string, BannerImage string, CreatorEarningRate string, PaymentAssetIds string) (*RespCreateCollection, error) {

	respPrepareTx, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/preparetx/getPrepareCreateCollectionTxInfo?account_name=%s", accountName))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respPrepareTx.Body)
	if err != nil {
		return nil, err
	}
	if respPrepareTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	resultPrepare := &RespetPreparetxInfo{}
	if err := json.Unmarshal(body, &resultPrepare); err != nil {
		return nil, err
	}
	tx, err := PrepareCreateCollectionTxInfo(c.keyManager, resultPrepare.Transtion)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.nftMarketURL+"/api/v1/collection/createCollection",
		url.Values{"short_name": {ShortName},
			"category_id":          {CategoryId},
			"collection_url":       {CollectionUrl},
			"external_link":        {ExternalLink},
			"twitter_link":         {TwitterLink},
			"instagram_link":       {TelegramLink},
			"discord_link":         {InstagramLink},
			"telegram_link":        {DiscordLink},
			"logo_image":           {LogoImage},
			"featured_image":       {FeaturedImage},
			"banner_image":         {BannerImage},
			"creator_earning_rate": {CreatorEarningRate},
			"payment_asset_ids":    {PaymentAssetIds},
			"transaction":          {tx}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespCreateCollection{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// api/v1/action/actionGetCollectionById

func (c *client) GetCollectionById(collectionId int64) (*RespGetCollectionByCollectionId, error) {
	request_query := fmt.Sprintf("query MyQuery {\n  actionGetCollectionById(collection_id: %d) {\n    collection {\n      account_name\n      banner_thumb\n    }\n  }\n}\n", collectionId)
	input := InputCollectionByIdActionBody{CollectionId: collectionId}
	action := ActionBody{Name: "actionGetCollectionById"}
	SessionVariables := SessionVariablesBody{XHasuraUserId: "x-hasura-role", XHasuraRole: "admin"}
	req := ReqGetCollectionById{
		Input:            input,
		Action:           action,
		SessionVariables: SessionVariables,
		RequestQuery:     request_query,
	}
	statusJSON, _ := json.Marshal(req)
	resp, err := http.Post(c.nftMarketURL+"/api/v1/action/actionGetCollectionById", "application/json", bytes.NewReader(statusJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetCollectionByCollectionId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (c *client) UpdateCollection(
	Id string,
	AccountName string,
	Name string,
	CollectionUrl string,
	Description string,
	CategoryId string,
	ExternalLink string,
	TwitterLink string,
	InstagramLink string,
	TelegramLink string,
	DiscordLink string,
	LogoImage string,
	FeaturedImage string,
	BannerImage string,
) (*RespUpdateCollection, error) {
	timestamp := time.Now().Unix()
	message := fmt.Sprintf("%dupdate_collection", timestamp)
	signature := SignMessage(c.keyManager, message)
	resp, err := http.PostForm(c.nftMarketURL+"/api/v1/collection/updateCollection",
		url.Values{
			"id":             {Id},
			"account_name":   {AccountName},
			"name":           {Name},
			"collection_url": {CollectionUrl},
			"description":    {Description},
			"category_id":    {CategoryId},
			"external_link":  {ExternalLink},
			"twitter_link":   {TwitterLink},
			"instagram_link": {InstagramLink},
			"telegram_link":  {TelegramLink},
			"discord_link":   {DiscordLink},
			"logo_image":     {LogoImage},
			"featured_image": {FeaturedImage},
			"banner_image":   {BannerImage},
			"timestamp":      {fmt.Sprintf("%d", timestamp)},
			"signature":      {signature}},
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespUpdateCollection{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetCollectionsByAccountIndex(AccountIndex int64) (*RespGetAccountCollections, error) {
	request_query := fmt.Sprintf("query MyQuery {\n  actionGetAccountCollections(account_index: %d) {\n    confirmedCollectionIdList\n    pendingCollections {\n      account_name\n      banner_image\n      banner_thumb\n      browse_count\n      category_id\n      created_at\n      creator_earning_rate\n      description\n      discord_link\n      expired_at\n      external_link\n      featured_Thumb\n      featured_image\n      floor_price\n      id\n      instagram_link\n      item_count\n      l2_collection_id\n      logo_image\n      logo_thumb\n      name\n      one_day_trade_volume\n      short_name\n      status\n      telegram_link\n      total_trade_volume\n      twitter_link\n    }\n  }\n}", AccountIndex)
	input := InputGetAccountCollectionsActionBody{AccountIndex: AccountIndex}
	action := ActionBody{Name: "actionGetAccountCollections"}
	SessionVariables := SessionVariablesBody{XHasuraUserId: "x-hasura-role", XHasuraRole: "admin"}
	req := ReqGetAccountCollections{
		Input:            input,
		Action:           action,
		SessionVariables: SessionVariables,
		RequestQuery:     request_query,
	}
	statusJSON, _ := json.Marshal(req)
	resp, err := http.Post(c.nftMarketURL+"/api/v1/action/actionGetAccountCollections", "application/json", bytes.NewReader(statusJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountCollections{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

/*=========================== nft =======================*/

func (c *client) MintNft(
	accountName string,
	CollectionId int64,
	NftUrl string, Name string,
	Description string, Media string,
	Properties string, Levels string, Stats string,
) (*RespCreateAsset, error) {

	ContentHash, err := calculateContentHash(accountName, CollectionId, Name, Properties, Levels, Stats)

	respPrepareTx, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/preparetx/getPrepareMintNftTxInfo?account_name=%s&collection_id=%d&name=%s&content_hash=%s", accountName, CollectionId, Name, ContentHash))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respPrepareTx.Body)
	if err != nil {
		return nil, err
	}
	if respPrepareTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultPrepare := &RespetPreparetxInfo{}
	if err := json.Unmarshal(body, &resultPrepare); err != nil {
		return nil, err
	}
	tx, err := PrepareMintNftTxInfo(c.keyManager, resultPrepare.Transtion)
	if err != nil {
		return nil, err
	}

	resp, err := http.PostForm(c.nftMarketURL+"/api/v1/asset/createAsset",
		url.Values{
			"collection_id": {fmt.Sprintf("%d", CollectionId)},
			"nft_url":       {NftUrl},
			"name":          {Name},
			"description":   {Description},
			"media":         {Media},
			"properties":    {Properties},
			"levels":        {Levels},
			"stats":         {Stats},
			"transaction":   {tx},
		},
	)

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespCreateAsset{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (c *client) GetNftByNftId(nftId int64) (*RespetAssetByAssetId, error) {
	request_query := fmt.Sprintf("query MyQuery {\n  actionGetAssetByAssetId(asset_id: %d) {\n    asset {\n      account_name\n      audio_thumb\n      collection_id\n      content_hash\n      created_at\n      creator_earning_rate\n      description\n      expired_at\n      id\n      image_thumb\n      levels\n      media\n      name\n      nft_index\n      properties\n      stats\n      status\n      video_thumb\n    }\n  }\n}\n", nftId)
	input := InputGetAssetByIdActionBody{AssetId: nftId}
	action := ActionBody{Name: "actionGetAssetByAssetId"}
	SessionVariables := SessionVariablesBody{XHasuraUserId: "x-hasura-role", XHasuraRole: "admin"}
	req := ReqGetAssetById{
		Input:            input,
		Action:           action,
		SessionVariables: SessionVariables,
		RequestQuery:     request_query,
	}
	statusJSON, _ := json.Marshal(req)
	resp, err := http.Post(c.nftMarketURL+"/api/v1/action/actionGetAssetByAssetId", "application/json", bytes.NewReader(statusJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespetAssetByAssetId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil

}

func (c *client) TransferNft(
	AssetId int64,
	accountName string,
	toAccountName string) (*ResqSendTransferNft, error) {
	respPrepareTx, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/preparetx/getPrepareTransferNftTxInfo?account_name=%s&to_account_name=%s&nft_id=%d", accountName, toAccountName, AssetId))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respPrepareTx.Body)
	if err != nil {
		return nil, err
	}
	if respPrepareTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	resultPrepare := &RespetPreparetxInfo{}
	if err := json.Unmarshal(body, &resultPrepare); err != nil {
		return nil, err
	}
	txInfo, err := PrepareTransferNftTxInfo(c.keyManager, resultPrepare.Transtion)

	resp, err := http.PostForm(c.nftMarketURL+"/api/v1/asset/sendTransferNft",
		url.Values{
			"asset_id":    {fmt.Sprintf("%d", AssetId)},
			"transaction": {txInfo},
		},
	)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &ResqSendTransferNft{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (c *client) WithdrawNft(accountName string, AssetId int64) (*ResqSendWithdrawNft, error) {
	respPrepareTx, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/preparetx/getPrepareWithdrawNftTxInfo?account_name=%s&nft_id=%d", accountName, AssetId))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respPrepareTx.Body)
	if err != nil {
		return nil, err
	}
	if respPrepareTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultPrepare := &RespetPreparetxInfo{}
	if err := json.Unmarshal(body, &resultPrepare); err != nil {
		return nil, err
	}

	txInfo, err := PrepareWithdrawNftTxInfo(c.keyManager, resultPrepare.Transtion)
	resp, err := http.PostForm(c.nftMarketURL+"/api/v1/asset/sendWithdrawNft",
		url.Values{
			"asset_id":    {fmt.Sprintf("%d", AssetId)},
			"transaction": {txInfo},
		},
	)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &ResqSendWithdrawNft{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (c *client) SellNft(accountName string, AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error) {
	respPrepareTx, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/preparetx/getPrepareOfferTxInfo?account_name=%s&nft_id=%d&money_id=%d&money_amount=%d&is_sell=true", accountName, AssetId, moneyType, AssetAmount))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respPrepareTx.Body)
	if err != nil {
		return nil, err
	}
	if respPrepareTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultPrepare := &RespetPreparetxInfo{}
	if err := json.Unmarshal(body, &resultPrepare); err != nil {
		return nil, err
	}

	tx, err := PrepareOfferTxInfo(c.keyManager, resultPrepare.Transtion, true)
	if err != nil {
		return nil, err
	}
	return c.Offer(accountName, tx)
}
func (c *client) BuyNft(accountName string, AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error) {
	respPrepareTx, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/preparetx/getPrepareOfferTxInfo?account_name=%s&nft_id=%d&money_id=%d&money_amount=%d&is_sell=false", accountName, AssetId, moneyType, AssetAmount))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respPrepareTx.Body)
	if err != nil {
		return nil, err
	}
	if respPrepareTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultPrepare := &RespetPreparetxInfo{}
	if err := json.Unmarshal(body, &resultPrepare); err != nil {
		return nil, err
	}

	tx, err := PrepareOfferTxInfo(c.keyManager, resultPrepare.Transtion, false)
	if err != nil {
		return nil, err
	}
	return c.Offer(accountName, tx)
}
func (c *client) Offer(accountName string, tx string) (*RespListOffer, error) {
	resp, err := http.PostForm(c.nftMarketURL+"/api/v1/offer/listOffer",
		url.Values{
			"accountName": {accountName},
			"transaction": {tx},
		},
	)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespListOffer{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetNextOfferId(AccountName string) (*RespGetNextOfferId, error) {
	resp, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/offer/getNextOfferId?account_name=%s", AccountName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetNextOfferId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (c *client) GetOfferById(OfferId int64) (*RespGetOfferByOfferId, error) {
	resp, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/offer/getOfferByOfferId?offer_id=%d", OfferId))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetOfferByOfferId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) AcceptOffer(accountName string, offerId int64, isSell bool, AssetAmount *big.Int) (*RespAcceptOffer, error) {
	respPrepareTx, err := http.Get(c.nftMarketURL + fmt.Sprintf("/api/v1/preparetx/getPrepareAtomicMatchWithTx?account_name=%s&offer_id=%d&money_id=%d&money_amount=%s&is_sell=%v", accountName, offerId, 0, AssetAmount.String(), isSell))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respPrepareTx.Body)
	if err != nil {
		return nil, err
	}
	if respPrepareTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultPrepare := &RespetPreparetxInfo{}
	if err := json.Unmarshal(body, &resultPrepare); err != nil {
		return nil, err
	}

	txInfo, err := PrepareAtomicMatchWithTx(c.keyManager, resultPrepare.Transtion, isSell, AssetAmount)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.nftMarketURL+"/api/v1/offer/acceptOffer",
		url.Values{
			"id":          {fmt.Sprintf("%d", offerId)},
			"transaction": {txInfo},
		},
	)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespAcceptOffer{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// ================ nft tool ===
func PrepareCreateCollectionTxInfo(key KeyManager, txInfoPrepare string) (string, error) {
	txInfo := &CreateCollectionTxInfo{}
	err := json.Unmarshal([]byte(txInfoPrepare), txInfo)
	if err != nil {
		return "", err
	}
	tx, err := ConstructCreateCollectionTx(key, txInfo) //sign tx message
	if err != nil {
		return "", err
	}
	return tx, nil
}
func PrepareMintNftTxInfo(key KeyManager, txInfoPrepare string) (string, error) {
	txInfo := &MintNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoPrepare), txInfo)
	if err != nil {
		return "", err
	}
	tx, err := ConstructMintNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, nil
}

func PrepareTransferNftTxInfo(key KeyManager, txInfoPrepare string) (string, error) {
	txInfo := &TransferNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoPrepare), txInfo)
	if err != nil {
		return "", err
	}
	tx, err := ConstructTransferNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func PrepareAtomicMatchWithTx(key KeyManager, txInfoPrepare string, isSell bool, AssetAmount *big.Int) (string, error) {
	txInfo := &AtomicMatchTxInfo{}
	err := json.Unmarshal([]byte(txInfoPrepare), txInfo)
	if err != nil {
		return "", err
	}
	if !isSell {
		signedTx, err := ConstructOfferTx(key, txInfo.BuyOffer)
		if err != nil {
			return "", err
		}
		signedOffer, _ := ParseOfferTxInfo(signedTx)
		txInfo.BuyOffer = signedOffer
		txInfo.BuyOffer.AssetAmount = AssetAmount
	}
	if isSell {
		signedTx, err := ConstructOfferTx(key, txInfo.SellOffer)
		if err != nil {
			return "", err
		}
		signedOffer, _ := ParseOfferTxInfo(signedTx)
		txInfo.SellOffer = signedOffer
		txInfo.SellOffer.AssetAmount = AssetAmount
	}

	tx, err := ConstructAtomicMatchTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}
func PrepareWithdrawNftTxInfo(key KeyManager, txInfoPrepare string) (string, error) {
	txInfo := &WithdrawNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoPrepare), txInfo)
	if err != nil {
		return "", err
	}
	tx, err := ConstructWithdrawNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}
func PrepareOfferTxInfo(key KeyManager, txInfoPrepare string, isSell bool) (string, error) {
	txInfo := &OfferTxInfo{}
	err := json.Unmarshal([]byte(txInfoPrepare), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.Type = 0
	if isSell {
		txInfo.Type = 1
	}
	tx, err := ConstructOfferTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func accountNameHash(accountName string) (res string, err error) {
	// TODO Keccak256
	words := strings.Split(accountName, ".")
	if len(words) != 2 {
		return "", errors.New("[AccountNameHash] invalid account name")
	}
	rootNode := make([]byte, 32)
	hashOfBaseNode := keccakHash(append(rootNode, keccakHash([]byte(words[1]))...))

	baseNode := big.NewInt(0).Mod(big.NewInt(0).SetBytes(hashOfBaseNode), curve.Modulus)
	baseNodeBytes := make([]byte, 32)
	baseNode.FillBytes(baseNodeBytes)

	nameHash := keccakHash([]byte(words[0]))
	subNameHash := keccakHash(append(baseNodeBytes, nameHash...))

	subNode := big.NewInt(0).Mod(big.NewInt(0).SetBytes(subNameHash), curve.Modulus)
	subNodeBytes := make([]byte, 32)
	subNode.FillBytes(subNodeBytes)

	res = common.Bytes2Hex(subNodeBytes)
	return res, nil
}
func keccakHash(value []byte) []byte {
	hashVal := crypto.Keccak256Hash(value)
	return hashVal[:]
}
func calculateContentHash(accountName string, collectionId int64, name string, _properties string, _levels string, _stats string) (string, error) {

	var (
		properties []Propertie
		levels     []Level
		stats      []Stat
	)
	err := json.Unmarshal([]byte(_properties), &properties)
	if err != nil {
		return "", fmt.Errorf("json Unmarshal err properties%s", _properties)
	}
	err = json.Unmarshal([]byte(_levels), &levels)
	if err != nil {
		return "", fmt.Errorf("json Unmarshal err levels%s", _levels)
	}
	err = json.Unmarshal([]byte(_stats), &stats)
	if err != nil {
		return "", fmt.Errorf("json Unmarshal err stats%s", _stats)
	}

	content := fmt.Sprintf("ACC:%s CID:%d NFT:%s", accountName, collectionId, name)

	if len(properties) == 0 {
		content = content + " PROPERTIES: empty"
	} else {
		content = content + " PROPERTIES: "
		m := make(map[string]string, 0)
		keys := make([]string, 0)
		for _, k := range properties {
			m[k.Name] = k.Value
			keys = append(keys, k.Name)
		}
		sort.Strings(keys)
		for _, k := range keys {
			content = content + fmt.Sprintf("k:%s v:%s", k, m[k])
		}
	}

	if len(levels) == 0 {
		content = content + " LEVELS: empty"
	} else {
		content = content + " LEVELS: "
		m := make(map[string]int64, 0)
		keys := make([]string, 0)
		for _, k := range levels {
			m[k.Name] = k.Value
			keys = append(keys, k.Name)
		}
		sort.Strings(keys)
		for _, k := range keys {
			content = content + fmt.Sprintf("k:%s v:%d", k, m[k])
		}
	}

	if len(stats) == 0 {
		content = content + " STATS: empty"
	} else {
		content = content + " STATS: "
		m := make(map[string]int64, 0)
		keys := make([]string, 0)
		for _, k := range stats {
			m[k.Name] = k.Value
			keys = append(keys, k.Name)
		}
		sort.Strings(keys)
		for _, k := range keys {
			content = content + fmt.Sprintf("k:%s v:%d", k, m[k])
		}
	}

	hFunc := mimc.NewMiMC()
	hFunc.Write([]byte(content))
	bytes := crypto.Keccak256Hash([]byte(content))
	fmt.Println("==nft content ==", content)
	return common.Bytes2Hex(bytes[:]), nil
}
func SignMessage(key KeyManager, message string) string {
	fmt.Println("message: ", message)
	sig, err := key.Sign([]byte(message), mimc.NewMiMC())
	if err != nil {
		panic("failed to sign message, err: " + err.Error())
	}

	signed := hex.EncodeToString(sig[:])
	fmt.Println("signed:", signed)
	return signed
}
func PackInput(abi *abi.ABI, abiMethod string, params ...interface{}) []byte {
	input, err := abi.Pack(abiMethod, params...)
	if err != nil {
		log.Error(abiMethod, " error", err)
	}
	return input
}
func sendContractTransaction(client *ethclient.Client, from, toAddress common.Address, value *big.Int, privateKey *ecdsa.PrivateKey, input []byte) common.Hash {
	// Ensure a valid value field and resolve the account nonce
	logger := log.New("func", "sendContractTransaction")
	nonce, err := client.PendingNonceAt(context.Background(), from)
	if err != nil {
		logger.Error("PendingNonceAt", "error", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	//gasPrice = big.NewInt(1000 000 000 000)
	if err != nil {
		log.Error("SuggestGasPrice", "error", err)
	}
	gasLimit := uint64(DefaultGasLimit) // in units

	//If the contract surely has code (or code is not needed), estimate the transaction

	msg := ethchain.CallMsg{From: from, To: &toAddress, GasPrice: gasPrice, Value: value, Data: input, GasFeeCap: big.NewInt(3000000000000)}
	gasLimit, err = client.EstimateGas(context.Background(), msg)
	if err != nil {
		logger.Error("Contract exec failed", "error", err)
	}
	if gasLimit < 1 {
		//gasLimit = 866328
		gasLimit = 2100000
	}
	gasLimit = uint64(DefaultGasLimit)

	// Create the transaction, sign it and schedule it for execution
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, input)

	chainID, _ := client.ChainID(context.Background())
	logger.Info("TxInfo", "TX data nonce ", nonce, " gasLimit ", gasLimit, " gasPrice ", gasPrice, " chainID ", chainID)
	signer := types.LatestSignerForChainID(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		log.Error("SignTx", "error", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Error("SendTransaction", "error", err)
	}
	return signedTx.Hash()
}

func getResult(conn *ethclient.Client, txHash common.Hash, contract bool) {
	logger := log.New("func", "getResult")
	logger.Info("Please waiting ", " txHash ", txHash.String())
	i := 0
	for {
		time.Sleep(time.Millisecond * 200)
		i++
		_, isPending, err := conn.TransactionByHash(context.Background(), txHash)
		if err != nil {
			logger.Info("TransactionByHash", "error", err)
		}
		if !isPending {
			break
		}
		if i > 20 {
			break
		}
	}

	queryTx(conn, txHash, contract, false)
}
func queryTx(conn *ethclient.Client, txHash common.Hash, contract bool, pending bool) {
	logger := log.New("func", "queryTx")
	if pending {
		_, isPending, err := conn.TransactionByHash(context.Background(), txHash)
		if err != nil {
			logger.Error("TransactionByHash", "error", err)
		}
		if isPending {
			println("In tx_pool no validator  process this, please query later")
			os.Exit(0)
		}
	}

	receipt, err := conn.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		for {
			time.Sleep(time.Millisecond * 200)
			receipt, err = conn.TransactionReceipt(context.Background(), txHash)
			if err == nil {
				break
			}
		}
		logger.Error("TransactionReceipt", "error", err)
	}

	if receipt.Status == types.ReceiptStatusSuccessful {
		//block, err := conn.BlockByHash(context.Background(), receipt.BlockHash)
		//if err != nil {
		//	logger.Error("BlockByHash", err)
		//}
		//logger.Info("Transaction Success", " block Number", receipt.BlockNumber.Uint64(), " block txs", len(block.Transactions()), "blockhash", block.Hash().Hex())
		logger.Info("Transaction Success", "block Number", receipt.BlockNumber.Uint64())
	} else if receipt.Status == types.ReceiptStatusFailed {
		//isContinueError = false
		logger.Info("Transaction Failed ", "Block Number", receipt.BlockNumber.Uint64())
	}
}

/* ================ legend ========================= */

func (c *client) GetTxsByPubKey(accountPk string, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getTxsByPubKey?account_pk=%s&offset=%d&limit=%d",
			accountPk, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsByPubKey{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getTxsByAccountName?account_name=%s&offset=%d&limit=%d",
			accountName, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsByAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) GetTxsByAccountIndexAndTxType(accountIndex int64, txType, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getTxsByAccountIndexAndTxType?account_index=%d&tx_type=%d&offset=%d&limit=%d",
			accountIndex, txType, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsByAccountIndexAndTxType{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) GetTxsListByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getTxsListByAccountIndex?account_index=%d&offset=%d&limit=%d", accountIndex, offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsListByAccountIndex{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, nil, err
	}
	return result.Total, result.Txs, nil
}

func (c *client) Search(info string) (*RespSearch, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/info/search?info=%s", info))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespSearch{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccounts(offset, limit uint32) (*RespGetAccounts, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/info/getAccounts?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccounts{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetGasFeeAssetList() (*RespGetGasFeeAssetList, error) {
	resp, err := http.Get(c.legendURL + "/api/v1/info/getGasFeeAssetList")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetGasFeeAssetList{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (int64, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/info/getWithdrawGasFee?asset_id=%d&withdraw_asset_id=%d&withdraw_amount=%d",
			assetId, withdrawAssetId, withdrawAmount))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetGasFee{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.GasFee, nil
}

func (c *client) GetGasFee(assetId uint32) (int64, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/info/getGasFee?asset_id=%d", assetId))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetGasFee{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.GasFee, nil
}

func (c *client) GetCurrencyPrices(symbol string) (*RespGetCurrencyPrices, error) {
	resp, err := http.Get(c.legendURL + "/api/v1/info/getCurrencyPrices")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetCurrencyPrices{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetCurrencyPriceBySymbol(symbol string) (*RespGetCurrencyPriceBySymbol, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/info/getCurrencyPriceBySymbol?symbol=%s", symbol))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetCurrencyPriceBySymbol{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAssetsList() (*RespGetAssetsList, error) {
	resp, err := http.Get(c.legendURL + "/api/v1/info/getAssetsList")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAssetsList{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetLayer2BasicInfo() (*RespGetLayer2BasicInfo, error) {
	resp, err := http.Get(c.legendURL + "/api/v1/info/getLayer2BasicInfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetLayer2BasicInfo{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetBlockByCommitment(blockCommitment string) (*Block, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/block/getBlockByCommitment?block_commitment=%s", blockCommitment))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetBlockByCommitment{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return &result.Block, nil
}

func (c *client) GetBalanceByAssetIdAndAccountName(assetId uint32, accountName string) (string, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/account/getBalanceByAssetIdAndAccountName?asset_id=%d&account_name=%s", assetId, accountName))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	result := &RespGetBalanceInfoByAssetIdAndAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result.Balance, nil
}

func (c *client) GetAccountStatusByAccountName(accountName string) (*RespGetAccountStatusByAccountName, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/account/getAccountStatusByAccountName?account_name=%s", accountName))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountStatusByAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccountInfoByAccountIndex(accountIndex int64) (*RespGetAccountInfoByAccountIndex, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/account/getAccountInfoByAccountIndex?account_index=%d", accountIndex))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountInfoByAccountIndex{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccountInfoByPubKey(accountPk string) (*RespGetAccountInfoByPubKey, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/account/getAccountInfoByPubKey?account_pk=%s", accountPk))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountInfoByPubKey{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetAccountStatusByAccountPk(accountPk string) (*RespGetAccountStatusByAccountPk, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/account/getAccountStatusByAccountPk?account_pk=%s", accountPk))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetAccountStatusByAccountPk{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *client) GetTxByHash(txHash string) (*RespGetTxByHash, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getTxByHash?tx_hash=%s", txHash))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	txResp := &RespGetTxByHash{}
	if err := json.Unmarshal(body, &txResp); err != nil {
		return nil, err
	}
	return txResp, nil
}

func (c *client) GetMempoolTxs(offset, limit uint32) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getMempoolTxs?offset=%d&limit=%d", offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	txsResp := &RespGetMempoolTxs{}
	if err := json.Unmarshal(body, &txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.MempoolTxs, nil
}
func (c *client) GetMempoolStatusTxsPending() (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getMempoolTxsPending?offset=%d&limit=%d", 0, 0))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	txsResp := &RespGetMempoolTxs{}
	if err := json.Unmarshal(body, &txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.MempoolTxs, nil
}
func (c *client) GetmempoolTxsByAccountName(accountName string) (total uint32, txs []*Tx, err error) {
	resp, err := http.Get(c.legendURL + "/api/v1/tx/getmempoolTxsByAccountName?account_name=" + accountName)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	txsResp := &RespGetmempoolTxsByAccountName{}
	if err := json.Unmarshal(body, &txsResp); err != nil {
		return 0, nil, err
	}
	return txsResp.Total, txsResp.Txs, nil
}

func (c *client) GetAccountInfoByAccountName(accountName string) (*AccountInfo, error) {
	resp, err := http.Get(c.legendURL + "/api/v1/account/getAccountInfoByAccountName?account_name=" + accountName)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	account := &AccountInfo{}
	if err := json.Unmarshal(body, &account); err != nil {
		return nil, err
	}
	return account, nil
}

func (c *client) GetNextNonce(accountIdx int64) (int64, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getNextNonce?account_index=%d", accountIdx))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetNextNonce{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.Nonce, nil
}

func (c *client) GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/tx/getTxsListByBlockHeight?block_height=%d&limit=%d&offset=%d", blockHeight, 0, 0))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &RespGetTxsListByBlockHeight{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result.Txs, nil
}

func (c *client) GetMaxOfferId(accountIndex uint32) (uint64, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/nft/getMaxOfferId?account_index=%d", accountIndex))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetMaxOfferId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.OfferId, nil
}

func (c *client) GetBlockByBlockHeight(blockHeight int64) (*Block, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/block/getBlockByBlockHeight?block_height=%d", blockHeight))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	res := &RespGetBlockByBlockHeight{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	return res.Block, nil
}

func (c *client) GetBlocks(offset, limit int64) (uint32, []*Block, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/block/getBlocks?limit=%d&offset=%d", offset, limit))
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, nil, fmt.Errorf(string(body))
	}
	res := &RespGetBlocks{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, nil, err
	}
	return res.Total, res.Blocks, nil
}

func (c *client) GetCurrentBlockHeight() (int64, error) {
	resp, err := http.Get(c.legendURL +
		fmt.Sprintf("/api/v1/block/getCurrentBlockHeight"))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespCurrentBlockHeight{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.Height, nil
}

func (c *client) SendCreateCollectionTx(txInfo string) (int64, error) {
	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendCreateCollectionTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	res := &RespSendCreateCollectionTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, err
	}
	return res.CollectionId, nil
}

func (c *client) SendMintNftTx(txInfo string) (string, error) {
	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendMintNftTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &RespSendMintNftTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *client) SignAndSendMintNftTx(tx *MintNftTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructMintNftTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendMintNftTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &RespSendMintNftTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *client) SignAndSendCreateCollectionTx(tx *CreateCollectionTxInfo) (int64, error) {
	if c.keyManager == nil {
		return 0, fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructCreateCollectionTx(c.keyManager, tx)
	if err != nil {
		return 0, err
	}

	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendCreateCollectionTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	res := &RespSendCreateCollectionTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return 0, err
	}
	return res.CollectionId, nil
}

func (c *client) SendCancelOfferTx(txInfo string) (string, error) {
	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendCancelOfferTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &RespSendCancelOfferTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *client) SignAndSendCancelOfferTx(tx *CancelOfferTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructCancelOfferTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendCancelOfferTx(txInfo)
}

func (c *client) SendAtomicMatchTx(txInfo string) (string, error) {
	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendAtomicMatchTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &RespSendAtomicMatchTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *client) SignAndSendAtomicMatchTx(tx *AtomicMatchTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructAtomicMatchTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendAtomicMatchTx(txInfo)
}

func (c *client) SendTransferNftTx(txInfo string) (string, error) {
	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendTransferNftTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &RespSendTransferNftTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *client) SignAndSendTransferNftTx(tx *TransferNftTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructTransferNftTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendTransferNftTx(txInfo)
}

func (c *client) SendWithdrawNftTx(txInfo string) (string, error) {
	resp, err := http.PostForm(c.legendURL+"/api/v1/tx/sendWithdrawNftTx",
		url.Values{"tx_info": {txInfo}})
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(string(body))
	}
	res := &RespSendWithdrawNftTx{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}
	return res.TxId, nil
}

func (c *client) SignAndSendWithdrawNftTx(tx *WithdrawNftTxInfo) (string, error) {
	if c.keyManager == nil {
		return "", fmt.Errorf("key manager is nil")
	}

	txInfo, err := ConstructWithdrawNftTx(c.keyManager, tx)
	if err != nil {
		return "", err
	}

	return c.SendWithdrawNftTx(txInfo)
}
