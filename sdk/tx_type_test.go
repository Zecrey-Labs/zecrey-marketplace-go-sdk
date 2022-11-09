package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"math/big"
	"testing"
	"time"
)

func TestParseCreateCollectionTxInfo(t *testing.T) {
	accountName := "sher.zec"
	seed := "28e1a3762f....."
	ShortName := fmt.Sprintf("collection_Shortname %d", time.Now().Second())
	CategoryId := "1"
	CollectionUrl := "-"
	ExternalLink := "-"
	TwitterLink := "-"
	InstagramLink := "-"
	TelegramLink := "-"
	DiscordLink := "-"
	LogoImage := "collection/dekkg5j1ptccq9ttliui"
	FeaturedImage := "collection/dekkg5j1ptccq9ttliui"
	BannerImage := "collection/dekkg5j1ptccq9ttliui"
	Description := "Description information"
	CreatorEarningRate := "1000"
	PaymentAssetIds := "[]"

	c := NewZecreyNftMarketSDK(accountName, seed)
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
		model.WithDescription(Description),
		model.WithPaymentAssetIds(PaymentAssetIds))
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
	var collectionId int64 = 209
	seed := "28e1a37......"
	accountName := "xxx"
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.GetCollectionById(collectionId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetCollectionById:", string(data))
}

