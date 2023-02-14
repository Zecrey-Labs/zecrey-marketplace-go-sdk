package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/ModuleTest"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/cancelOffer"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/createCollection"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/listOffer"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/mintNft"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/transferNft"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/util"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/withdrawNft"
	"github.com/ethereum/go-ethereum/common"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	curve "github.com/zecrey-labs/zecrey-crypto/ecc/ztwistededwards/tebn254"
	"github.com/zecrey-labs/zecrey-crypto/util/ecdsaHelper"
	legendSdk "github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"
	"github.com/zeromicro/go-zero/core/conf"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

var configFile = flag.String("f",
	"./config.yaml", "the config file")

var cfg Config
var client *sdk.Client

func testAll() {
	conf.MustLoad(*configFile, &cfg)
	_client, err := sdk.NewClient(cfg.AccountName, cfg.Seed)
	client = _client
	if err != nil {
		panic(err)
	}

	for i := 1; i < 30; i++ {
		createCollectionCorrectBatch(i)
		createCollectionWrongBatch(i)
		mintNftCorrectOnce(i)
		mintNftCorrectWrongBatch(i)
		makeOfferCorrectBatch(i)
		makeOfferWrongBatch(i)
		transferNftCorrectOnce(i)
		transferNftWrongBatch(i)
		withdrawNftCorrectOnce(i)
		withdrawNftWrongBatch(i)
		acceptOfferWrongBatch(i)
		time.Sleep(60 * time.Second)
	}
	time.Sleep(10 * time.Minute)
	panic("==== test over !!!")
}

