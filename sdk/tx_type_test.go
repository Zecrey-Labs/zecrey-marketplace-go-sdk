package sdk

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	//nftMarketUrl string = "http://34.111.87.92/"
	//legendUrl    string = "https://dev-legend-app.zecrey.com"
	nftMarketUrl string = "http://localhost:9999"
	legendUrl    string = "https://test-legend-app.zecrey.com"
)

func TestGenerateAddLiquidity(t *testing.T) {
	var a = AddLiquidityTxInfo{
		FromAccountIndex:  0,
		PairIndex:         0,
		AssetAId:          0,
		AssetAAmount:      big.NewInt(10000),
		AssetBId:          0,
		AssetBAmount:      big.NewInt(100),
		LpAmount:          big.NewInt(995),
		KLast:             big.NewInt(50000),
		TreasuryAmount:    big.NewInt(3),
		GasAccountIndex:   0,
		GasFeeAssetId:     0,
		GasFeeAssetAmount: big.NewInt(200),
		ExpiredAt:         1654656781000,
		Nonce:             1,
		Sig:               []byte("QgkTDbEq3Pq7AjidooPyfHmlSa1VuBAgqv57XjOT7yQC6OzNBv6YQLSm6U1BmPKA/qzFhfpnVFR8jL64kX/W+g=="),
	}

	aBytes, err := json.Marshal(a)
	assert.Nil(t, err)
	log.Println(string(aBytes))
}

func TestParseAddLiquidityTxInfo(t *testing.T) {
	txInfo := "{\"FromAccountIndex\":0,\"PairIndex\":0,\"AssetAId\":0,\"AssetAAmount\":10000,\"AssetBId\":0,\"AssetBAmount\":100,\"LpAmount\":995,\"KLast\":50000,\"TreasuryAmount\":3,\"GasAccountIndex\":0,\"GasFeeAssetId\":0,\"GasFeeAssetAmount\":200,\"ExpiredAt\":1654656781000,\"Nonce\":1,\"Sig\":\"UWdrVERiRXEzUHE3QWppZG9vUHlmSG1sU2ExVnVCQWdxdjU3WGpPVDd5UUM2T3pOQnY2WVFMU202VTFCbVBLQS9xekZoZnBuVkZSOGpMNjRrWC9XK2c9PQ==\"}"
	var addLiquidityTx *AddLiquidityTxInfo

	err := json.Unmarshal([]byte(txInfo), &addLiquidityTx)
	assert.Nil(t, err)
	log.Println(addLiquidityTx)
}

