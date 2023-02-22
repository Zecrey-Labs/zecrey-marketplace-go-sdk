package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/zecrey-labs/zecrey-crypto/util/eddsaHelper"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	zecreyLegendRpc "github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey-legend"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
)

const (
	//nftMarketUrl = "http://localhost:7777"

	nftMarketUrl = "https://test-legend-nft.zecrey.com"
	legendUrl    = "https://test-legend-app.zecrey.com"
	//hasuraUrl    = "https://legend-marketplace.hasura.app/v1/graphql"
	hasuraUrl = "https://hasura.zecrey.com/v1/graphql" //test
	//hasuraAdminKey = "j76XNG0u72QWBt4gS167wJlhnFNHSI5A6R1427KGJyMrFWI7s8wOvz1vmA4DsGos" //test
	hasuraAdminKey = "zecreyLegendTest@Hasura" //test

	//nftMarketUrl   = "https://dev-legend-nft.zecrey.com"
	//legendUrl      = "https://dev-legend-app.zecrey.com"
	//hasuraUrl      = "https://legend-market-dev.hasura.app/v1/graphql"
	//hasuraAdminKey = "kqWAsFWVvn61mFuiuQ5yqJkWpu5VS1B5FGTdFzlVlQJ9fMTr9yNIjOnN3hERC9ex" //dev

	//nftMarketUrl   = "https://qa-legend-nft.zecrey.com"
	//legendUrl      = "https://qa-legend-app.zecrey.com"
	//hasuraUrl      = "https://legend-market-qa.hasura.app/v1/graphql"
	//hasuraAdminKey = "M5tpo0dWWjYdW0erD0mHqwcRSObUowSprpS7Q3K33SNQ0dcXkPeL63tpoka9dTBw" //qa

	hasuraTimeDeadline = 15 //15s
	chainRpcUrl        = "https://data-seed-prebsc-1-s1.binance.org:8545"
	QueryNftUrl        = "https://deep-index.moralis.io"
	QueryNftUrlKey     = ""
	DefaultGasLimit    = 5000000
	NameSuffix         = ".zec"
	MinGasFee          = 100000000000000 // 0.0001BNB

	BNBAssetId = 0
	LEGAssetId = 1
	REYAssetId = 2
)

var (
	GlobalAssetId = REYAssetId
)

type Client struct {
	accountName    string
	l2pk           string
	seed           string
	nftMarketUrl   string
	legendUrl      string
	providerClient *_rpc.ProviderClient
	keyManager     KeyManager
}

func NewClient(accountName, seed string) (*Client, error) {
	keyManager, err := NewSeedKeyManager(seed)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong seed:%s", seed))
	}
	l2pk, err := eddsaHelper.GetEddsaCompressedPublicKey(seed)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong GetEddsaCompressedPublicKey :%s", err.Error()))
	}
	connEth, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	return &Client{
		accountName:    fmt.Sprintf("%s%s", accountName, NameSuffix),
		seed:           seed,
		l2pk:           l2pk,
		nftMarketUrl:   nftMarketUrl,
		legendUrl:      legendUrl,
		providerClient: connEth,
		keyManager:     keyManager,
	}, nil
}
func NewClientNoSuffix(accountName, seed string) (*Client, error) {
	keyManager, err := NewSeedKeyManager(seed)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong seed:%s", seed))
	}
	l2pk, err := eddsaHelper.GetEddsaCompressedPublicKey(seed)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong GetEddsaCompressedPublicKey :%s", err.Error()))
	}
	connEth, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	return &Client{
		accountName:    fmt.Sprintf("%s", accountName),
		seed:           seed,
		l2pk:           l2pk,
		nftMarketUrl:   nftMarketUrl,
		legendUrl:      legendUrl,
		providerClient: connEth,
		keyManager:     keyManager,
	}, nil
}
func (c *Client) SetKeyManager(keyManager KeyManager) {
	c.keyManager = keyManager
}

