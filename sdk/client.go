package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/zecrey-labs/zecrey-crypto/util/eddsaHelper"
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
	//nftMarketUrl = "http://localhost:9999"
	nftMarketUrl = "https://test-legend-nft.zecrey.com"
	//nftMarketUrl = "https://dev-legend-nft.zecrey.com"
	//nftMarketUrl = "https://qa-legend-nft.zecrey.com"

	legendUrl = "https://qa-legend-app.zecrey.com"
	//legendUrl = "https://dev-legend-app.zecrey.com"
	//legendUrl    = "https://test-legend-app.zecrey.com"

	//hasuraUrl          = "https://legend-market-dev.hasura.app/v1/graphql"
	hasuraUrl = "https://legend-market-qa.hasura.app/v1/graphql"
	//hasuraUrl          = "https://legend-marketplace.hasura.app/v1/graphql"
	hasuraAdminKey = "M5tpo0dWWjYdW0erD0mHqwcRSObUowSprpS7Q3K33SNQ0dcXkPeL63tpoka9dTBw" //qa
	//hasuraAdminKey     = "j76XNG0u72QWBt4gS167wJlhnFNHSI5A6R1427KGJyMrFWI7s8wOvz1vmA4DsGos"//test
	hasuraTimeDeadline = 15 //15s

	chainRpcUrl     = "https://data-seed-prebsc-1-s1.binance.org:8545"
	DefaultGasLimit = 5000000
	NameSuffix      = ".zec"
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
	l2pk := eddsaHelper.GetEddsaPublicKey(seed[2:])
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

func (c *Client) SetKeyManager(keyManager KeyManager) {
	c.keyManager = keyManager
}

func (c *Client) CreateCollection(ShortName string, CategoryId string, CreatorEarningRate string, ops ...model.CollectionOption) (*RespCreateCollection, error) {
	cp := &model.CollectionParams{}
	for _, do := range ops {
		do.F(cp)
	}

	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCreateCollectionTxInfo?account_name=%s", c.accountName))
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
	tx, err := sdkCreateCollectionTxInfo(c.keyManager, resultSdk.Transtion, cp.Description)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/collection/createCollection",
		url.Values{"short_name": {ShortName},
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

	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkMintNftTxInfo?account_name=%s&collection_id=%d&name=%s&content_hash=%s", c.accountName, CollectionId, Name, ContentHash))
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
	tx, err := sdkMintNftTxInfo(c.keyManager, resultSdk.Transtion)
	if err != nil {
		return nil, err
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

func (c *Client) TransferNft(
	AssetId int64,
	toAccountName string) (*ResqSendTransferNft, error) {
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
	result := &ResqSendTransferNft{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) WithdrawNft(AssetId int64) (*ResqSendWithdrawNft, error) {
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

	txInfo, err := sdkWithdrawNftTxInfo(c.keyManager, resultSdk.Transtion)
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
	result := &ResqSendWithdrawNft{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Client) CreateSellOffer(AssetId int64, AssetType int64, AssetAmount *big.Int) (*RespListOffer, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkOfferTxInfo?account_name=%s&nft_id=%d&money_id=%d&money_amount=%d&is_sell=true", c.accountName, AssetId, AssetType, AssetAmount))
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

	tx, err := sdkOfferTxInfo(c.keyManager, resultSdk.Transtion, true)
	if err != nil {
		return nil, err
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
	tx, err := sdkOfferTxInfo(c.keyManager, resultSdk.Transtion, false)
	if err != nil {
		return nil, err
	}
	return c.Offer(c.accountName, tx)
}

func (c *Client) CancelOffer(offerId int64) (*RespCancelOffer, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCancelOfferTxInfo?account_name=%s&offerId=%d", c.accountName, offerId))
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
	tx, err := sdkOfferTxInfo(c.keyManager, resultSdk.Transtion, false)
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
		return nil, fmt.Errorf(string(body))
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

func (c *Client) AcceptOffer(offerId int64, isSell bool, AssetAmount *big.Int) (*RespAcceptOffer, error) {
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAtomicMatchWithTx?account_name=%s&offer_id=%d&money_id=%d&money_amount=%s&is_sell=%v", c.accountName, offerId, 0, AssetAmount.String(), isSell))
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

	txInfo, err := sdkAtomicMatchWithTx(c.keyManager, resultSdk.Transtion, isSell, AssetAmount)
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

/*
GetMyInfo accountName、l2pk、seed
*/
func (c *Client) GetMyInfo() (accountName string, l2pk string, seed string) {
	return c.accountName, c.l2pk, c.seed
}

func sdkCreateCollectionTxInfo(key KeyManager, txInfoSdk, Description string) (string, error) {
	txInfo := &CreateCollectionTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	//reset
	txInfo.GasFeeAssetAmount = big.NewInt(1000000000000000)
	txInfo.Introduction = Description
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
	txInfo.GasFeeAssetAmount = big.NewInt(1000000000000000)
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
	txInfo.GasFeeAssetAmount = big.NewInt(1000000000000000)
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
		signedTx, err := constructOfferTx(key, txInfo.BuyOffer)
		if err != nil {
			return "", err
		}
		signedOffer, _ := parseOfferTxInfo(signedTx)
		txInfo.BuyOffer = signedOffer
		txInfo.BuyOffer.AssetAmount = AssetAmount
	}
	if isSell {
		signedTx, err := constructOfferTx(key, txInfo.SellOffer)
		if err != nil {
			return "", err
		}
		signedOffer, _ := parseOfferTxInfo(signedTx)
		txInfo.SellOffer = signedOffer
		txInfo.SellOffer.AssetAmount = AssetAmount
	}

	tx, err := constructAtomicMatchTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func sdkWithdrawNftTxInfo(key KeyManager, txInfoSdk string) (string, error) {
	txInfo := &WithdrawNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.GasFeeAssetAmount = big.NewInt(1000000000000000)
	tx, err := constructWithdrawNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}

func sdkOfferTxInfo(key KeyManager, txInfoSdk string, isSell bool) (string, error) {
	txInfo := &OfferTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.Type = 0
	if isSell {
		txInfo.Type = 1
	}
	tx, err := constructOfferTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
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
