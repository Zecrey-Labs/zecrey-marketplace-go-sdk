package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/ethereum/go-ethereum/common"
	vegeta "github.com/tsenart/vegeta/v12/lib"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestCreateCollection(t *testing.T) {
	accountName := "alice"
	seed := "asdfasdfasdf98fd05c70sdafasdfasdffdasdfsadfsdfsfdasdf30383efcb3954631"
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

func TestGetCollectionById(t *testing.T) {
	var collectionId int64 = 7
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
	var accountIndex int64 = 2
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
	var nftId int64 = 3
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
	accountName := "jarry"
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
	var CollectionId int64 = 93
	seed := "13243b9a9fdec6dc90c7cc1eb1c939134dfb659d2f0asdfas5413213213213213"
	accountName := "bob"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	NftUrl := "-"
	Name := fmt.Sprintf("nftName2:%s", accountName)
	Description := fmt.Sprintf("%s `s nft", accountName)
	Media := "collection/dz5hwqaszpwtflg0sfz4"
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
		t.Fatal(err)
	}
	data, err := json.Marshal(ret)
	fmt.Println("MintNft:", string(data))
}

func TestGetNftByNftId(t *testing.T) {
	var nftId int64 = 2
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
	var OfferId int64 = 1
	result, err := GetOfferById(OfferId)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("Offer:", string(data))
}

func TestTransferNft(t *testing.T) {
	seed := "13243b9a9fdec6dc90c7cc1eb1c939134dfb659d2f0asdfas5413213213213213"
	accountName := "bob"

	var nftId int64 = 277
	//accountName := "alice"
	//seed := "asdfasdfasdf98fd05c70sdafasdfasdffdasdfsadfsdfsfdasdf30383efcb3954631"
	toAccountName := "alice"

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
	var AssetId int64 = 277
	seed := "13243b9a9fdec6dc90c7cc1eb1c939134dfb659d2f0asdfas5413213213213213"
	accountName := "bob"
	tol1Address := "0x< a l1 address you want to withdraw nft>"
	c, err := NewClient(accountName, seed)
	if err != nil {
		t.Fatal(err)
	}
	result, err := c.WithdrawNft(AssetId, tol1Address)
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	fmt.Println("WithdrawNft:", string(data))
}