func TestParseCreateCollectionTxInfo(t *testing.T) {
	accountName := "sher.zec"
	accountSeed := "28e1a3762f....."
	ShortName := fmt.Sprintf("collection_Shortname %d", time.Now().Second())
	CategoryId := "1"
	CollectionUrl := "-"
	ExternalLink := "-"
	TwitterLink := "-"
	InstagramLink := "-"
	TelegramLink := "-"
	DiscordLink := "-"
	LogoImage := "collection/cbenqstwzx5uy9oedjrb"
	FeaturedImage := "collection/cbenqstwzx5uy9oedjrb"
	BannerImage := "collection/cbenqstwzx5uy9oedjrb"
	CreatorEarningRate := "6666"
	PaymentAssetIds := "[]"
	keyManager, err := NewSeedKeyManager(accountSeed)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret, err := c.CreateCollection(accountName, ShortName, CategoryId, CollectionUrl,
		ExternalLink, TwitterLink, InstagramLink, TelegramLink, DiscordLink, LogoImage,
		FeaturedImage, BannerImage, CreatorEarningRate, PaymentAssetIds)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret)
	fmt.Println("CreateCollection:", string(data))
	ret2, err := c.GetCollectionById(ret.Collection.Id)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(ret2)
	fmt.Println("GetCollectionById:", string(data))
}
func TestGetCollectionById(t *testing.T) {
	var collectionId int64 = 54
	accountSeed := "28e1a3762f....."
	keyManager, err := NewSeedKeyManager(accountSeed)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret2, err := c.GetCollectionById(collectionId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetCollectionById:", string(data))
}

func TestGetCollectionByAccountIndex(t *testing.T) {
	var accountIndex int64 = 2
	accountSeed := "28e1a3762f....."
	keyManager, err := NewSeedKeyManager(accountSeed)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret2, err := c.GetCollectionsByAccountIndex(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetCollectionsByAccountIndex:", string(data))
}

func TestUpdateCollection(t *testing.T) {
	Id := "54"
	AccountName := "sher.zec"
	Name := "zw-sdk--collection-update"
	CollectionUrl := "-"
	Description := "-"
	CategoryId := "1"
	ExternalLink := "-"
	TwitterLink := "-"
	InstagramLink := "-"
	TelegramLink := "-"
	DiscordLink := "-"
	LogoImage := "collection/cbenqstwzx5uy9oedjrb"
	FeaturedImage := "collection/cbenqstwzx5uy9oedjrb"
	BannerImage := "collection/cbenqstwzx5uy9oedjrb"

	var AccountIndex int64 = 2
	accountSeed := "28e1a3762f....."
	keyManager, err := NewSeedKeyManager(accountSeed)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret, err := c.UpdateCollection(Id, AccountName,
		Name, CollectionUrl, Description, CategoryId,
		ExternalLink, TwitterLink, InstagramLink, TelegramLink,
		DiscordLink, LogoImage, FeaturedImage, BannerImage)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret)
	fmt.Println("UpdateCollection:", string(data))

	ret2, err := c.GetCollectionsByAccountIndex(AccountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(ret2)
	fmt.Println("GetCollectionsByAccountIndex:", string(data))
}

func TestMintNft(t *testing.T) {

	accountName := "sher.zec"
	var CollectionId int64 = 54
	NftUrl := "-"
	Name := "-"
	Description := "nft-sdk-Description"
	Media := "collection/cbenqstwzx5uy9oedjrb"
	key := fmt.Sprintf("zw:%s:%d", accountName, 2)
	value := "red1"
	assetProperty := Propertie{
		Name:  key,
		Value: value,
	}
	assetLevel := Level{
		Name:     "assetLevel",
		Value:    123,
		MaxValue: 123,
	}
	assetStats := Stat{
		Name:     "StatType",
		Value:    456,
		MaxValue: 456,
	}
	// get content hash

	_Properties := []Propertie{assetProperty}
	_Levels := []Level{assetLevel}
	_Stats := []Stat{assetStats}

	_PropertiesByte, err := json.Marshal(_Properties)
	_LevelsByte, err := json.Marshal(_Levels)
	_StatsByte, err := json.Marshal(_Stats)
	accountSeed := "28e1a3762f....."
	keyManager, err := NewSeedKeyManager(accountSeed)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret, err := c.MintNft(accountName,
		CollectionId,
		NftUrl, Name,
		Description, Media,
		string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret)
	fmt.Println("MintNft:", string(data))

	ret2, err := c.GetNftByNftId(ret.Asset.Id)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(ret2)
	fmt.Println("GetNftByNftId:", string(data))
}
func TestGetNftByNftId(t *testing.T) {
	var nftId int64 = 140
	accountSeed := "28e1a3762f....."
	keyManager, err := NewSeedKeyManager(accountSeed)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret2, err := c.GetNftByNftId(nftId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetNftByNftId:", string(data))
}

func TestTransferNft(t *testing.T) {
	var AssetId int64 = 140
	privateKey := "28e1a3762f....."
	accountName := "sher.zec"
	toAccountName := "gavin.zec"
	keyManager, err := NewSeedKeyManager(privateKey)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret2, err := c.TransferNft(AssetId, accountName, toAccountName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("TransferNft:", string(data))
}

func TestWithdrawNft(t *testing.T) {
	var AssetId int64 = 140
	privateKey := "17673b9a9....."
	accountName := "gavin.zec"
	keyManager, err := NewSeedKeyManager(privateKey)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret2, err := c.WithdrawNft(accountName, AssetId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("WithdrawNft:", string(data))
}

func TestSellOffer(t *testing.T) {
	var AssetId int64 = 139
	privateKey := "28e1a3762f....."
	accountName := "sher.zec"
	keyManager, err := NewSeedKeyManager(privateKey)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)

	ret2, err := c.SellNft(accountName, AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("SellNft:", string(data))
}

func TestBuyOffer(t *testing.T) {
	var AssetId int64 = 139
	privateKey := "17673b9a9....."
	accountName := "gavin.zec"
	keyManager, err := NewSeedKeyManager(privateKey)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)

	ret2, err := c.BuyNft(accountName, AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("BuyNft:", string(data))
}

func TestAcceptOffer(t *testing.T) {
	var offerId int64 = 7
	privateKey := "17673b9a9....."
	accountName := "gavin.zec"
	keyManager, err := NewSeedKeyManager(privateKey)
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl, keyManager)
	ret2, err := c.AcceptOffer(accountName, offerId, false, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("AcceptOffer:", string(data))
}
