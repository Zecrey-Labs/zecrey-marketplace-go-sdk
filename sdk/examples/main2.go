package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
)

func main() {
	c, err := sdk.NewClient("sher.zec", "28e1a3762ff9944e9a4ad79477b756ef0aff3d2af76f0f40a0c3ec6ca76cf24b")
	if err != nil {
		fmt.Println("NewClient err:", err)
		return
	}
	filesName := []string{"./eightBits"}
	for _, collectionName := range filesName {
		// 1. creat collection
		collectionLogoFile, err := sdk.UploadMedia(collectionName + "/collection_icon.png")
		if err != nil {
			fmt.Println("UploadMedia err:", err)
			return
		}
		collectionBannerFile, err := sdk.UploadMedia(collectionName + "/collection_cover.png")
		if err != nil {
			fmt.Println("UploadMedia err:", err)
			return
		}
		collection := sdk.Colletcion{
			ShortName:          fmt.Sprintf("%v_%d", collectionName, time.Now().Second()),
			CategoryId:         "1",
			CreatorEarningRate: "1",
			LogoImage:          collectionLogoFile.PublicId,
			BannerImage:        collectionBannerFile.PublicId,
			PaymentAssetIds:    "[]",
			CollectionUrl:      "_",
			ExternalLink:       "_",
			TwitterLink:        "_",
			InstagramLink:      "_",
			TelegramLink:       "_",
			DiscordLink:        "_",
			FeaturedImage:      "_",
			Description:        "_",
		}
		collectionResp, err := c.CreateCollection(collection)
		if err != nil {
			fmt.Println("CreateCollection err:", err)
		}

		// 2. mint nft
		rd, err := ioutil.ReadDir(collectionName + "/nft")
		if err != nil {
			fmt.Println("read dir fail:", err)
			return
		}
		for _, fi := range rd {
			if !fi.IsDir() {
				nftfileName := collectionName + "/nft/" + fi.Name()
				nftImageResp, err := sdk.UploadMedia(nftfileName)
				if err != nil {
					fmt.Println("UploadMedia err:", err)
					return
				}
				_Properties := []sdk.Propertie{sdk.Propertie{
					Name:  fmt.Sprintf("zw:%s:%d", c.AccountName, 2),
					Value: "red1",
				}}
				_Levels := []sdk.Level{sdk.Level{
					Name:     "assetLevel",
					Value:    123,
					MaxValue: 123,
				}}
				_Stats := []sdk.Stat{sdk.Stat{
					Name:     "StatType",
					Value:    456,
					MaxValue: 456,
				}}
				_PropertiesByte, _ := json.Marshal(_Properties)
				_LevelsByte, _ := json.Marshal(_Levels)
				_StatsByte, _ := json.Marshal(_Stats)
				nftInfo := sdk.Mintnft{
					CollectionId: collectionResp.Collection.Id,
					NftUrl:       "",
					Name:         nftfileName[:len(nftfileName)-4],
					TreasuryRate: 30,
					Description:  "",
					Media:        nftImageResp.PublicId,
					Properties:   string(_PropertiesByte),
					Levels:       string(_LevelsByte),
					Stats:        string(_StatsByte),
				}
				resp, err := c.MintNft(nftInfo)
				if err != nil {
					fmt.Println("MintNft err:", err)
				}
				fmt.Println("MintNft resp:", resp)
			}
		}
	}
}