func TestSellOffer(t *testing.T) {
	var AssetId int64 = 3
	seed := "asdfasdfasdf98fd05c70sdafasdfasdffdasdfsadfsdfsfdasdf30383efcb3954631"
	accountName := "alice"
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

func TestBuyOffer(t *testing.T) {
	var AssetId int64 = 3
	seed := "13243b9a9fdec6dc90c7cc1eb1c939134dfb659d2f0asdfas5413213213213213"
	accountName := "bob"

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

func TestCancelOffer(t *testing.T) {
	var OfferId int64 = 4
	seed := "asdfasdfasdf98fd05c70sdafasdfasdffdasdfsadfsdfsfdasdf30383efcb3954631"
	accountName := "alice"
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
	var offerId int64 = 6
	seed := "13243b9a9fdec6dc90c7cc1eb1c939134dfb659d2f0asdfas5413213213213213"
	accountName := "bob"
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
	accountName := "alice"
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

func TestDepositNft(t *testing.T) {
	accountName := "alice"
	privateKey := "0xe94a8b4ddd33b2865a89bb764d70a0c3e3276007ece8f114a47a4e9581ec3567"
	_nftL1Address := common.HexToAddress("0x805e286D05388911cCdB10E3c7b9713415607c72")
	_nftL1TokenId := big.NewInt(511)

	depositNftTransaction, err := DepositNft(accountName, privateKey, _nftL1Address, _nftL1TokenId)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(depositNftTransaction)
}
func TestDepositBNB(t *testing.T) {
	accountName := "alice"
	privateKey := "0xe94a8b4ddd33b2865a89bb764d70a0c3e3276007ece8f114a47a4e9581ec3567"
	depositBnbTransaction, err := DepositBNB(accountName, privateKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(depositBnbTransaction)
}
func TestFullExit(t *testing.T) {
	accountName := "alice"
	privateKey := "0xe94a8b4ddd33b2865a89bb764d70a0c3e3276007ece8f114a47a4e9581ec3567"
	_asset := common.HexToAddress("0x805e286D05388911cCdB10E3c7b9713415607c72")
	fullExitTransaction, err := FullExit(accountName, privateKey, _asset)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fullExitTransaction)
}
func TestFullExitNft(t *testing.T) {
	accountName := "alice"
	privateKey := "0xe94a8b4ddd33b2865a89bb764d70a0c3e3276007ece8f114a47a4e9581ec3567"
	_nftIndex := uint32(511)
	fullExitNftTransaction, err := FullExitNft(accountName, privateKey, _nftIndex)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fullExitNftTransaction)
}
func TestWithdraw(t *testing.T) {
	accountName := "alice"
	privateKey := "0xe94a8b4ddd33b2865a89bb764d70a0c3e3276007ece8f114a47a4e9581ec3567"
	_owner := common.HexToAddress("0x805e286D05388911cCdB10E3c7b9713415607c72")
	_token := common.HexToAddress("0x805e286D05388911cCdB10E3c7b9713415607c72")
	_amount := big.NewInt(10000)

	withdrawTransaction, err := Withdraw(accountName, privateKey, _owner, _token, _amount)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(withdrawTransaction)
}

func TestGetAccountIsRegistered(t *testing.T) {
	accountName := "alice"
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

func TestGetAccountIndex(t *testing.T) {
	accountName := "alice"
	accountInfo, err := GetAccountIndex(accountName)
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
	filePath := "/Users/user0/Documents/collection2222/8 Bits/8_Bit_Komo_Shooter.png"
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

func TestGetGetLayer2BasicInfo(t *testing.T) {
	result, err := GetLayer2BasicInfo()
	if err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}

/* batch create nft  */
func TestUploadMediaInBatch(t *testing.T) {
	paths := []string{
		//"/Users/user0/Documents/collection2222/Weather Map of Lizard-34",
		"/Users/user0/Documents/collection2222/Portrait",
		"/Users/user0/Documents/collection2222/Piece",
		"/Users/user0/Documents/collection2222/End of the World",
		"/Users/user0/Documents/collection2222/Comics",
		"/Users/user0/Documents/collection2222/CG",
		"/Users/user0/Documents/collection2222/Blocks",
		"/Users/user0/Documents/collection2222/Ball",
		"/Users/user0/Documents/collection2222/8 Bits",
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
				fn1 = strings.TrimSuffix(fn, "_cover.jpg")
				fn1 = strings.TrimSuffix(fn, "_icon.jpg")
				fn1 = strings.TrimSuffix(fn, "_icon")
				fn1 = strings.TrimSuffix(fn, "_icon")
				fn1 = strings.TrimSuffix(fn, "")
				fn1 = strings.TrimSuffix(fn, ".png")
				fn1 = strings.TrimSuffix(fn, ".jpg")
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
		"Portrait": {"collection/xj5dg8wxxlu53iaaxq1s", "collection/m1bxivrprc2sxykorvjq",
			"Digital avatars of Bob's friends."},
		"End of the World": {"collection/glsvirjr4uizzfpbpg4e", "collection/vnywfhyqysjdnb5zopth", "The resources on the Lizard Planet have been exhausted, and this will be the last period of the Lizardmen's stay here. No one knows whether they can escape from this planet and find a new place to live in the universe."},
		"Comics":           {"collection/cqrcempz1shhrgp2ijtv", "collection/asiinpdtj5xxfvfdwkub", "The rich man Bob ties his life closely to Zecrey, because Zecrey provides him with more conveniences."},
		"CG":               {"collection/owvlmlbzqb7skhh1vai5", "collection/wls6pkmgtgukahtkwnvh", "There is a group of Lizardmen living on an abandoned mine planet. They are divided into three races, namely Zecrey, Komodoensis and Reptoids. The Lizardman is a highly civilized creature in the universe, but according to the current declassified data, no one can know where the Lizardman came from."},
		"Blocks":           {"collection/bxzvucgyeama4vxopjyq", "collection/qizphcugm5rii6spwqox", "Anamite Ore is a rare mineral element found on Planet Lizard-34. Relying on the technology of the Zecrey Tribe, a huge amount of energy can be extracted from the ore. So this ore is also the source of energy for many giant machines."},
		"Ball":             {"collection/v6tayu18nocskfchg04d", "collection/qhnnhopawg8lth0tfc8u", "It is said that the ancestors of Zecrey were trying to use the energy extracted from Anamite Ore to cultivate some organic matter before the Great End War, but this project was aborted with the outbreak of the war."},
		"8 Bits":           {"collection/mlxhzfn1ee5nfvoc2dzt", "collection/uelcwluwxbap4vlggixf", "In the Zecrey world, the Lizardmen will develop Anamite Ore resources on this planet. Each race has deployed its most excellent fighters in face of the imminent battle."},
	}
	for CName, Infos := range collections {
		accountName := "alice"
		seed := "asdfasdfasdf98fd05c70sdafasdfasdffdasdfsadfsdfsfdasdf30383efcb3954631"
		ShortName := CName
		CategoryId := "1"
		LogoImage := Infos[0]
		FeaturedImage := Infos[0]
		BannerImage := Infos[1]
		Description := Infos[2]
		CreatorEarningRate := "200"
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

func TestMintNftInBatch(t *testing.T) {
	nfts := map[string]string{
		"8_Bit_Komo_Shooter":  "collection/mpallit5lf4uk6nevfrd",
		"8_Bit_Komo_Warrior":  "collection/vmji7607tn4pvjtmpxua",
		"8_Bit_Zec_Scientist": "collection/iu3ywffre7inylomvgyc",
		"8_Bit_Zec_Warrior":   "collection/xolwgsun7q1bgpdhffz7",
		"Komo_Student":        "collection/kcfwsruzxeolkhr6ve30",
		"Zec_Student":         "collection/yzxtxfdffavshvytxuxy",
	}
	for nftName, url := range nfts {
		var CollectionId int64 = 7
		accountName := "alice"
		seed := "asdfasdfasdf98fd05c70sdafasdfasdffdasdfsadfsdfsfdasdf30383efcb3954631 "
		c, err := NewClient(accountName, seed)
		if err != nil {
			t.Fatal(err)
		}

		NftUrl := "-"
		Name := fmt.Sprintf("%s", nftName)
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
			CollectionId, NftUrl, Name, Description, Media,
			string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
		if err != nil {
			t.Fatal(err)
		}
		data, err := json.Marshal(ret)
		fmt.Println("MintNft:", string(data))
	}
}

func TestQueryEfficiency_getSellOffers(t *testing.T) {
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  offer(where: {status: {_eq: \"%d\"}, direction: {_eq: \"1\"}}) {\n    id\n    l2_offer_id\n    asset_id\n    counterpart_id\n    payment_asset_id\n    payment_asset_amount\n    signature\n    status\n    direction\n    expired_at\n   created_at\n    asset {\n      id\n      nft_index\n      name\n      collection_id\n      content_hash\n      create_tx_hash\n      creator_earning_rate\n      description\n      expired_at\n      image_thumb\n      l1_token_id\n      last_payment_asset_amount\n      last_payment_asset_id\n      media_detail {\n        url\n      }\n      nft_url\n      status\n      video_thumb\n      asset_stats {\n        max_value\n        key\n      }\n      asset_properties {\n        key\n        value\n      }\n      asset_levels {\n        key\n        max_value\n        value\n      }\n    }\n  }\n}\n","variables":{}}
`, 1)
	vegetaTest("getSellOffers", queryStr)
	TestQueryEfficiency_getAccountAssets(t)
}
func TestQueryEfficiency_getAccountAssets(t *testing.T) {
	accountIndex := 4
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  actionGetAccountAssets(account_index: %d) {\n    confirmedAssetIdList\n    pendingAssets {\n      account_name\n      audio_thumb\n      collection_id\n      content_hash\n      created_at\n      creator_earning_rate\n      description\n      expired_at\n      id\n      image_thumb\n      levels\n      media\n      name\n      nft_index\n      properties\n      stats\n      status\n      video_thumb\n    }\n  }\n}\n","variables":{}}
`, accountIndex)
	vegetaTest("getAccountAssets", queryStr)
	//TestQueryEfficiency_getAccountCollections(t)
}
func TestQueryEfficiency_getAccountCollections(t *testing.T) {
	accountIndex := 4
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  actionGetAccountCollections(account_index: %d) {\n    confirmedCollectionIdList\n    pendingCollections {\n      status\n      account_name\n      banner_image\n      banner_thumb\n      browse_count\n      category_id\n      created_at\n      creator_earning_rate\n      description\n      discord_link\n      expired_at\n      external_link\n      featured_Thumb\n      featured_image\n      floor_price\n      id\n      instagram_link\n      item_count\n      l2_collection_id\n      logo_image\n      logo_thumb\n      name\n      one_day_trade_volume\n      short_name\n      telegram_link\n      total_trade_volume\n      twitter_link\n    }\n  }\n}\n","variables":{}}
`, accountIndex)
	vegetaTest("getAccountCollections", queryStr)
	TestQueryEfficiency_getAccountOffers(t)
}
func TestQueryEfficiency_getAccountOffers(t *testing.T) {
	accountIndex := 4
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  actionGetAccountOffers(account_index: %d) {\n    confirmedOfferIdList\n    pendingOffers {\n      account_name\n      asset_id\n      created_at\n      direction\n      expired_at\n      id\n      payment_asset_amount\n      payment_asset_id\n      signature\n      status\n    }\n  }\n}","variables":{}}
`, accountIndex)
	vegetaTest("getAccountOffers", queryStr)
	TestQueryEfficiency_getAssetByAssetId(t)
}
func TestQueryEfficiency_getAssetByAssetId(t *testing.T) {
	assetId := 3
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  actionGetAssetByAssetId(asset_id: %d) {\n    asset {\n      account_name\n      audio_thumb\n      collection_id\n      content_hash\n      created_at\n      creator_earning_rate\n      description\n      expired_at\n      id\n      image_thumb\n      levels\n      media\n      name\n      nft_index\n      properties\n      stats\n      status\n      video_thumb\n    }\n  }\n}\n","variables":{}}
`, assetId)
	vegetaTest("getAssetByAssetId", queryStr)
	TestQueryEfficiency_getAssetOffers(t)
}
func TestQueryEfficiency_getAssetOffers(t *testing.T) {
	assetId := 3
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  actionGetAssetOffers(asset_id: %d) {\n    confirmedOfferIdList\n    pendingOffers {\n      account_name\n      asset_id\n      created_at\n      direction\n      expired_at\n      id\n      payment_asset_amount\n      payment_asset_id\n      signature\n      status\n    }\n  }\n}","variables":{}}
`, assetId)
	vegetaTest("getAssetOffers", queryStr)
	//TestQueryEfficiency_getCollectionAssets(t)
}
func TestQueryEfficiency_getCollectionAssets(t *testing.T) {
	collectionId := 4
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  actionGetCollectionAssets(collection_id: %d) {\n    ConfirmedCollectionIdList\n    pendingAssets {\n      account_name\n      audio_thumb\n      collection_id\n      content_hash\n      created_at\n      creator_earning_rate\n      description\n      expired_at\n      id\n      image_thumb\n      levels\n      media\n      name\n      nft_index\n      properties\n      stats\n      status\n      video_thumb\n    }\n  }\n}","variables":{}}
`, collectionId)
	vegetaTest("getCollectionAssets", queryStr)
	TestQueryEfficiency_getCollectionById(t)
}
func TestQueryEfficiency_getCollectionById(t *testing.T) {
	collectionId := 4
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  actionGetCollectionById(collection_id: %d) {\n    collection {\n      account_name\n      banner_image\n      banner_thumb\n      browse_count\n      category_id\n      created_at\n      creator_earning_rate\n      description\n      discord_link\n      expired_at\n      external_link\n      featured_Thumb\n      featured_image\n      floor_price\n      id\n      instagram_link\n      item_count\n      l2_collection_id\n      logo_image\n      logo_thumb\n      name\n      one_day_trade_volume\n      short_name\n      status\n      telegram_link\n      total_trade_volume\n      twitter_link\n    }\n  }\n}\n","variables":{}}
`, collectionId)
	vegetaTest("getCollectionById", queryStr)
	TestQueryEfficiency_rankNft(t)
}
func TestQueryEfficiency_rankNft(t *testing.T) {
	assetID := 3
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached ($_in: [bigint!] = %d) {\n  asset(where: {id: {_in: $_in}}) {\n    account {\n      account_index\n      account_name\n    }\n    collection {\n      account_id\n      account_index\n      collection_stats {\n        collection_id\n        day_growth_rate\n        day_trade_volume\n        floor_price\n        item_count\n        last_week_trade_volume\n        month_trade_volume\n        total_trade_volume\n        week_growth_rate\n        week_trade_volume\n        yesterday_trade_volume\n      }\n      collection_url\n      description\n      id\n      l2_collection_id\n    }\n    offers(limit: 1, offset: 0, where: {deleted_at: {_is_null: true}, expired_at: {_gt: \"123465\"}, direction: {_eq: \"1\"}}, order_by: {payment_asset_amount: desc_nulls_last}) {\n      payment_asset_amount\n    }\n  }\n}\n","variables":{}}
`, assetID)
	vegetaTest("rankNft", queryStr)
	TestQueryEfficiency_getTableAccountCollectionsInfo(t)
}
func TestQueryEfficiency_getTableAccountCollectionsInfo(t *testing.T) {
	accountIndex := 3
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached($_in: [bigint!] = %d) {\n  collection(where: {account_index: {_eq: \"3\"}}) {\n    account_id\n    account_index\n    deleted_at\n    description\n    discord_link\n    expired_at\n  }\n}","variables":{}}
`, accountIndex)
	vegetaTest("getTableAccountCollections", queryStr)
	TestQueryEfficiency_hotNft(t)
}
func TestQueryEfficiency_hotNft(t *testing.T) {
	assetID := 3
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached($_in: [bigint!] = %d) {\n  asset(offset: 100, limit: 100) {\n    account {\n      account_index\n      account_name\n      assets_aggregate(where: {status: {_in: \"1\", _nin: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25]}}) {\n        aggregate {\n          count\n        }\n      }\n    }\n    collection {\n      l2_collection_id\n      id\n    }\n    offers_aggregate {\n      aggregate {\n        sum {\n          payment_asset_amount\n        }\n      }\n    }\n    offers(where: {status: {_in: \"1\"}, id: {_nin: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25]}, direction: {_eq: \"1\"}, deleted_at: {_is_null: true}}, order_by: {created_at: desc_nulls_last}) {\n      id\n      l2_offer_id\n      payment_asset_id\n      payment_asset_amount\n      direction\n    }\n  }\n}\n","variables":{}}
`, assetID)
	vegetaTest("hotNft", queryStr)
	TestQueryEfficiency_popularNftCollection(t)
}
func TestQueryEfficiency_popularNftCollection(t *testing.T) {
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  collection(order_by: {collection_stat: {total_trade_volume: desc_nulls_last, day_trade_volume: desc_nulls_last, browse_count: desc_nulls_last, week_trade_volume: desc_nulls_first, month_trade_volume: desc_nulls_last}}, where: {status: {_eq: \"1\"}}, limit: 10, offset: 10) {\n    account {\n      account_index\n      account_name\n    }\n    id\n    name\n    logo_thumb\n    shortname\n    status\n    collection_stats {\n      day_growth_rate\n      day_trade_volume\n      floor_price\n      item_count\n      last_week_trade_volume\n      month_trade_volume\n      total_trade_volume\n      week_growth_rate\n      week_trade_volume\n      yesterday_trade_volume\n    }\n  }\n}\n","variables":{}}
`)
	vegetaTest("popularNftCollection", queryStr)
	TestQueryEfficiency_VolumeIncrease(t)
}
func TestQueryEfficiency_VolumeIncrease(t *testing.T) {
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  collection {\n    collection_stats(order_by: {week_growth_rate: asc}) {\n      browse_count\n      day_growth_rate\n      day_trade_volume\n      floor_price\n      item_count\n      last_week_trade_volume\n      month_trade_volume\n      total_trade_volume\n      week_growth_rate\n      week_trade_volume\n      yesterday_trade_volume\n    }\n    name\n    description\n    deleted_at\n    expired_at\n    discord_link\n    external_link\n    featured_image\n    featured_thumb\n    id\n  }\n}\n","variables":{}}
`)
	vegetaTest("VolumeIncrease(weekly)", queryStr)
	TestQueryEfficiency_RankIng(t)
}
func TestQueryEfficiency_RankIng(t *testing.T) {
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  collection {\n    collection_stats(order_by: {day_trade_volume: asc}) {\n      browse_count\n      day_growth_rate\n      day_trade_volume\n      floor_price\n      item_count\n      last_week_trade_volume\n      month_trade_volume\n      total_trade_volume\n      week_growth_rate\n      week_trade_volume\n      yesterday_trade_volume\n    }\n    name\n    description\n    deleted_at\n    expired_at\n    discord_link\n    external_link\n    featured_image\n    featured_thumb\n    id\n  }\n}\n","variables":{}}
`)
	vegetaTest("RankIng(day_trade_volume)", queryStr)
	TestQueryEfficiency_ReceiveOrSendOffer(t)
}
func TestQueryEfficiency_ReceiveOrSendOffer(t *testing.T) {
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  offer(where: {_or: {account: {account_index: {_eq: \"4\"}}, direction: {_eq: \"0\"}}}, order_by: {created_at: desc_nulls_last}) {\n    account {\n      account_index\n      account_name\n      banner_image\n      updated_at\n      twitter_link\n      status\n      pub_key\n    }\n    asset {\n      account_id\n      collection_id\n      content_hash\n      create_tx_hash\n      collection {\n        account_id\n        account_index\n        banner_thumb\n        category_id\n        banner_image\n      }\n    }\n    counterpart_id\n    created_at\n    direction\n    deleted_at\n    expired_at\n    l2_offer_id\n    id\n  }\n}\n","variables":{}}
`)
	vegetaTest("ReceiveOrSend", queryStr)
	TestQueryEfficiency_CollectionActivity(t)
}
func TestQueryEfficiency_CollectionActivity(t *testing.T) {
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery @cached {\n  activity_tx(where: {nft_index: {_eq: \"3\"}}, order_by: {created_at: desc_nulls_last}) {\n    block_height\n    collection_id\n    created_at\n    deleted_at\n    from_account_index\n    from_account_name\n  }\n}","variables":{}}
`)
	vegetaTest("nftActivity", queryStr)
}

func vegetaTest(apiPath, queryStr string) {
	var metrics vegeta.Metrics
	var Freq = 250
	var duration = 60 * time.Second
	for i := 1; i < 100; i++ {
		var data = []byte(queryStr)
		req, err := http.NewRequest(http.MethodPost, hasuraUrl, bytes.NewReader(data))
		if err != nil {
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-hasura-access-key", hasuraAdminKey)
		rate := vegeta.Rate{Freq: Freq, Per: time.Second}
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: "POST",
			URL:    hasuraUrl,
			Body:   data,
			Header: req.Header,
		})
		attacker := vegeta.NewAttacker()
		for res := range attacker.Attack(targeter, rate, duration, apiPath) {
			metrics.Add(res)
		}
		metrics.Close()
		if metrics.Success < 1 {
			break
		}
		Freq += 100
	}
	fmt.Println(fmt.Sprintf("API Path: %s duration:%d Freq:%d\n Success:%f\n 50th percentile:%s\n 95th percentile:%s\n 99th percentile:%s\n", apiPath, duration, Freq, metrics.Success, metrics.Latencies.P50, metrics.Latencies.P95, metrics.Latencies.P99))
	fmt.Println(metrics.Errors)
}
