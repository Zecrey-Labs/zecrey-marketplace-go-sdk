package sdk

type ZecreyLegendSDK interface {
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

func NewZecreyNftMarketSDK(url string) ZecreyLegendSDK {
	return &client{
		zkbasURL: url,
	}
}