func (c *Client) CreateCollection(ShortName string, CategoryId string, CreatorEarningRate string, ops ...model.CollectionOption) (*RespCreateCollection, error) {
	cp := model.GetDefaultCollection()
	for _, do := range ops {
		do.F(cp)
	}

	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCreateCollectionTxInfo?account_name=%s", c.accountName))
	if err != nil {
		return nil, fmt.Errorf("sdk http get err:%s", err)
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, fmt.Errorf("sdk ioutil read err:%s", string(body))
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sdk statusCode %d err:%s", respSdkTx.StatusCode, string(body))
	}

	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, fmt.Errorf("json unmarshal 1 err:%s", err)
	}
	tx, err := sdkCreateCollectionTxInfo(c.keyManager, resultSdk.Transtion, cp.Description, ShortName)
	if err != nil {
		return nil, fmt.Errorf("sdkCreateCollectionTxInfo err:%s", err)
	}
	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/collection/createCollection",
		url.Values{
			"short_name":           {ShortName},
			"category_id":          {CategoryId},
			"collection_url":       {cp.CollectionUrl},
			"external_link":        {cp.ExternalLink},
			"twitter_link":         {cp.TwitterLink},
			"instagram_link":       {cp.TelegramLink},
			"discord_link":         {cp.InstagramLink},
			"telegram_link":        {cp.DiscordLink},
			"logo_image":           {cp.LogoImage},
			"featured_image":       {cp.FeaturedImage},
			"banner_image":         {cp.BannerImage},
			"creator_earning_rate": {CreatorEarningRate},
			"payment_asset_ids":    {cp.PaymentAssetIds},
			"transaction":          {tx}})
	if err != nil {
		return nil, fmt.Errorf("post err:%s", err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("resp.Body err:%s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status2Code %d err:%s", resp.StatusCode, string(body))
	}
	result := &RespCreateCollection{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("json unmarshal 2 err:%s", err)
	}
	return result, nil
}

