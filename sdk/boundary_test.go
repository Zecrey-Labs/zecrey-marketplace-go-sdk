package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

var (
	//len([]byte(boundaryStr))  = 164
	boundaryStr = "https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg"
	//len([]byte(boundaryStr))  = 1148
	boundaryStr2 = "https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg" +
		"https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg"
)

func TestBoundary_CreateCollection_url(t *testing.T) {
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	ShortName := "MyNft1"
	CategoryId := "1"
	//size 128
	CollectionUrl := boundaryStr
	ExternalLink := boundaryStr
	TwitterLink := boundaryStr
	InstagramLink := boundaryStr
	TelegramLink := boundaryStr
	DiscordLink := boundaryStr
	LogoImage := "collection/haxiuyotbowltzv5ubok"
	FeaturedImage := "collection/haxiuyotbowltzv5ubok"
	BannerImage := "collection/haxiuyotbowltzv5ubok"
	Description := "Description information"
	CreatorEarningRate := "200"

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.CreateCollection(ShortName, CategoryId, CreatorEarningRate,
		model.WithCollectionUrl(CollectionUrl),
		model.WithExternalLink(ExternalLink),
		model.WithTwitterLink(TwitterLink),
		model.WithInstagramLink(InstagramLink),
		model.WithTelegramLink(TelegramLink),
		model.WithDiscordLink(DiscordLink),
		model.WithLogoImage(LogoImage),
		model.WithFeaturedImage(FeaturedImage),
		model.WithBannerImage(BannerImage),
		model.WithDescription(Description))
	if err != nil {
		if strings.Contains(err.Error(), fmt.Errorf("value too long for type character varying(128) (SQLSTATE 22001)").Error()) {
			fmt.Println("pass")
			return
		}
		t.Fatal(err)
	}
}

//TestBoundary_CreateCollection_CreatorEarningRate test CreatorEarningRate
func TestBoundary_CreateCollection_CreatorEarningRate(t *testing.T) {
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	ShortName := "MyNft4"
	CategoryId := "1"
	LogoImage := "collection/haxiuyotbowltzv5ubok"
	FeaturedImage := "collection/haxiuyotbowltzv5ubok"
	BannerImage := "collection/haxiuyotbowltzv5ubok"
	Description := "Description information"
	CreatorEarningRate := "2000000000000000000000000000"
	//CreatorEarningRate := "-200"

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.CreateCollection(ShortName, CategoryId, CreatorEarningRate,
		model.WithLogoImage(LogoImage),
		model.WithFeaturedImage(FeaturedImage),
		model.WithBannerImage(BannerImage),
		model.WithDescription(Description))
	if err != nil {
		if strings.Contains(err.Error(), fmt.Errorf("cannot parsed as int").Error()) {
			fmt.Println("pass")
			return
		}
		t.Fatal(err)
	}
}

func TestBoundary_CreateCollection_VeritySign(t *testing.T) {
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	ShortName := "MyNft5"
	CategoryId := "1"
	LogoImage := "collection/haxiuyotbowltzv5ubok"
	FeaturedImage := "collection/haxiuyotbowltzv5ubok"
	BannerImage := "collection/haxiuyotbowltzv5ubok"
	Description := "Description information"
	CreatorEarningRate := "10"

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCreateCollectionTxInfo?account_name=%s", c.accountName))
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		t.Fatal(err)
	}
	if respSdkTx.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf(string(body)))
	}

	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		t.Fatal(err)
	}
	tx, err := sdkCreateCollectionTxInfo(c.keyManager, resultSdk.Transtion, Description, ShortName)
	if err != nil {
		t.Fatal(err)
	}
	//--------- change value ------
	txInfo := &CreateCollectionTxInfo{}
	err = json.Unmarshal([]byte(tx), txInfo)
	txInfo.AccountIndex = 2
	data, _ := json.Marshal(txInfo)

	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/collection/createCollection",
		url.Values{
			"short_name":           {ShortName},
			"category_id":          {CategoryId},
			"collection_url":       {"-"},
			"external_link":        {"-"},
			"twitter_link":         {"-"},
			"instagram_link":       {"-"},
			"discord_link":         {"-"},
			"telegram_link":        {"-"},
			"logo_image":           {LogoImage},
			"featured_image":       {FeaturedImage},
			"banner_image":         {BannerImage},
			"creator_earning_rate": {CreatorEarningRate},
			"payment_asset_ids":    {"[]"},
			"transaction":          {string(data)}})
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		if strings.Contains(string(body), fmt.Errorf("invalid signature").Error()) {
			fmt.Println("pass")
			return
		}
		t.Fatal(fmt.Errorf(string(body)))
	}
}

