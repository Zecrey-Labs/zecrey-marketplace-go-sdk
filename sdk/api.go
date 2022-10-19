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

	SetKeyManager(keyManager KeyManager)

	// GetTxsListByBlockHeight return txs in block
	GetTxsListByBlockHeight(blockHeight uint32) ([]*Tx, error)

	// GetAccountInfoByAccountName returns account info (mainly pubkey) by using account_name
	GetAccountInfoByAccountName(accountName string) (*AccountInfo, error)

	// GetNextNonce returns nonce of account
	GetNextNonce(accountIdx int64) (int64, error)

	// GetMaxOfferId returns max offer id for an account
	GetMaxOfferId(accountIndex uint32) (uint64, error)

	// GetBlocks return total blocks num and block list
	GetBlocks(offset, limit int64) (uint32, []*Block, error)

	// GetBlockByBlockHeight returns block by height
	GetBlockByBlockHeight(blockHeight int64) (*Block, error)

	GetMempoolTxs(offset, limit uint32) (total uint32, txs []*Tx, err error)

	GetMempoolStatusTxsPending() (total uint32, txs []*Tx, err error)

	GetmempoolTxsByAccountName(accountName string) (total uint32, txs []*Tx, err error)

	GetBalanceByAssetIdAndAccountName(assetId uint32, accountName string) (string, error)

	GetAccountStatusByAccountName(accountName string) (*RespGetAccountStatusByAccountName, error)

	GetAccountStatusByAccountPk(accountPk string) (*RespGetAccountStatusByAccountPk, error)

	GetTxByHash(txHash string) (*RespGetTxByHash, error)

	GetBlockByCommitment(blockCommitment string) (*Block, error)

	GetTxsByPubKey(accountPk string, offset, limit uint32) (total uint32, txs []*Tx, err error)

	GetTxsByAccountName(accountName string, offset, limit uint32) (total uint32, txs []*Tx, err error)

	GetTxsByAccountIndexAndTxType(accountIndex int64, txType, offset, limit uint32) (total uint32, txs []*Tx, err error)

	GetTxsListByAccountIndex(accountIndex int64, offset, limit uint32) (total uint32, txs []*Tx, err error)

	Search(info string) (*RespSearch, error)

	GetAccounts(offset, limit uint32) (*RespGetAccounts, error)

	GetGasFeeAssetList() (*RespGetGasFeeAssetList, error)

	GetWithdrawGasFee(assetId, withdrawAssetId uint32, withdrawAmount uint64) (int64, error)

	GetGasFee(assetId uint32) (int64, error)

	GetCurrencyPrices(symbol string) (*RespGetCurrencyPrices, error)

	GetCurrencyPriceBySymbol(symbol string) (*RespGetCurrencyPriceBySymbol, error)

	GetAssetsList() (*RespGetAssetsList, error)

	GetLayer2BasicInfo() (*RespGetLayer2BasicInfo, error)

	GetAccountInfoByPubKey(accountPk string) (*RespGetAccountInfoByPubKey, error)

	GetAccountInfoByAccountIndex(accountIndex int64) (*RespGetAccountInfoByAccountIndex, error)

	SendMintNftTx(txInfo string) (string, error)

	SendCreateCollectionTx(txInfo string) (int64, error)

	GetCurrentBlockHeight() (int64, error)

	SignAndSendWithdrawNftTx(tx *WithdrawNftTxInfo) (string, error)

	SendWithdrawNftTx(txInfo string) (string, error)

	SignAndSendTransferNftTx(tx *TransferNftTxInfo) (string, error)

	SendTransferNftTx(txInfo string) (string, error)

	SignAndSendAtomicMatchTx(tx *AtomicMatchTxInfo) (string, error)

	SendAtomicMatchTx(txInfo string) (string, error)

	SignAndSendCancelOfferTx(tx *CancelOfferTxInfo) (string, error)

	SendCancelOfferTx(txInfo string) (string, error)

	SignAndSendCreateCollectionTx(tx *CreateCollectionTxInfo) (int64, error)

	SignAndSendMintNftTx(tx *MintNftTxInfo) (string, error)
}

func NewZecreyNftMarketSDK(legendUrl, nftmarketUrl string) ZecreyLegendSDK {
	return &client{
		nftMarketURL: nftmarketUrl,
		legendURL:    legendUrl,
	}
}