func TestGetCollectionsByAccountIndex(t *testing.T) {
	var accountIndex int64 = 4
	accountName := "xxx"
	seed := "ee823...."
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.GetCollectionsByAccountIndex(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GetCollectionsByAccountIndex:", ret2)
}

func TestGetAccountNFTs(t *testing.T) {
	var accountIndex int64 = 4
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.GetAccountNFTs(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GetAccountNFTsConfirmedList:", ret2.ConfirmedAssetIdList)
	fmt.Println("GetAccountNFTsPendingAssets:", ret2.PendingAssets)
}

func TestGetAccountOffers(t *testing.T) {
	var accountIndex int64 = 4
	accountName := "xxx"
	seed := "ee823a...."
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.GetAccountOffers(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GetAccountNFTsConfirmedList:", ret2.ConfirmedOfferIdList)
	fmt.Println("GetAccountNFTsPendingOffers:", ret2.PendingOffers)
}

func TestGetNftOffers(t *testing.T) {
	var nftId int64 = 7
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.GetNftOffers(nftId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GetNFTsConfirmedList:", ret2.ConfirmedOfferIdList)
	fmt.Println("GetNFTsPendingOffers:", ret2.PendingOffers)
}

func TestUpdateCollection(t *testing.T) {
	Id := "54"
	AccountName := "sher"
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
	c := NewZecreyNftMarketSDK(AccountName, seed)
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

	ret2, err := c.GetCollectionsByAccountIndex(AccountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(ret2)
	fmt.Println("GetCollectionsByAccountIndex:", string(data))
}

func TestMintNft(t *testing.T) {
	var CollectionId int64 = 15
	accountName := "sher"
	seed := "28e1a3762ff99...."
	c := NewZecreyNftMarketSDK(accountName, seed)

	NftUrl := "-"
	Name := "-"
	Description := "nft-sdk-Description"
	Media := "collection/dekkg5j1ptccq9ttliui"
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

	ret2, err := c.GetNftById(ret.Asset.Id)
	if err != nil {
		t.Fatal(err)
	}
	data, err = json.Marshal(ret2)
	fmt.Println("GetNftById:", string(data))
}

func TestGetNftByNftId(t *testing.T) {
	var nftId int64 = 9
	accountName := "xxx"
	seed := "ee823a7....."
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.GetNftById(nftId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetNftById:", string(data))
}

func TestGetOfferById(t *testing.T) {
	var OfferId int64 = 5
	accountName := "xxx"
	seed := "ee823a7...."
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.GetOfferById(OfferId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("Offer:", string(data))
}

func TestTransferNft(t *testing.T) {
	var nftId int64 = 6
	accountName := "sher"
	seed := "28e1a37........."
	toAccountName := "gavin"

	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.TransferNft(nftId, toAccountName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("TransferNft:", string(data))
}

func TestWithdrawNft(t *testing.T) {
	var AssetId int64 = 6
	seed := "17673b9a9fd.........."
	accountName := "gavin"

	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.WithdrawNft(AssetId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("WithdrawNft:", string(data))
}

func TestSellOffer(t *testing.T) {
	var AssetId int64 = 11
	accountName := "sher"
	seed := "28e1a3762ff9.........."

	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.CreateSellOffer(AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("CreateSellOffer:", string(data))
}

func TestBuyOffer(t *testing.T) {
	var AssetId int64 = 9
	accountName := "sher"
	seed := "28e1a3762ff99...."

	c := NewZecreyNftMarketSDK(accountName, seed)

	ret2, err := c.CreateBuyOffer(AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("CreateBuyOffer:", string(data))
}
func TestCancelOffer(t *testing.T) {
	var OfferId int64 = 9
	accountName := "sher"
	seed := "28e1a3762ff99...."
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.CancelOffer(OfferId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("CancelOffer:", string(data))
}

func TestAcceptOffer(t *testing.T) {
	var offerId int64 = 21
	seed := "28e1a3762ff99...."
	accountName := "gavin"
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret2, err := c.AcceptOffer(offerId, false, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("AcceptOffer:", string(data))
}

func TestCreateL1Account(t *testing.T) {
	l1Addr, privateKeyStr, l2pk, seed, err := CreateL1Account()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("l1Addr:", l1Addr)
	fmt.Println("seed:", seed)
	fmt.Println("privateKeyStr:", privateKeyStr)
	fmt.Println("l2pk:", l2pk)
	fmt.Println("err:", err)
	//l1Addr: 0xD207262DEA01aE806fA2dCaEdd489Bd2f5FABcFE
	//seed: 0x6a1a320d14790f2d9aa9a37769f4833d583a3f7f974fd452a3990aeb0e7a6052
	//privateKeyStr: 1a061a8e74cee1ce2e2ddd29f5afea99ecfbaf1998b6d349a8c09a368e637b8e
	//l2pk: 06278b99871f1d64fcc83bd27713cbf743d957c510a245d6bfb0eae888e35452274a2b4c8c7b7424f25d7d187661225111753197248fa045fd872aa662fdcb24
}

func TestGetSeedAndL2Pk(t *testing.T) {
	privateKey := "0xe94a8b4ddd33b2865a89bb764d70a0c3e3276007ece8f114a47a4e9581ec3567"
	l2pk, seed, err := GetSeedAndL2Pk(privateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("seed:", seed)
	fmt.Println("l2pk:", l2pk)
	fmt.Println("err:", err)
	//l1Addr: 0xD207262DEA01aE806fA2dCaEdd489Bd2f5FABcFE
	//seed: 0x6a1a320d14790f2d9aa9a37769f4833d583a3f7f974fd452a3990aeb0e7a6052
	//privateKeyStr: 1a061a8e74cee1ce2e2ddd29f5afea99ecfbaf1998b6d349a8c09a368e637b8e
	//l2pk: 22fc6f5d74c8639245462a0af6b5c931bd209c04034b28421a60336635ab85950a3163e68ec29319ca200fac009408369b0a1f75200a118aded920cd240e1358
}

func TestRegisterAccountWithPrivateKey(t *testing.T) {
	accountName := "zhangwei1111"
	l1Addr := "0x805e286D05388911cCdB10E3c7b9713415607c72"
	//seed := "0x7ea589236ac7e6034a40ad31f27a6ea1bbaeb7746ba5e8d3408a3abb480a8688"
	//l2pk := "22fc6f5d74c8639245462a0af6b5c931bd209c04034b28421a60336635ab85950a3163e68ec29319ca200fac009408369b0a1f75200a118aded920cd240e1358"
	privateKey := "0xe94a8b4ddd33b2865a89bb764d70a0c3e3276007ece8f114a47a4e9581ec3567"
	client, err := RegisterAccountWithPrivateKey(accountName, l1Addr, privateKey)
	if err != nil {
		t.Fatal(err)
	}
	accountName, l2pk, seed := client.GetMyInfo()
	fmt.Println(fmt.Sprintf("registerAccountRet:\naccountName=%s\nl2pk=%s\nseed=%s", accountName, l2pk, seed))
}

func TestGetAccountIsRegistered(t *testing.T) {
	accountName := "6633332"
	ret2, err := GetAccountIsRegistered(accountName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(ret2)
	fmt.Println("GetAccountIsRegistered:", string(data))
}

func TestGetAccountByAccountName(t *testing.T) {
	accountName := "xxx"
	seed := "ee823a........."
	c := NewZecreyNftMarketSDK(accountName, seed)
	accountInfo, err := c.GetAccountByAccountName(accountName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(accountInfo)
	fmt.Println(string(data))
}

func TestApplyRegisterHost(t *testing.T) {
	l1Addr := "0x805e286D05388911cCdB10E3c7b9713415607c72"
	l2pk := "22fc6f5d74c8639245462a0af6b5c931bd209c04034b28421a60336635ab85950a3163e68ec29319ca200fac009408369b0a1f75200a118aded920cd240e1358"
	accountName := "6633332"
	seed := "0x7ea589236ac7e6034a40ad31f27a6ea1bbaeb7746ba5e8d3408a3abb480a8688"
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret, err := c.ApplyRegisterHost(fmt.Sprintf("%s", accountName), l2pk, l1Addr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp: %v\n", ret.Ok)
}

func TestUploadMeida(t *testing.T) {
	accountName := "xxx"
	seed := "0x7ea5892......."
	filePath := "/Users/xxx/Downloads/xxx.jpg"
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret, err := c.UploadMedia(filePath)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp: %v\n", ret)
}

func TestGetCategories(t *testing.T) {
	accountName := "xxx"
	seed := "ee823a7........."
	c := NewZecreyNftMarketSDK(accountName, seed)
	ret, err := c.GetCategories()
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range ret.Categories {
		fmt.Printf("Categories Id: %d\n", v.Id)
		fmt.Printf("Categories Name: %s\n", v.Name)
	}
}