func (c *Client) UpdateCollection(Id string, Name string, ops ...model.CollectionOption) (*RespUpdateCollection, error) {
	cp := &model.CollectionParams{}
	for _, do := range ops {
		do.F(cp)
	}
	CategoryId := "1"
	timestamp := time.Now().Unix()
	message := fmt.Sprintf("%dupdate_collection", timestamp)
	signature := SignMessage(c.keyManager, message)
	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/collection/updateCollection",
		url.Values{
			"id":             {Id},
			"account_name":   {c.accountName},
			"name":           {Name},
			"collection_url": {cp.CollectionUrl},
			"description":    {cp.Description},
			"category_id":    {CategoryId},
			"external_link":  {cp.ExternalLink},
			"twitter_link":   {cp.TwitterLink},
			"instagram_link": {cp.InstagramLink},
			"telegram_link":  {cp.TelegramLink},
			"discord_link":   {cp.DiscordLink},
			"logo_image":     {cp.LogoImage},
			"featured_image": {cp.FeaturedImage},
			"banner_image":   {cp.BannerImage},
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

func (c *Client) MintNft(CollectionId int64, NftUrl string, Name string, Description string, Media string, Properties string, Levels string, Stats string) (*RespCreateAsset, error) {

	ContentHash, err := calculateContentHash(c.accountName, CollectionId, Name, Properties, Levels, Stats)
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkMintNftTxInfo?treasury_rate=20&account_name=%s&collection_id=%d&name=%s&content_hash=%s", c.accountName, CollectionId, Name, ContentHash))
	//fmt.Println(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkMintNftTxInfo?account_name=%s&collection_id=%d&name=%s&content_hash=%s", c.accountName, CollectionInfo, Name, ContentHash))
	if err != nil {
		return nil, fmt.Errorf("sdk http get err:%s", err)
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, fmt.Errorf("sdk ioutil read err:%s", string(body))
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sdk statusCode %d err:%s", respSdkTx.StatusCode, string(body))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, fmt.Errorf("json unmarshal 1 err:%s", err)
	}
	tx, err := sdkMintNftTxInfo(c.keyManager, resultSdk.Transtion)
	if err != nil {
		return nil, fmt.Errorf("sdkMintNftTxInfo err:%s", err)
	}

	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/asset/createAsset",
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
		return nil, fmt.Errorf("resp.Body err:%s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status2Code %d err:%s", resp.StatusCode, string(body))
	}
	result := &RespCreateAsset{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) TransferNft(
	AssetId int64,
	toAccountName string) (*RespSendTransferNft, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkTransferNftTxInfo?account_name=%s&to_account_name=%s%s&nft_id=%d", c.accountName, toAccountName, NameSuffix, AssetId))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, err
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, err
	}
	txInfo, err := sdkTransferNftTxInfo(c.keyManager, resultSdk.Transtion)

	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/asset/sendTransferNft",
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
	result := &RespSendTransferNft{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) WithdrawNft(AssetId int64, tol1Address string) (*RespSendWithdrawNft, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkWithdrawNftTxInfo?account_name=%s&nft_id=%d", c.accountName, AssetId))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, err
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, err
	}

	txInfo, err := sdkWithdrawNftTxInfo(c.keyManager, resultSdk.Transtion, tol1Address)
	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/asset/sendWithdrawNft",
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
	result := &RespSendWithdrawNft{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Withdraw(tol1Address string, assetId, assetAmount int64) (*RespSendWithdrawTx, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkWithdrawTxInfo?account_name=%s", c.accountName))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, err
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, err
	}

	txInfo, err := sdkWithdrawTxInfo(c.keyManager, resultSdk.Transtion, tol1Address, assetAmount, assetId)
	resp, err := http.PostForm(c.legendUrl+"/api/v1/tx/sendWithdrawTx",
		url.Values{
			"tx_info": {txInfo},
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
	result := &RespSendWithdrawTx{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateSellOffer(AssetId int64, AssetType int64, AssetAmount *big.Int) (*RespListOffer, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkOfferTxInfo?account_name=%s&nft_id=%d&money_id=%d&money_amount=%d&is_sell=true", c.accountName, AssetId, AssetType, AssetAmount))
	if err != nil {
		return nil, fmt.Errorf("sdk get %s", err.Error())
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, fmt.Errorf("sdk readbody %s", err.Error())
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sdk code%d %s", respSdkTx.StatusCode, string(body))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, fmt.Errorf("sdk resultSdk %s", err.Error())
	}

	tx, err := sdkOfferTxInfo(c.keyManager, resultSdk.Transtion, AssetAmount, true)
	if err != nil {
		return nil, fmt.Errorf("sdk convert  %s", err.Error())
	}
	return c.Offer(c.accountName, tx)
}

func (c *Client) CreateBuyOffer(AssetId int64, AssetType int64, AssetAmount *big.Int) (*RespListOffer, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkOfferTxInfo?account_name=%s&nft_id=%d&money_id=%d&money_amount=%d&is_sell=false", c.accountName, AssetId, AssetType, AssetAmount))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, err
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, err
	}
	tx, err := sdkOfferTxInfo(c.keyManager, resultSdk.Transtion, AssetAmount, false)
	if err != nil {
		return nil, err
	}
	return c.Offer(c.accountName, tx)
}

func (c *Client) CancelOffer(offerId int64) (*RespCancelOffer, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCancelOfferTxInfo?account_name=%s&offer_id=%d", c.accountName, offerId))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, err
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("sdk status2Code %d err:%s", respSdkTx.StatusCode, string(body))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, err
	}
	tx, err := sdkCancelOfferTxInfo(c.keyManager, resultSdk.Transtion)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/offer/cancelOffer",
		url.Values{
			"id":          {fmt.Sprintf("%d", offerId)},
			"transaction": {tx},
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
		return nil, fmt.Errorf("status2Code %d err:%s", resp.StatusCode, string(body))
	}
	result := &RespCancelOffer{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) Offer(accountName string, tx string) (*RespListOffer, error) {
	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/offer/listOffer",
		url.Values{
			"accountName": {accountName},
			"transaction": {tx},
		},
	)
	if err != nil {
		return nil, fmt.Errorf(" postform  %s", err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(" io read all  %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("post code%d %s", resp.StatusCode, string(body))
	}
	result := &RespListOffer{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("json unmarshal %s", err.Error())
	}
	return result, nil
}

func (c *Client) AcceptOffer(offerId int64, isSell bool, assetAmount *big.Int) (*RespAcceptOffer, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAtomicMatchWithTx?account_name=%s&offer_id=%d&money_id=%d&money_amount=%s&is_sell=%v", c.accountName, offerId, 0, assetAmount.String(), isSell))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, err
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, err
	}

	txInfo, err := sdkAtomicMatchWithTx(c.keyManager, resultSdk.Transtion, isSell, assetAmount)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/offer/acceptOffer",
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

func (c *Client) Deposit(accountName, privateKey string, assetId, assetAmount int64, BEP20token common.Address) (*types.Transaction, error) {
	accountName = fmt.Sprintf("%s%s", accountName, NameSuffix)
	var chainId *big.Int
	chainId, err := c.providerClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	authCli, err := _rpc.NewAuthClient(c.providerClient, privateKey, chainId)
	if err != nil {
		return nil, err
	}
	//get base contract address
	resp, err := GetLayer2BasicInfo()
	if err != nil {
		return nil, err
	}
	ZecreyLegendContract := resp.ContractAddresses[0]

	gasPrice, err := c.providerClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return nil, err
	}

	transactOpts, err := zecreyLegendRpc.ConstructTransactOptsWithValue(c.providerClient, authCli, gasPrice, DefaultGasLimit, assetAmount)
	if err != nil {
		return nil, err
	}
	var depositTransaction *types.Transaction
	if assetId == BNBAssetId {
		depositTransaction, err = zecreyInstance.DepositBNB(transactOpts, accountName)
		if err != nil {
			return nil, err
		}
	}
	if assetId != BNBAssetId {
		depositTransaction, err = zecreyInstance.DepositBEP20(transactOpts, BEP20token, big.NewInt(assetAmount), accountName)
		if err != nil {
			return nil, err
		}
	}

	return depositTransaction, nil
}
func (c *Client) DepositNft(accountName, privateKey string, _nftL1Address common.Address, _nftL1TokenId *big.Int) (*types.Transaction, error) {
	var chainId *big.Int
	chainId, err := c.providerClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	authCli, err := _rpc.NewAuthClient(c.providerClient, privateKey, chainId)
	if err != nil {
		return nil, err
	}
	//get base contract address
	resp, err := GetLayer2BasicInfo()
	if err != nil {
		return nil, err
	}
	ZecreyLegendContract := resp.ContractAddresses[0]
	ZnsPriceOracle := resp.ContractAddresses[1]

	gasPrice, err := c.providerClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return nil, err
	}

	priceOracleInstance, err := zecreyLegendRpc.LoadStablePriceOracleInstance(c.providerClient, ZnsPriceOracle)
	if err != nil {
		return nil, err
	}

	amount, err := zecreyLegendRpc.Price(priceOracleInstance, accountName)
	if err != nil {
		return nil, err
	}
	transactOpts, err := zecreyLegendRpc.ConstructTransactOptsWithValue(c.providerClient, authCli, gasPrice, DefaultGasLimit, amount.Int64())
	if err != nil {
		return nil, err
	}
	depositNftTransaction, err := zecreyInstance.DepositNft(transactOpts, accountName, _nftL1Address, _nftL1TokenId)
	if err != nil {
		return nil, err
	}
	return depositNftTransaction, nil
}
func (c *Client) FullExit(accountName, privateKey string, assetAddress common.Address, amount int64) (*types.Transaction, error) {
	accountName = fmt.Sprintf("%s%s", accountName, NameSuffix)
	var chainId *big.Int
	chainId, err := c.providerClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	authCli, err := _rpc.NewAuthClient(c.providerClient, privateKey, chainId)
	if err != nil {
		return nil, err
	}
	//get base contract address
	resp, err := GetLayer2BasicInfo()
	if err != nil {
		return nil, err
	}
	ZecreyLegendContract := resp.ContractAddresses[0]
	//ZnsPriceOracle := resp.ContractAddresses[1]

	gasPrice, err := c.providerClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return nil, err
	}

	//priceOracleInstance, err := zecreyLegendRpc.LoadStablePriceOracleInstance(c.providerClient, ZnsPriceOracle)
	//if err != nil {
	//	return nil, err
	//}

	//amount, err := zecreyLegendRpc.Price(priceOracleInstance, accountName)
	//if err != nil {
	//	return nil, err
	//}
	transactOpts, err := zecreyLegendRpc.ConstructTransactOptsWithValue(c.providerClient, authCli, gasPrice, DefaultGasLimit, amount)
	if err != nil {
		return nil, err
	}
	fullExitTransaction, err := zecreyInstance.RequestFullExit(transactOpts, accountName, assetAddress)
	if err != nil {
		return nil, err
	}
	return fullExitTransaction, nil
}
func (c *Client) FullExitNft(accountName, privateKey string, _nftIndex uint32) (*types.Transaction, error) {
	//accountName = fmt.Sprintf("%s%s", accountName, NameSuffix)
	var chainId *big.Int
	chainId, err := c.providerClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	authCli, err := _rpc.NewAuthClient(c.providerClient, privateKey, chainId)
	if err != nil {
		return nil, err
	}
	//get base contract address
	resp, err := GetLayer2BasicInfo()
	if err != nil {
		return nil, err
	}
	ZecreyLegendContract := resp.ContractAddresses[0]
	ZnsPriceOracle := resp.ContractAddresses[1]

	gasPrice, err := c.providerClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return nil, err
	}

	priceOracleInstance, err := zecreyLegendRpc.LoadStablePriceOracleInstance(c.providerClient, ZnsPriceOracle)
	if err != nil {
		return nil, err
	}

	amount, err := zecreyLegendRpc.Price(priceOracleInstance, accountName)
	if err != nil {
		return nil, err
	}
	transactOpts, err := zecreyLegendRpc.ConstructTransactOptsWithValue(c.providerClient, authCli, gasPrice, DefaultGasLimit, amount.Int64())
	if err != nil {
		return nil, err
	}

	fullExitNftTransaction, err := zecreyInstance.RequestFullExitNft(transactOpts, accountName, _nftIndex)
	if err != nil {
		return nil, err
	}
	return fullExitNftTransaction, nil
}

/*
GetMyInfo accountName、l2pk、seed
*/
func (c *Client) GetMyInfo() (accountName string, l2pk string, seed string) {
	return c.accountName, c.l2pk, c.seed
}
func (c *Client) GetKeyManager() KeyManager {
	return c.keyManager
}

func sdkCreateCollectionTxInfo(key KeyManager, txInfoSdk, Description, ShortName string) (string, error) {
	txInfo := &CreateCollectionTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	//reset
	txInfo.GasFeeAssetId = int64(GlobalAssetId)
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	txInfo.Introduction = Description
	txInfo.Name = ShortName
	tx, err := constructCreateCollectionTx(key, txInfo) //sign tx message
	if err != nil {
		return "", err
	}
	return tx, nil
}

func sdkMintNftTxInfo(key KeyManager, txInfoSdk string) (string, error) {
	txInfo := &MintNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.GasFeeAssetId = int64(GlobalAssetId)
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	tx, err := constructMintNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, nil
}

func sdkTransferNftTxInfo(key KeyManager, txInfoSdk string) (string, error) {
	txInfo := &TransferNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.GasFeeAssetId = int64(GlobalAssetId)
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	tx, err := constructTransferNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func sdkAtomicMatchWithTx(key KeyManager, txInfoSdk string, isSell bool, AssetAmount *big.Int) (string, error) {
	txInfo := &AtomicMatchTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	if !isSell {
		txInfo.BuyOffer.AssetId = int64(GlobalAssetId)
		txInfo.BuyOffer.AssetAmount = AssetAmount
		signedTx, err := constructOfferTx(key, txInfo.BuyOffer)
		if err != nil {
			return "", err
		}
		signedOffer, _ := parseOfferTxInfo(signedTx)
		txInfo.BuyOffer = signedOffer

	}
	if isSell {
		txInfo.SellOffer.AssetId = int64(GlobalAssetId)
		txInfo.SellOffer.AssetAmount = AssetAmount
		signedTx, err := constructOfferTx(key, txInfo.SellOffer)
		if err != nil {
			return "", err
		}
		signedOffer, _ := parseOfferTxInfo(signedTx)
		txInfo.SellOffer = signedOffer

	}
	txInfo.GasFeeAssetId = int64(GlobalAssetId)
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	tx, err := constructAtomicMatchTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func sdkWithdrawNftTxInfo(key KeyManager, txInfoSdk string, tol1Address string) (string, error) {
	txInfo := &WithdrawNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.GasFeeAssetId = int64(GlobalAssetId)
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	txInfo.ToAddress = tol1Address
	tx, err := constructWithdrawNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}
func sdkWithdrawTxInfo(key KeyManager, txInfoSdk string, tol1Address string, assetAmount, assetId int64) (string, error) {
	txInfo := &WithdrawTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.GasFeeAssetId = int64(GlobalAssetId)
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	txInfo.AssetId = assetId
	txInfo.AssetAmount = big.NewInt(assetAmount)
	txInfo.ToAddress = tol1Address
	tx, err := constructWithdrawTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func sdkOfferTxInfo(key KeyManager, txInfoSdk string, AssetAmount *big.Int, isSell bool) (string, error) {
	txInfo := &OfferTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.Type = 0
	if isSell {
		txInfo.Type = 1
	}
	txInfo.AssetId = int64(GlobalAssetId)
	txInfo.AssetAmount = AssetAmount
	tx, err := constructOfferTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func sdkCancelOfferTxInfo(key KeyManager, tx string) (string, error) {
	txInfo := &CancelOfferTxInfo{}
	err := json.Unmarshal([]byte(tx), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.GasFeeAssetId = int64(GlobalAssetId)
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	convertedTx := ConvertCancelOfferTxInfo(txInfo)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeCancelOfferMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
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
	return common.Bytes2Hex(bytes[:]), nil
}
