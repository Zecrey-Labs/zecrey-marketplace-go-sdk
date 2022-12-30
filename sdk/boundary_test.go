package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"math/big"
	"testing"
)

func TestBoundary_CreateCollection(t *testing.T) {
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	ShortName := "MyNft1"
	CategoryId := "1"
	CollectionUrl := "https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg"
	ExternalLink := "https://weibo.com/alice"
	TwitterLink := "https://twitter.com/alice"
	InstagramLink := "https://www.instagram.com/alice/"
	TelegramLink := "https://tgstat.com/channel/@alice"
	DiscordLink := "https://discord.com/api/v10/applications/<aliceid>/commands"
	LogoImage := "collection/aug788rsfbsnj3i7leqf"
	FeaturedImage := "collection/aug788rsfbsnj3i7leqf"
	BannerImage := "collection/aug788rsfbsnj3i7leqf"
	Description := "Description information"
	CreatorEarningRate := "200"

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	ret, err := c.CreateCollection(ShortName, CategoryId, CreatorEarningRate,
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
	fmt.Println("CreateCollection:", string(data))

	result, err := GetCollectionById(ret.Collection.Id)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(result)
	fmt.Println("GetCollectionById:", string(data))
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
func TestBoundary_MintNft(t *testing.T) {
	var CollectionId int64 = 11
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	NftUrl := "-"
	Name := fmt.Sprintf("nftName1231:%s", accountName)
	Description := fmt.Sprintf("%s `s nft", accountName)
	Media := "collection/aug788rsfbsnj3i7leqf"
	// get content hash
	_Properties := []Propertie{}
	_Levels := []Level{}
	_Stats := []Stat{}

	_PropertiesByte, err := json.Marshal(_Properties)
	_LevelsByte, err := json.Marshal(_Levels)
	_StatsByte, err := json.Marshal(_Stats)

	ret, err := c.MintNft(
		CollectionId,
		NftUrl, Name,
		Description, Media,
		string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret)
	fmt.Println("MintNft:", string(data))
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
