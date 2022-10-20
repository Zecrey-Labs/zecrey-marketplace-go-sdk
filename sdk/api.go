package sdk

import "math/big"

type ZecreyLegendSDK interface {

	//nftmarket

	CreateCollection(
		accountName string, accountSeed string, ShortName string, CategoryId string, CollectionUrl string,
		ExternalLink string, TwitterLink string, InstagramLink string, TelegramLink string, DiscordLink string, LogoImage string,
		FeaturedImage string, BannerImage string, CreatorEarningRate string, PaymentAssetIds string) (*RespCreateCollection, error)
	GetCollectionById(collectionId int64) (*RespGetCollectionByCollectionId, error)

	GetCollectionsByAccountIndex(AccountIndex int64) (*RespGetAccountCollections, error)

	UpdateCollection(
		Id string, AccountName string, privateKey string,
		Name string, CollectionUrl string, Description string, CategoryId string,
		ExternalLink string, TwitterLink string, InstagramLink string, TelegramLink string,
		DiscordLink string, LogoImage string, FeaturedImage string, BannerImage string,
	) (*RespUpdateCollection, error)

	MintNft(
		privateKey string, accountName string,
		CollectionId int64, l2NftCollectionId int64,
		NftUrl string, Name string,
		Description string, Media string,
		Properties string, Levels string, Stats string,
	) (*RespCreateAsset, error)

	GetNftByNftId(nftId int64) (*RespetAssetByAssetId, error)

	TransferNft(AssetId int64, privateKey string, accountName string, toAccountName string) (*ResqSendTransferNft, error)

	WithdrawNft(privateKey string, accountName string, AssetId int64) (*ResqSendWithdrawNft, error)

	SellNft(privateKey string, accountName string, AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error)
	BuyNft(privateKey string, accountName string, AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error)
	AcceptOffer(privateKey string, accountName string, offerId int64, isSell bool, assetId int64, AssetAmount *big.Int) (*RespAcceptOffer, error)
}

func NewZecreyNftMarketSDK(legendUrl, nftmarketUrl string) ZecreyLegendSDK {
	return &client{
		nftMarketURL: nftmarketUrl,
		legendURL:    legendUrl,
	}
}