func testCreateCollection(accountNum int) {
	var legendClients []legendSdk.ZecreyLegendSDK
	wg := sync.WaitGroup{}
	wg.Add(accountNum)
	now := time.Now()
	for index := 0; index < accountNum; index++ {
		privateKey, err := ethercrypto.LoadECDSA(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/test_account_in_dev_count_1000/%s", fmt.Sprintf("key%d", index)))
		privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
		l1Addr, err := ecdsaHelper.GenerateL1Address(privateKey)
		_, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
		legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
		sk, err := curve.GenerateEddsaPrivateKey(seed)
		res, err := legendClient.GetAccountInfoByPubKey(hex.EncodeToString(sk.PublicKey.Bytes()))
		if err != nil {
			fmt.Printf("NewZecreyLegendSDK failed:%v", err)
			return
		} else {
			legendClients = append(legendClients, legendClient)
		}
		CollectionUrl := "https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg"
		ExternalLink := "https://weibo.com/alice"
		TwitterLink := "https://twitter.com/alice"
		InstagramLink := "https://www.instagram.com/alice/"
		TelegramLink := "https://tgstat.com/channel/@alice"
		DiscordLink := "https://discord.com/api/v10/applications/<aliceid>/commands"
		LogoImage := "collection/j9w9z4dmcd2beufvkxkp"
		FeaturedImage := "collection/j9w9z4dmcd2beufvkxkp"
		BannerImage := "collection/j9w9z4dmcd2beufvkxkp"
		Description := "Description information"

		go func(res *legendSdk.RespGetAccountInfoByPubKey, seed, l1Addr string, index int) {
			defer wg.Done()
			client, err := sdk.NewClientNoSuffix(res.AccountName, seed)
			if err != nil {
				panic(err)
			}
			createCollectionClient := createCollection.InitCtx(client, common.HexToAddress(l1Addr))
			now := time.Now()
			count := 1
			errInfo := createCollectionClient.CreateCollectionTest(count, index, func(t *createCollection.RandomOptionParam) {
				t.CategoryId = 1
				t.CreatorEarningRate = 200
				t.RandomShortName = true
				t.Ops = []model.CollectionOption{model.WithCollectionUrl(CollectionUrl),
					model.WithExternalLink(ExternalLink),
					model.WithTwitterLink(TwitterLink),
					model.WithInstagramLink(InstagramLink),
					model.WithTelegramLink(TelegramLink),
					model.WithDiscordLink(DiscordLink),
					model.WithLogoImage(LogoImage),
					model.WithFeaturedImage(FeaturedImage),
					model.WithBannerImage(BannerImage),
					model.WithDescription(Description)}
			})
			fmt.Println(fmt.Sprintf("index=%d  sendCount=%d time=%v errs=%v", index, count, time.Now().Sub(now), errInfo))
		}(res, seed, l1Addr, index)
	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("==== test over all time=%v ", time.Now().Sub(now)))
}

func testUpdateCollection(repeat int) {
	var legendClients []legendSdk.ZecreyLegendSDK
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	for index := 0; index < repeat; index++ {
		privateKey, err := ethercrypto.LoadECDSA(filepath.Join("./sdk/test", util.DefaultDir, util.KeyDir, fmt.Sprintf("key%d", index)))
		privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
		l2pk, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
		legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
		res, err := legendClient.GetAccountInfoByPubKey(l2pk)
		if err != nil {
			fmt.Printf("NewZecreyLegendSDK failed:%v", err)
			panic(fmt.Sprintf("NewZecreyLegendSDK failed:%v", err))
		} else {
			fmt.Printf("AccountName:%v", res.AccountName)
			legendClients = append(legendClients, legendClient)
		}
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

		go func(res *legendSdk.RespGetAccountInfoByPubKey, seed string, index int) {
			defer wg.Done()
			client, err := sdk.NewClient(res.AccountName, seed)
			l1addr, err := sdk.GetAccountL1Address(res.AccountName)
			mintNftClient := createCollection.InitCtx(client, l1addr)
			data, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/medias%d.json", index))
			var Medias []string
			json.Unmarshal(data, &Medias)
			fmt.Println("counts:", len(Medias))
			now := time.Now()
			err = mintNftClient.CreateCollectionTest(len(Medias), index, func(t *createCollection.RandomOptionParam) {
				t.CategoryId = 1
				t.CreatorEarningRate = 200
				t.RandomShortName = true
				t.Ops = []model.CollectionOption{
					model.WithCollectionUrl(CollectionUrl),
					model.WithExternalLink(ExternalLink),
					model.WithTwitterLink(TwitterLink),
					model.WithInstagramLink(InstagramLink),
					model.WithTelegramLink(TelegramLink),
					model.WithDiscordLink(DiscordLink),
					model.WithLogoImage(LogoImage),
					model.WithFeaturedImage(FeaturedImage),
					model.WithBannerImage(BannerImage),
					model.WithDescription(Description)}
			})
			fmt.Println(fmt.Sprintf("==== test over time=%v err=%s", time.Now().Sub(now), err.Error()))
		}(res, seed, index)
	}
	wg.Wait()
}
func testMintNftPer() {
	client, _ := sdk.NewClient("amber1", "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f")
	mintNftClient := mintNft.InitCtx(client, common.HexToAddress("0x09E45d6FcF322c4D93E6aFE7076601FF10BA942E"))
	data, _ := ioutil.ReadFile("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/medias.json")
	var Medias []string
	json.Unmarshal(data, &Medias)
	fmt.Println("counts:", len(Medias))
	now := time.Now()
	err := mintNftClient.MitNftTest(len(Medias), 0, func(t *mintNft.RandomOptionParam) {
		t.CollectionId = 6
		t.RandomNftUrl = true
		t.RandomName = true
		t.RandomDescription = true
		t.Properties = "[]"
		t.Levels = "[]"
		t.Stats = "[]"
		t.Medias = Medias
	})
	fmt.Println(fmt.Sprintf("==== test over time=%v err=%s", time.Now().Sub(now), err.Error()))
}
func testMintNft(repeat int) {
	collectionIds := []int64{1193, 1192, 1194, 1195, 1196, 1197, 1201, 1360}
	var legendClients []legendSdk.ZecreyLegendSDK
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	for index := 0; index < repeat; index++ {
		privateKey, _ := ethercrypto.LoadECDSA(filepath.Join("./sdk/test", util.DefaultDir, util.KeyDir, fmt.Sprintf("key%d", index)))
		privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
		_, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
		legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
		sk, err := curve.GenerateEddsaPrivateKey(seed)
		res, err := legendClient.GetAccountInfoByPubKey(hex.EncodeToString(sk.PublicKey.Bytes()))
		if err != nil {
			fmt.Printf("NewZecreyLegendSDK failed:%v", err)
		} else {
			fmt.Printf("AccountName:%v", res.AccountName)
			legendClients = append(legendClients, legendClient)
		}

		go func(res *legendSdk.RespGetAccountInfoByPubKey, seed string, index int) {
			defer wg.Done()
			client, err := sdk.NewClientNoSuffix(res.AccountName, seed)
			if err != nil {
				panic(err)
			}
			l1Addr, err := ecdsaHelper.GenerateL1Address(privateKey)
			if err != nil {
				panic(err)
			}
			mintNftClient := mintNft.InitCtx(client, common.HexToAddress(l1Addr))
			data, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias_dev/medias%d.json", index))
			var Medias []string
			json.Unmarshal(data, &Medias)
			now := time.Now()
			errInfo := mintNftClient.MitNftTest(len(Medias), index, func(t *mintNft.RandomOptionParam) {
				t.CollectionId = collectionIds[index]
				t.RandomNftUrl = true
				t.RandomName = true
				t.RandomDescription = true
				t.Properties = "[]"
				t.Levels = "[]"
				t.Stats = "[]"
				t.Medias = Medias
			})
			fmt.Println(fmt.Sprintf("index=%d  send time=%v allTimes=%v errs=%v", index, len(Medias), time.Now().Sub(now), errInfo))
		}(res, seed, index)
	}
	wg.Wait()
}
func testListOffer(repeat int) {
	client, _ := sdk.NewClient("amber1", "ee823a72698fd05c70fbdf36ba2ea467d33cf628c94ef030383efcb39581e43f")
	listOfferClient := listOffer.InitCtx(client, common.HexToAddress("0x09E45d6FcF322c4D93E6aFE7076601FF10BA942E"))
	data, _ := ioutil.ReadFile("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/medias.json")
	var Medias []string
	json.Unmarshal(data, &Medias)
	fmt.Println("counts:", len(Medias))
	now := time.Now()
	err := listOfferClient.ListOfferTest()
	fmt.Println(fmt.Sprintf("==== test over time=%v err=%s", time.Now().Sub(now), err.Error()))
}
func testCancelOffer(repeat int) {
	var legendClients []legendSdk.ZecreyLegendSDK
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	for index := 0; index < repeat; index++ {
		privateKey, err := ethercrypto.LoadECDSA(filepath.Join("./sdk/test", util.DefaultDir, util.KeyDir, fmt.Sprintf("key%d", index)))
		privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
		l2pk, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
		legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
		res, err := legendClient.GetAccountInfoByPubKey(l2pk)
		if err != nil {
			fmt.Printf("NewZecreyLegendSDK failed:%v", err)
			panic(fmt.Sprintf("NewZecreyLegendSDK failed:%v", err))
		} else {
			fmt.Printf("AccountName:%v", res.AccountName)
			legendClients = append(legendClients, legendClient)
		}

		go func(res *legendSdk.RespGetAccountInfoByPubKey, seed string, index int) {
			defer wg.Done()
			client, err := sdk.NewClient(res.AccountName, seed)
			l1addr, err := sdk.GetAccountL1Address(res.AccountName)
			cancelOfferClient := cancelOffer.InitCtx(client, l1addr)
			data, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/medias%d.json", index))
			var Medias []string
			json.Unmarshal(data, &Medias)
			fmt.Println("counts:", len(Medias))
			now := time.Now()
			err = cancelOfferClient.CancelOfferTest()
			fmt.Println(fmt.Sprintf("==== test over time=%v err=%s", time.Now().Sub(now), err.Error()))
		}(res, seed, index)
	}
	wg.Wait()

}
func testAcceptOffer(repeat int) {
	var legendClients []legendSdk.ZecreyLegendSDK
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	for index := 0; index < repeat; index++ {
		privateKey, err := ethercrypto.LoadECDSA(filepath.Join("./sdk/test", util.DefaultDir, util.KeyDir, fmt.Sprintf("key%d", index)))
		privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
		l2pk, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
		legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
		res, err := legendClient.GetAccountInfoByPubKey(l2pk)
		if err != nil {
			fmt.Printf("NewZecreyLegendSDK failed:%v", err)
			panic(fmt.Sprintf("NewZecreyLegendSDK failed:%v", err))
		} else {
			fmt.Printf("AccountName:%v", res.AccountName)
			legendClients = append(legendClients, legendClient)
		}

		go func(res *legendSdk.RespGetAccountInfoByPubKey, seed string, index int) {
			defer wg.Done()
			client, err := sdk.NewClient(res.AccountName, seed)
			l1addr, err := sdk.GetAccountL1Address(res.AccountName)
			cancelOfferClient := cancelOffer.InitCtx(client, l1addr)
			data, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/medias%d.json", index))
			var Medias []string
			json.Unmarshal(data, &Medias)
			fmt.Println("counts:", len(Medias))
			now := time.Now()
			err = cancelOfferClient.CancelOfferTest()
			fmt.Println(fmt.Sprintf("==== test over time=%v err=%s", time.Now().Sub(now), err.Error()))
		}(res, seed, index)
	}
	wg.Wait()
}
func testTransferNft(repeat int) {
	var legendClients []legendSdk.ZecreyLegendSDK
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	for index := 0; index < repeat; index++ {
		privateKey, _ := ethercrypto.LoadECDSA(filepath.Join(".", util.DefaultDir, util.KeyDir, fmt.Sprintf("%d", index)))
		privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
		l2pk, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
		legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
		res, err := legendClient.GetAccountInfoByPubKey(l2pk)
		if err != nil {
			fmt.Printf("NewZecreyLegendSDK failed:%v", err)
		} else {
			fmt.Printf("AccountName:%v", res.AccountName)
			legendClients = append(legendClients, legendClient)
		}

		go func(res *legendSdk.RespGetAccountInfoByPubKey, seed string, index int) {
			defer wg.Done()
			client, err := sdk.NewClient(res.AccountName, seed)
			l1addr, err := sdk.GetAccountL1Address(res.AccountName)
			transferNftClient := transferNft.InitCtx(client, l1addr)
			data, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/medias%d.json", index))
			var Medias []string
			json.Unmarshal(data, &Medias)
			fmt.Println("counts:", len(Medias))
			now := time.Now()
			err = transferNftClient.TransferNftTest(func(t *transferNft.RandomOptionParam) {

			})
			fmt.Println(fmt.Sprintf("==== test over time=%v err=%s", time.Now().Sub(now), err.Error()))
		}(res, seed, index)

	}
	wg.Wait()
}
func testWithdrawNft(repeat int) {
	var legendClients []legendSdk.ZecreyLegendSDK
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	for index := 0; index < repeat; index++ {
		privateKey, _ := ethercrypto.LoadECDSA(filepath.Join(".", util.DefaultDir, util.KeyDir, fmt.Sprintf("%d", index)))
		privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
		l2pk, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
		legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
		res, err := legendClient.GetAccountInfoByPubKey(l2pk)
		if err != nil {
			fmt.Printf("NewZecreyLegendSDK failed:%v", err)
		} else {
			fmt.Printf("AccountName:%v", res.AccountName)
			legendClients = append(legendClients, legendClient)
		}

		go func(res *legendSdk.RespGetAccountInfoByPubKey, seed string, index int) {
			defer wg.Done()
			client, err := sdk.NewClient(res.AccountName, seed)
			l1addr, err := sdk.GetAccountL1Address(res.AccountName)
			withdrawNftClient := withdrawNft.InitCtx(client, l1addr)
			data, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/medias%d.json", index))
			var Medias []string
			json.Unmarshal(data, &Medias)
			fmt.Println("counts:", len(Medias))
			now := time.Now()
			err = withdrawNftClient.WithdrawNftTest(func(t *withdrawNft.RandomOptionParam) {

			})
			fmt.Println(fmt.Sprintf("==== test over time=%v err=%s", time.Now().Sub(now), err.Error()))
		}(res, seed, index)

	}
	wg.Wait()
}

//func main() {
//	//testMintNft(7)
//	testCreateCollection(1)
//}

func main() {
	ModuleTest.StartTest(1, ModuleTest.TxTypeMatch)
}
