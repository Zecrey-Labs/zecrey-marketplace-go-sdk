package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"io/ioutil"
	"math/big"
	"strings"
	"testing"
)

func TestCreateCollection(t *testing.T) {
	accountName := "amber1"
	seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
	ShortName := "MyNft"
	CategoryId := "1"
	CollectionUrl := "https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg"
	ExternalLink := "https://weibo.com/alice"
	TwitterLink := "https://twitter.com/alice"
	InstagramLink := "https://www.instagram.com/alice/"
	TelegramLink := "https://tgstat.com/channel/@alice"
	DiscordLink := "https://discord.com/api/v10/applications/<aliceid>/commands"
	LogoImage := "collection/tvbish3l1djoeqfgqjy9"
	FeaturedImage := "collection/tvbish3l1djoeqfgqjy9"
	BannerImage := "collection/tvbish3l1djoeqfgqjy9"
	Description := "Description information"
	CreatorEarningRate := "1000"

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

func TestGetCollectionById(t *testing.T) {
	var collectionId int64 = 209
	result, err := GetCollectionById(collectionId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestGetCollectionsByAccountIndex(t *testing.T) {
	var accountIndex int64 = 4
	result, err := GetCollectionsByAccountIndex(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestGetAccountNFTs(t *testing.T) {
	var accountIndex int64 = 4
	result, err := GetAccountNFTs(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func TestGetAccountOffers(t *testing.T) {
	var accountIndex int64 = 4
	result, err := GetAccountOffers(accountIndex)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestGetNftOffers(t *testing.T) {
	var nftId int64 = 7
	result, err := GetNftOffers(nftId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestUpdateCollection(t *testing.T) {
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

func TestMintNft(t *testing.T) {
	nfts := map[string]string{

		//"Lizard_34_Station":  "collection/fhqtogwwejr9ecpbzglt",
		"Lizard_34_Broken": "collection/ukosklfy7zxiodcacqx8",
		"collection_cover": "collection/dl4c13dewzsbwmjpqfk1",
		"Lizard_34_Scape":  "collection/niphsjvc4wyvjed15srp",
		"Lizard_34_Tidal":  "collection/kslbwhiwavoifvbapuot",
		"Lizard_34_Sunset": "collection/emwbiowe1zfonhbs7us2",
		"Lizard_34_Moon":   "collection/iyiffvmegjbxkxy63jtd",
	}
	for nftName, url := range nfts {
		var CollectionId int64 = 12
		accountName := "amber1"
		seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f "
		c, err := NewClient(accountName, seed)
		if err != nil {
			t.Fatal(err)
		}

		NftUrl := "-"
		Name := nftName
		Description := nftName
		Media := url
		// get content hash
		_Properties := []Propertie{}
		_Levels := []Level{}
		_Stats := []Stat{}

		_PropertiesByte, err := json.Marshal(_Properties)
		_LevelsByte, err := json.Marshal(_Levels)
		_StatsByte, err := json.Marshal(_Stats)

		ret, err := c.MintNft(
			CollectionId,
			NftUrl, Name, 20,
			Description, Media,
			string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
		if err != nil {
			t.Fatal(err)
		}
		data, err := json.Marshal(ret)
		fmt.Println("MintNft:", string(data))
		result, err := GetNftById(ret.Asset.Id)
		if err != nil {
			t.Fatal(err)
		}
		data, err = json.Marshal(result)
		fmt.Println("GetNftById:", string(data))
	}

}

func TestGetNftByNftId(t *testing.T) {
	var nftId int64 = 24
	result, err := GetNftById(nftId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestGetOfferById(t *testing.T) {
	var OfferId int64 = 5
	result, err := GetOfferById(OfferId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("Offer:", string(data))
}

func TestTransferNft(t *testing.T) {
	var nftId int64 = 6
	accountName := "sher"
	seed := "28e1a37........."
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

func TestWithdrawNft(t *testing.T) {
	var AssetId int64 = 6
	seed := "17673b9a9fd.........."
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

func TestSellOffer(t *testing.T) {
	var AssetId int64 = 383
	accountName := "sher"
	seed := "28e1a3762ff9.........."

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.CreateSellOffer(AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("CreateSellOffer:", string(data))
}

func TestBuyOffer(t *testing.T) {
	var AssetId int64 = 9
	accountName := "sher"
	seed := "28e1a3762ff99...."

	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}

	result, err := c.CreateBuyOffer(AssetId, 0, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("CreateBuyOffer:", string(data))
}

func TestCancelOffer(t *testing.T) {
	var OfferId int64 = 9
	accountName := "sher"
	seed := "28e1a3762ff99...."
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

func TestAcceptOffer(t *testing.T) {
	var offerId int64 = 21
	seed := "28e1a3762ff99...."
	accountName := "gavin"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.AcceptOffer(offerId, false, big.NewInt(1000000))
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
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
	result, err := IfAccountRegistered(accountName)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("IfAccountRegistered:", string(data))
}

func TestGetAccountByAccountName(t *testing.T) {
	accountName := "alice"
	accountInfo, err := GetAccountByAccountName(accountName)
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
	ret, err := ApplyRegisterHost(fmt.Sprintf("%s", accountName), l2pk, l1Addr)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp: %v\n", ret.Ok)
}

func TestUploadMeida(t *testing.T) {
	filePath := "/Users/zhangwei/Documents/飞书20221112-142909.jpg"
	result, err := UploadMedia(filePath)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestGetCategories(t *testing.T) {
	result, err := GetCategories()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestGetListingOffers(t *testing.T) {
	var isSell int64 = 1
	result, err := GetListingOffers(isSell)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

func TestUploadMediaInBatch(t *testing.T) {
	paths := []string{
		//"/Users/zhangwei/Documents/collection2222/Weather Map of Lizard-34",
		"/Users/zhangwei/Documents/collection2222/Portrait",
		//"/Users/zhangwei/Documents/collection2222/Piece",
		//"/Users/zhangwei/Documents/collection2222/End of the World",
		//"/Users/zhangwei/Documents/collection2222/Comics",
		//"/Users/zhangwei/Documents/collection2222/CG",
		//"/Users/zhangwei/Documents/collection2222/Blocks",
		//"/Users/zhangwei/Documents/collection2222/Ball",
		//"/Users/zhangwei/Documents/collection2222/8 Bits"
	}
	for p := 0; p < len(paths); p++ {
		path := paths[p]
		files, _ := ioutil.ReadDir(path)
		for i := 0; i < len(files); i++ {
			result, err := UploadMedia(path + "/" + files[i].Name())
			if err == nil {
				strs := strings.Split(path, "/")
				cName := strs[len(strs)-1]
				fn := files[i].Name()
				fn1 := strings.TrimSuffix(fn, "_cover.png")
				fn1 = strings.TrimSuffix(fn, "_icon.png")
				fn1 = strings.TrimSuffix(fn, "_icon.jpg")
				fn1 = strings.TrimSuffix(fn, "_icon.jpg")
				fn1 = strings.TrimSuffix(fn, ".jpg")
				fn1 = strings.TrimSuffix(fn, ".png")
				//collection cover
				if strings.Contains(fn, "_cover.png") {
					fmt.Println(fmt.Sprintf("Collection Cover:\"%s\":{\"%s\",\"none\"},", cName, result.PublicId))
					continue
				}
				//collection
				if strings.Contains(fn, "_icon.png") {
					fmt.Println(fmt.Sprintf("Collection Icon:\"%s\":{\"none\",\"%s\"},", cName, result.PublicId))
					continue
				}
				//nft
				fmt.Println(fmt.Sprintf("Collection:%s nft:\"%s\":\"%s\",", cName, fn1, result.PublicId))
			}
		}
	}
}

func TestCreateCollectionInBatch(t *testing.T) {
	collections := map[string][]string{
		"Lizard_34_Broken": {"collection/ahykviwc0suhoyzusb5q", "collection/ahykviwc0suhoyzusb5q",
			"There is a group of Lizardmen living on an abandoned mine planet. They are divided into three races, namely Zecrey, Komodoensis and Reptoids. The Lizardman is a highly civilized creature in the universe, but according to the current declassified data, no one can know where the Lizardman came from."},
	}
	for CName, Infos := range collections {
		accountName := "amber1"
		seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f"
		ShortName := CName
		CategoryId := "1"
		LogoImage := Infos[0]
		FeaturedImage := Infos[0]
		BannerImage := Infos[1]
		Description := Infos[2]
		CreatorEarningRate := "1000"
		c, err := NewClient(accountName, seed)
		if err != nil {
			t.Fatal(err)
		}
		ret, err := c.CreateCollection(ShortName, CategoryId, CreatorEarningRate,
			model.WithLogoImage(LogoImage),
			model.WithFeaturedImage(FeaturedImage),
			model.WithBannerImage(BannerImage),
			model.WithDescription(Description))
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("%s ,%d", ShortName, ret.Collection.Id))
	}
}

/*
   End of the World,12
*/
func TestMintNftInBatch(t *testing.T) {
	nfts := map[string]string{
		//"End of the World":  "collection/fhqtogwwejr9ecpbzglt",
	}
	for nftName, url := range nfts {
		var CollectionId int64 = 12
		accountName := "amber1"
		seed := "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f "
		c, err := NewClient(accountName, seed)
		if err != nil {
			t.Fatal(err)
		}

		NftUrl := "-"
		Name := nftName
		Description := nftName
		Media := url
		// get content hash
		var _Properties []Propertie
		var _Levels []Level
		var _Stats []Stat

		_PropertiesByte, err := json.Marshal(_Properties)
		_LevelsByte, err := json.Marshal(_Levels)
		_StatsByte, err := json.Marshal(_Stats)

		ret, err := c.MintNft(
			CollectionId,
			NftUrl, Name, 20,
			Description, Media,
			string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
		if err != nil {
			t.Fatal(err)
		}
		data, err := json.Marshal(ret)
		fmt.Println("MintNft:", string(data))
		result, err := GetNftById(ret.Asset.Id)
		if err != nil {
			t.Fatal(err)
		}
		data, err = json.Marshal(result)
		fmt.Println("GetNftById:", string(data))
	}

}
