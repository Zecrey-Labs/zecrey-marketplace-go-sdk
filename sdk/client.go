package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/zecrey-labs/zecrey-crypto/util/eddsaHelper"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
)

const (
	//nftMarketUrl = "http://localhost:9999"
	// nftMarketUrl = "https://test-legend-nft.zecrey.com"
	nftMarketUrl = "https://dev-legend-nft.zecrey.com"
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
	AccountName    string
	L2pk           string
	Seed           string
	NftMarketUrl   string
	LegendUrl      string
	ProviderClient *_rpc.ProviderClient
	KeyManager     KeyManager
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
		AccountName:    accountName,
		Seed:           seed,
		L2pk:           l2pk,
		NftMarketUrl:   nftMarketUrl,
		LegendUrl:      legendUrl,
		ProviderClient: connEth,
		KeyManager:     keyManager,
	}, nil
}

func (c *Client) SetKeyManager(keyManager KeyManager) {
	c.KeyManager = keyManager
}

type Colletcion struct {
	ShortName          string
	CategoryId         string
	CreatorEarningRate string
	CollectionUrl      string
	ExternalLink       string
	TwitterLink        string
	InstagramLink      string
	TelegramLink       string
	DiscordLink        string
	LogoImage          string
	FeaturedImage      string
	BannerImage        string
	PaymentAssetIds    string
	Description        string
}

func (c *Client) CreateCollection(collection Colletcion) (*RespCreateCollection, error) {
	respSdkTx, err := http.Get(c.NftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCreateCollectionTxInfo?account_name=%s", c.AccountName))
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
	tx, err := sdkCreateCollectionTxInfo(c.KeyManager, resultSdk.Transtion, collection.Description)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/collection/createCollection",
		url.Values{"short_name": {collection.ShortName},
			"category_id":          {collection.CategoryId},
			"collection_url":       {collection.CollectionUrl},
			"external_link":        {collection.ExternalLink},
			"twitter_link":         {collection.TwitterLink},
			"instagram_link":       {collection.TelegramLink},
			"discord_link":         {collection.InstagramLink},
			"telegram_link":        {collection.DiscordLink},
			"logo_image":           {collection.LogoImage},
			"featured_image":       {collection.FeaturedImage},
			"banner_image":         {collection.BannerImage},
			"creator_earning_rate": {collection.CreatorEarningRate},
			"payment_asset_ids":    {collection.PaymentAssetIds},
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
	signature := SignMessage(c.KeyManager, message)
	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/collection/updateCollection",
		url.Values{
			"id":             {Id},
			"account_name":   {c.AccountName},
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

type Mintnft struct {
	CollectionId int64
	NftUrl       string
	Name         string
	TreasuryRate int64
	Description  string
	Media        string
	Properties   string
	Levels       string
	Stats        string
}

func (c *Client) MintNft(nftInfo Mintnft) (*RespCreateAsset, error) {
	ContentHash, err := calculateContentHash(c.AccountName, nftInfo.CollectionId, nftInfo.Name, nftInfo.Properties, nftInfo.Levels, nftInfo.Stats)
	respSdkTx, err := http.Get(c.NftMarketUrl +
		fmt.Sprintf("/api/v1/sdk/getSdkMintNftTxInfo?account_name=%s&collection_id=%d&name=%s&content_hash=%streasury_rate%d",
			c.AccountName, nftInfo.CollectionId, nftInfo.Name, ContentHash, nftInfo.TreasuryRate))
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
	tx, err := sdkMintNftTxInfo(c.KeyManager, resultSdk.Transtion)
	if err != nil {
		return nil, err
	}

	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/asset/createAsset",
		url.Values{
			"collection_id": {fmt.Sprintf("%d", nftInfo.CollectionId)},
			"nft_url":       {nftInfo.NftUrl},
			"name":          {nftInfo.Name},
			"description":   {nftInfo.Description},
			"media":         {nftInfo.Media},
			"properties":    {nftInfo.Properties},
			"levels":        {nftInfo.Levels},
			"stats":         {nftInfo.Stats},
			"transaction":   {tx},
			"treasury_rate": {fmt.Sprintf("%d", nftInfo.TreasuryRate)},
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
	respSdkTx, err := http.Get(c.NftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkTransferNftTxInfo?account_name=%s&to_account_name=%s%s&nft_id=%d", c.AccountName, toAccountName, NameSuffix, AssetId))
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
	txInfo, err := sdkTransferNftTxInfo(c.KeyManager, resultSdk.Transtion)

	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/asset/sendTransferNft",
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
	respSdkTx, err := http.Get(c.NftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkWithdrawNftTxInfo?account_name=%s&nft_id=%d", c.AccountName, AssetId))
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

	txInfo, err := sdkWithdrawNftTxInfo(c.KeyManager, resultSdk.Transtion)
	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/asset/sendWithdrawNft",
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
	respSdkTx, err := http.Get(c.NftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkOfferTxInfo?account_name=%s&nft_id=%d&money_id=%d&money_amount=%d&is_sell=true", c.AccountName, AssetId, AssetType, AssetAmount))
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

	tx, err := sdkOfferTxInfo(c.KeyManager, resultSdk.Transtion, true)
	if err != nil {
		return nil, err
	}
	return c.Offer(c.AccountName, tx)
}

func (c *Client) CreateBuyOffer(AssetId int64, AssetType int64, AssetAmount *big.Int) (*RespListOffer, error) {
	respSdkTx, err := http.Get(c.NftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkOfferTxInfo?account_name=%s&nft_id=%d&money_id=%d&money_amount=%d&is_sell=false", c.AccountName, AssetId, AssetType, AssetAmount))
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
	tx, err := sdkOfferTxInfo(c.KeyManager, resultSdk.Transtion, false)
	if err != nil {
		return nil, err
	}
	return c.Offer(c.AccountName, tx)
}

func (c *Client) CancelOffer(offerId int64) (*RespCancelOffer, error) {
	respSdkTx, err := http.Get(c.NftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCancelOfferTxInfo?account_name=%s&offerId=%d", c.AccountName, offerId))
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
	tx, err := sdkOfferTxInfo(c.KeyManager, resultSdk.Transtion, false)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/offer/cancelOffer",
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
	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/offer/listOffer",
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
	respSdkTx, err := http.Get(c.NftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAtomicMatchWithTx?account_name=%s&offer_id=%d&money_id=%d&money_amount=%s&is_sell=%v", c.AccountName, offerId, 0, AssetAmount.String(), isSell))
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

	txInfo, err := sdkAtomicMatchWithTx(c.KeyManager, resultSdk.Transtion, isSell, AssetAmount)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(c.NftMarketUrl+"/api/v1/offer/acceptOffer",
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
	return c.AccountName, c.L2pk, c.Seed
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
