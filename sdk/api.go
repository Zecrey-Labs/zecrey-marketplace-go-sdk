package sdk

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	"math/big"
)

type ZecreyNftMarketSDK interface {
	CreateL1Account() (l1Addr, privateKeyStr, l2pk, seed string, err error)

	RegisterAccountWithPrivateKey(accountName, l1Addr, l2pk, privateKey, seed string) (ZecreyNftMarketSDK, error)

	GetAccountByAccountName(accountName string) (*RespGetAccountByAccountName, error)

	ApplyRegisterHost(accountName string, l2Pk string, OwnerAddr string) (*RespApplyRegisterHost, error)

	CreateCollection(ShortName string, CategoryId string, CreatorEarningRate string,
		ops ...model.CollectionOption) (*RespCreateCollection, error)

	GetCollectionById(collectionId int64) (*RespGetCollectionByCollectionId, error)

	GetCollectionsByAccountIndex(AccountIndex int64) (*RespGetAccountCollections, error)

	UpdateCollection(Id string, Name string,
		ops ...model.CollectionOption) (*RespUpdateCollection, error)

	MintNft(
		CollectionId int64,
		NftUrl string, Name string,
		Description string, Media string,
		Properties string, Levels string, Stats string,
	) (*RespCreateAsset, error)

	GetNftByNftId(nftId int64) (*RespetAssetByAssetId, error)

	TransferNft(AssetId int64, toAccountName string) (*ResqSendTransferNft, error)

	WithdrawNft(AssetId int64) (*ResqSendWithdrawNft, error)

	SellNft(AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error)

	BuyNft(AssetId int64, moneyType int64, AssetAmount *big.Int) (*RespListOffer, error)

	AcceptOffer(offerId int64, isSell bool, AssetAmount *big.Int) (*RespAcceptOffer, error)
}

func NewZecreyNftMarketSDK(accountName, seed string) ZecreyNftMarketSDK {
	keyManager, err := NewSeedKeyManager(seed)
	if err != nil {
		panic(fmt.Sprintf("wrong seed:%s", seed))
	}
	connEth, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		panic(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	return &client{
		accountName:    accountName,
		nftMarketURL:   nftMarketUrl,
		legendURL:      legendUrl,
		providerClient: connEth,
		keyManager:     keyManager,
	}
}
