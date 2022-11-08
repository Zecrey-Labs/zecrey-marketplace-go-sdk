package sdk

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/zecrey-labs/zecrey-crypto/util/eddsaHelper"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

type ZecreyNftMarketSDK interface {
	GetAccountByAccountName(accountName string) (*RespGetAccountByAccountName, error)

	GetCategories() (*RespGetCollectionCategories, error)

	GetCollectionById(collectionId int64) (*RespGetCollectionByCollectionId, error)

	GetCollectionsByAccountIndex(AccountIndex int64) (*RespGetAccountCollections, error)

	GetAccountNFTs(AccountIndex int64) (*RespGetAccountAssets, error)

	GetAccountOffers(AccountIndex int64) (*RespGetAccountOffers, error)

	GetNftOffers(nftId int64) (*RespGetAssetOffers, error)

	GetNftById(nftId int64) (*RespetAssetByAssetId, error)

	GetOfferById(OfferId int64) (*RespGetOfferByOfferId, error)

	GetMyInfo() (accountName string, l2pk string, seed string)

	ApplyRegisterHost(accountName string, l2Pk string, OwnerAddr string) (*RespApplyRegisterHost, error)

	UploadMedia(filePath string) (*RespMediaUpload, error)

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

	AcceptOffer(offerId int64, isSell bool, AssetAmount *big.Int) (*RespAcceptOffer, error)
}

//NewZecreyNftMarketSDK public
func NewZecreyNftMarketSDK(accountName, seed string) ZecreyNftMarketSDK {
	keyManager, err := NewSeedKeyManager(seed)
	if err != nil {
		panic(fmt.Sprintf("wrong seed:%s", seed))
	}
	l2pk := eddsaHelper.GetEddsaPublicKey(seed[2:])
	connEth, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		panic(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	return &client{
		accountName:    fmt.Sprintf("%s%s", accountName, NameSuffix),
		seed:           seed,
		l2pk:           l2pk,
		nftMarketURL:   nftMarketUrl,
		legendURL:      legendUrl,
		providerClient: connEth,
		keyManager:     keyManager,
	}
}
