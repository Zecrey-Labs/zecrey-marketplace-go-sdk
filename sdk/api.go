package sdk

import "math/big"

type ZecreyLegendSDK interface {

	//nftmarket

	CreateCollection(
		accountName string, ShortName string, CategoryId string, CollectionUrl string,
		ExternalLink string, TwitterLink string, InstagramLink string, TelegramLink string, DiscordLink string, LogoImage string,
		FeaturedImage string, BannerImage string, CreatorEarningRate string, PaymentAssetIds string) (*RespCreateCollection, error)
	GetCollectionById(collectionId int64) (*RespGetCollectionByCollectionId, error)

	GetCollectionsByAccountIndex(AccountIndex int64) (*RespGetAccountCollections, error)

	UpdateCollection(
		Id string, AccountName string,
		Name string, CollectionUrl string, Description string, CategoryId string,
		ExternalLink string, TwitterLink string, InstagramLink string, TelegramLink string,
		DiscordLink string, LogoImage string, FeaturedImage string, BannerImage string,
	) (*RespUpdateCollection, error)

	MintNft(
		accountName string,
		CollectionId int64,
		NftUrl string, Name string,
		Description string, Media string,
		Properties string, Levels string, Stats string,
	) (*RespCreateAsset, error)

	GetNftByNftId(nftId int64) (*RespetAssetByAssetId, error)

	TransferNft(AssetId int64, accountName string, toAccountName string) (*ResqSendTransferNft, error)

	WithdrawNft(accountName string, AssetId int64) (*ResqSendWithdrawNft, error)

	SellNft(accountName string, AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error)
	BuyNft(accountName string, AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error)
	AcceptOffer(accountName string, offerId int64, isSell bool, AssetAmount *big.Int) (*RespAcceptOffer, error)
}

func NewZecreyNftMarketSDK(legendUrl, nftmarketUrl string, keyManager KeyManager) ZecreyLegendSDK {
	return &client{
		nftMarketURL: nftmarketUrl,
		legendURL:    legendUrl,
		keyManager:   keyManager,
	}
}
