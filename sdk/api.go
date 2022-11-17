package sdk

import (
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"math/big"
)

type ZecreyNftMarketSDK interface {
	GetMyInfo() (accountName string, l2pk string, seed string)

	CreateCollection(ShortName string, CategoryId string, CreatorEarningRate string,
		ops ...model.CollectionOption) (*RespCreateCollection, error)

	UpdateCollection(Id string, Name string,
		ops ...model.CollectionOption) (*RespUpdateCollection, error)

	MintNft(
		CollectionId int64,
		NftUrl string, Name string,
		Description string, Media string,
		Properties string, Levels string, Stats string,
	) (*RespCreateAsset, error)

	TransferNft(AssetId int64, toAccountName string) (*ResqSendTransferNft, error)

	WithdrawNft(AssetId int64) (*ResqSendWithdrawNft, error)

	CreateSellOffer(AssetId int64, AssetType int64, AssetAmount *big.Int) (*RespListOffer, error)

	CreateBuyOffer(AssetId int64, AssetType int64, AssetAmount *big.Int) (*RespListOffer, error)

	CancelOffer(offerId int64) (*RespCancelOffer, error)

	AcceptOffer(offerId int64, isSell bool, AssetAmount *big.Int) (*RespAcceptOffer, error)
}
