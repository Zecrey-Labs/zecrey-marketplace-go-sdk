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
	nftMarketUrl string = "https://test-legend-nft.zecrey.com"
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
	accountSeed := "28e1a376........."
	ShortName := fmt.Sprintf("collection_Shortname %d", time.Now().Second())
	CategoryId := "1"
	CollectionUrl := "-"
	ExternalLink := "-"
	TwitterLink := "-"
	InstagramLink := "-"
	TelegramLink := "-"
	DiscordLink := "-"
	LogoImage := "collection/e2rdcsgitkxonxqjcmkg"
	FeaturedImage := "collection/e2rdcsgitkxonxqjcmkg"
	BannerImage := "collection/e2rdcsgitkxonxqjcmkg"
	CreatorEarningRate := "6666"
	PaymentAssetIds := "[]"
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret, err := c.CreateCollection(accountName, accountSeed, ShortName, CategoryId, CollectionUrl,
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
	var collectionId int64 = 142
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret2, err := c.GetCollectionById(collectionId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetCollectionById:", string(data))
}

func TestGetCollectionByAccountIndex(t *testing.T) {
	var accountIndex int64 = 4
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret2, err := c.GetCollectionsByAccountIndex(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetCollectionsByAccountIndex:", string(data))
}

func TestUpdateCollection(t *testing.T) {
	Id := "142"
	AccountName := "sher.zec"
	privateKey := "28e1a376........."
	Name := "zw-sdk--collection-update"
	CollectionUrl := "-"
	Description := "-"
	CategoryId := "1"
	ExternalLink := "-"
	TwitterLink := "-"
	InstagramLink := "-"
	TelegramLink := "-"
	DiscordLink := "-"
	LogoImage := "collection/e2rdcsgitkxonxqjcmkg"
	FeaturedImage := "collection/e2rdcsgitkxonxqjcmkg"
	BannerImage := "collection/e2rdcsgitkxonxqjcmkg"

	var AccountIndex int64 = 2

	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret, err := c.UpdateCollection(Id, AccountName, privateKey,
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

	privateKey := "28e1a376........."
	accountName := "sher.zec"
	var CollectionId int64 = 142
	var l2NftCollectionId int64 = 0
	NftUrl := "-"
	Name := "-"
	Description := "nft-sdk-Description"
	Media := "collection/e2rdcsgitkxonxqjcmkg"
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
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret, err := c.MintNft(privateKey, accountName,
		CollectionId, l2NftCollectionId,
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
	var nftId int64 = 301
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret2, err := c.GetNftByNftId(nftId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetNftByNftId:", string(data))
}

func TestTransferNft(t *testing.T) {
	var AssetId int64 = 301
	privateKey := "28e1a376........."
	accountName := "sher.zec"
	toAccountName := "gavin.zec"
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret2, err := c.TransferNft(AssetId, privateKey, accountName, toAccountName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("TransferNft:", string(data))
}

func TestWithdrawNft(t *testing.T) {
	var AssetId int64 = 301
	privateKey := "17673b.........."
	accountName := "gavin.zec"
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret2, err := c.WithdrawNft(privateKey, accountName, AssetId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("WithdrawNft:", string(data))
}

func TestSellOffer(t *testing.T) {
	var AssetId int64 = 215
	privateKey := "28e1a376........."
	accountName := "sher.zec"
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)

	ret2, err := c.SellNft(privateKey, accountName, AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("SellNft:", string(data))
}

func TestBuyOffer(t *testing.T) {
	var AssetId int64 = 217
	privateKey := "28e1a376........."
	accountName := "sher.zec"
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)

	ret2, err := c.BuyNft(privateKey, accountName, AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("BuyNft:", string(data))
}

func TestAcceptOffer(t *testing.T) {
	var offerId int64 = 199
	privateKey := "ee823a743f......"
	accountName := "amber1.zec"
	c := NewZecreyNftMarketSDK(legendUrl, nftMarketUrl)
	ret2, err := c.AcceptOffer(privateKey, accountName, offerId, true, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("AcceptOffer:", string(data))
}
