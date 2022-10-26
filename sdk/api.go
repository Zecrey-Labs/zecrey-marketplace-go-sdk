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

	GetAccountByAccountName(accountName string) (address string, err error)

	ApplyRegisterHost(accountName string, l2Pk string, OwnerAddr string) (*RespApplyRegisterHost, error)

	CreateCollection(accountName string, ShortName string, CategoryId string, CreatorEarningRate string,
		ops ...model.CollectionOption) (*RespCreateCollection, error)

	GetCollectionById(collectionId int64) (*RespGetCollectionByCollectionId, error)

	GetCollectionsByAccountIndex(AccountIndex int64) (*RespGetAccountCollections, error)

	UpdateCollection(Id string, AccountName string, Name string,
		ops ...model.CollectionOption) (*RespUpdateCollection, error)

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

func NewZecreyNftMarketSDK(keyManager KeyManager) ZecreyNftMarketSDK {
	connEth, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		panic(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	return &client{
		nftMarketURL:   nftMarketUrl,
		legendURL:      legendUrl,
		providerClient: connEth,
		keyManager:     keyManager,
	}
}