func TestBoundary_UpdateCollection(t *testing.T) {
	Id := "54"
	accountName := "sher"
	Name := "zw-sdk--collection-update"
	CollectionUrl := "-"
	Description := "-"
	ExternalLink := "-"
	TwitterLink := "-"
	InstagramLink := "-"
	TelegramLink := "-"
	DiscordLink := "-"
	LogoImage := "collection/cbenqstwzx5uy9oedjrb"
	FeaturedImage := "collection/cbenqstwzx5uy9oedjrb"
	BannerImage := "collection/cbenqstwzx5uy9oedjrb"

	var AccountIndex int64 = 2
	seed := "28e1a3762f....."
	c, err := NewClient(accountName, seed)
	if err != nil {
		panic(err)
	}
	ret, err := c.UpdateCollection(Id, Name,
		model.WithCollectionUrl(CollectionUrl),
		model.WithExternalLink(ExternalLink),
		model.WithTwitterLink(TwitterLink),
		model.WithInstagramLink(InstagramLink),
		model.WithTelegramLink(TelegramLink),
		model.WithDiscordLink(DiscordLink),
		model.WithLogoImage(LogoImage),
		model.WithFeaturedImage(FeaturedImage),
		model.WithBannerImage(BannerImage),
		model.WithDescription(Description))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret)
	fmt.Println("UpdateCollection:", string(data))

	result, err := GetCollectionsByAccountIndex(AccountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(result)
	fmt.Println("GetCollectionsByAccountIndex:", string(data))
}
func TestBoundary_MintNft_UrlOrName(t *testing.T) {
	var CollectionId int64 = 81
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	NftUrl := boundaryStr2
	Name := fmt.Sprintf("%s", boundaryStr2)
	Description := fmt.Sprintf("%s `s nft", boundaryStr2)
	Media := "collection/aug788rsfbsnj3i7leqf"
	// get content hash
	var _Properties []Propertie
	var _Levels []Level
	var _Stats []Stat

	_PropertiesByte, err := json.Marshal(_Properties)
	_LevelsByte, err := json.Marshal(_Levels)
	_StatsByte, err := json.Marshal(_Stats)

	ret, err := c.MintNft(
		CollectionId,
		NftUrl, Name,
		Description, Media,
		string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
	if err != nil {
		if strings.Contains(err.Error(), fmt.Errorf("value too long for type character varying(1024) (SQLSTATE 22001)").Error()) {
			fmt.Println("pass")
			return
		}
		t.Fatal(err)
	}
	data, err := json.Marshal(ret)
	fmt.Println("MintNft:", string(data))
}

//TestBoundary_MintNft_CreatorEarningRate  test txInfo.CreatorTreasuryRate = -1 or txInfo.CreatorTreasuryRate > 10000
func TestBoundary_MintNft_CreatorEarningRate(t *testing.T) {
	var CollectionId int64 = 81
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	NftUrl := "-"
	Name := fmt.Sprintf("nftName 2:%s", accountName)
	Description := fmt.Sprintf("%s `s nft", accountName)
	Media := "collection/aug788rsfbsnj3i7leqf"
	// get content hash
	var _Properties []Propertie
	var _Levels []Level
	var _Stats []Stat

	_PropertiesByte, err := json.Marshal(_Properties)
	_LevelsByte, err := json.Marshal(_Levels)
	_StatsByte, err := json.Marshal(_Stats)

	ContentHash, err := calculateContentHash(c.accountName, CollectionId, Name, string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
	respSdkTx, err := http.Get(c.nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkMintNftTxInfo?treasury_rate=20&account_name=%s&collection_id=%d&name=%s&content_hash=%s", c.accountName, CollectionId, Name, ContentHash))
	if err != nil {
		t.Fatal(err)
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		t.Fatal(err)
	}
	if respSdkTx.StatusCode != http.StatusOK {
		t.Fatal(fmt.Errorf(string(body)))
	}
	resultSdk := &RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		t.Fatal(err)
	}
	txInfo := &MintNftTxInfo{}
	err = json.Unmarshal([]byte(resultSdk.Transtion), txInfo)
	if err != nil {
		t.Fatal(err)
	}
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	//txInfo.CreatorTreasuryRate = -1
	txInfo.CreatorTreasuryRate = 1000000000
	tx, err := constructMintNftTx(c.keyManager, txInfo)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.PostForm(c.nftMarketUrl+"/api/v1/asset/createAsset",
		url.Values{
			"collection_id": {fmt.Sprintf("%d", CollectionId)},
			"nft_url":       {NftUrl},
			"name":          {Name},
			"description":   {Description},
			"media":         {Media},
			"properties":    {string(_PropertiesByte)},
			"levels":        {string(_LevelsByte)},
			"stats":         {string(_StatsByte)},
			"transaction":   {tx},
		},
	)

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		if strings.Contains(string(body), fmt.Errorf("CreatorTreasuryRate should  not be less than 0").Error()) {
			fmt.Println("pass")
			return
		}
		if strings.Contains(string(body), fmt.Errorf("CreatorTreasuryRate should not be larger than 65535").Error()) {
			fmt.Println("pass")
			return
		}
		t.Fatal(fmt.Errorf(string(body)))
	}
	result := &RespCreateAsset{}
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatal(err)
	}

	data, err := json.Marshal(result)
	fmt.Println("MintNft:", string(data))
	fmt.Println("pass")
}
func TestBoundary_TransferNft(t *testing.T) {
	var nftId int64 = 4
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	toAccountName := "gavin"

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.TransferNft(nftId, toAccountName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("TransferNft:", string(data))
}
func TestBoundary_WithdrawNft(t *testing.T) {
	var AssetId int64 = 4
	seed := "17673b9a9fdec6dc90c7cc1eb1c939134dfb659d2f08edbe071e5c45f343d008"
	accountName := "gavin"

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.WithdrawNft(AssetId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("WithdrawNft:", string(data))
}
func TestBoundary_SellOffer(t *testing.T) {
	var AssetId int64 = 3
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	accountName := "amber1"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.CreateSellOffer(AssetId, 0, big.NewInt(1230000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("CreateSellOffer:", string(data))
}
func TestBoundary_BuyOffer(t *testing.T) {
	var AssetId int64 = 3
	seed := "17673b9a9fdec6dc90c7cc1eb1c939134dfb659d2f08edbe071e5c45f343d008"
	accountName := "gavin"

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}

	result, err := c.CreateBuyOffer(AssetId, 0, big.NewInt(1230000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("CreateBuyOffer:", string(data))
}
func TestBoundary_CancelOffer(t *testing.T) {
	var OfferId int64 = 4
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	accountName := "amber1"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.CancelOffer(OfferId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("CancelOffer:", string(data))
}
func TestBoundary_AcceptOffer(t *testing.T) {
	var offerId int64 = 6
	seed := "17673b9a9fdec6dc90c7cc1eb1c939134dfb659d2f08edbe071e5c45f343d008"
	accountName := "gavin"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.AcceptOffer(offerId, false, big.NewInt(1230000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("AcceptOffer:", string(data))
}
