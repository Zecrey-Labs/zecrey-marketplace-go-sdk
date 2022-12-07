package main

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

/*
upload media
from the user`s perspective
Explain clearly

1.Create an account in the contract，Please ensure that your L1 address has enough tokens to pay the GasFee.
2.Synchronize the registered account to the legend
3.Create nft in your collection. By default, you will have a collection that belongs to you. CollectionId=0
4.You can also create your own collection:
	4.1 Get all category
	4.2 Select a category to create your own collection
	4.3 Check all your collections
5.Create an offer to sell your nft
6.Transfer your nft
7.Withdraw your nft
*/
func main1() {

	//1.Create  account in the contract，Please ensure that your L1 address has enough tokens to pay the GasFee.
	//A account
	var _sdkA *sdk.Client
	var _sdkB *sdk.Client
	l1Addr := "0xD207262DEA01aE806fA2dCaEdd489Bd2f5FABcFE"
	var accountInfo *sdk.RespGetAccountByAccountName
	{
		accountName := "A_account" //set your account a name
		privateKey := "1a061a8e74cee1ce2e2ddd29f5afea99ecfbaf1998b6d349a8c09a368e637b8e"
		_, err := sdk.RegisterAccountWithPrivateKey(accountName, l1Addr, privateKey)
		if err != nil {
			panic(err)
		}
		accountInfo, err = sdk.GetAccountByAccountName(accountName)
		if err != nil {
			panic(err)
		}
	}
	//B account
	{
		accountName := "B_Account" //set your account a name
		privateKey := "1a061a8e74cee1ce2e2ddd29f5afea99ecfbaf1998b6d349a8c09a368e637b8e"
		l2Addr := "0xD207262DEA01aE806fA2dCaEdd489Bd2f5FABcFE"
		//Get legend seed and l2Pk
		sdkB, err := sdk.RegisterAccountWithPrivateKey(accountName, l2Addr, privateKey)
		if err != nil {
			panic(err)
		}
		_sdkB = sdkB
	}

	//2.Synchronize the registered account to the legend
	if accountInfo == nil {
		accountName, l2Pk, _ := _sdkA.GetMyInfo()
		ok, err := sdk.ApplyRegisterHost(accountName, l2Pk, l1Addr)
		if err != nil {
			panic(err)
		}
		log.Info("ApplyRegisterHost:", ok)
	}

	//3.Create nft in your collection. By default, you will have a collection that belongs to you. CollectionId=0
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
	Description := "Description information"
	CreatorEarningRate := "6666"
	PaymentAssetIds := "[]"

	retCollection, err := _sdkA.CreateCollection(ShortName, CategoryId, CreatorEarningRate,
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
		panic(err)
	}
	log.Info("create collection Info:", retCollection)
	accountName, _, _ := _sdkA.GetMyInfo()
	var CollectionId int64 = 54
	NftUrl := "-"
	Name := "-"
	DescriptionNft := "nft-sdk-Description"
	Media := "collection/cbenqstwzx5uy9oedjrb"
	key := fmt.Sprintf("zw:%s:%d", accountName, 2)
	value := "red1"
	assetProperty := sdk.Propertie{
		Name:  key,
		Value: value,
	}
	assetLevel := sdk.Level{
		Name:     "assetLevel",
		Value:    123,
		MaxValue: 123,
	}
	assetStats := sdk.Stat{
		Name:     "StatType",
		Value:    456,
		MaxValue: 456,
	}
	// get content hash

	_Properties := []sdk.Propertie{assetProperty}
	_Levels := []sdk.Level{assetLevel}
	_Stats := []sdk.Stat{assetStats}

	_PropertiesByte, err := json.Marshal(_Properties)
	_LevelsByte, err := json.Marshal(_Levels)
	_StatsByte, err := json.Marshal(_Stats)

	retNft, err := _sdkA.MintNft(
		CollectionId,
		NftUrl, Name,
		DescriptionNft, Media,
		string(_PropertiesByte), string(_LevelsByte), string(_StatsByte))
	log.Info("create nft Info:", retNft)

	//4.Create an offer to sell your nft
	var assetId int64 = 1
	var assetType int64 = 0
	assetAmount := big.NewInt(1000000) //
	retListOffer, err := _sdkB.CreateBuyOffer(assetId, assetType, assetAmount)

	//5.Accept offer
	retAccept, err := _sdkA.AcceptOffer(retListOffer.Offer.Id, true, assetAmount)
	log.Info("accept Offer", retAccept)

	//5.1 check nft
	retGetNftByNftId, err := sdk.GetNftById(retAccept.Offer.AssetId)
	if err != nil {
		panic(err)
	}
	log.Info("retGetNftByNftId", retGetNftByNftId.Asset.AccountName)

	////6.Transfer your nft
	//ResqSendTransferNft, err := _sdkA.TransferNft(assetId, "B_Account")
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("retGetNftByNftId", ResqSendTransferNft.Success)
	//
	////7.Withdraw your nft
	//ResqSendWithdrawNft, err := _sdkA.WithdrawNft(assetId)
	//if err != nil {
	//	panic(err)
	//}
	//log.Info("ResqSendWithdrawNft", ResqSendWithdrawNft.Success)

}
