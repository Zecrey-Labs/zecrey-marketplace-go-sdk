package sdk

type Asset struct {
	Id         uint32
	BalanceEnc string
}

type AccountInfo struct {
	Index     uint32   `json:"account_index"`
	Name      string   `json:"account_name"`
	Nonce     int64    `json:"nonce"`
	AccountPk string   `json:"account_pk"`
	Assets    []*Asset `json:"assets"`
}

type RawTx struct {
	TxType uint32
	TxInfo string //globalrpc => sendAddliquidity.go
	TxHash string
}

type TxHash struct {
	TxHash    string `json:"tx_hash"`
	CreatedAt int64  `json:"created_at"`
}

type TxDetail struct {
	TxId            int64  `json:"tx_id"`
	AssetId         int64  `json:"asset_id"`
	AssetType       int64  `json:"asset_type"`
	AccountIndex    int64  `json:"account_index"`
	AccountName     string `json:"account_name"`
	Balance         string `json:"balance"`
	BalanceDelta    string `json:"balance_delta"`
	Order           int64  `json:"order"`
	AccountOrder    int64  `json:"account_order"`
	Nonce           int64  `json:"nonce"`
	CollectionNonce int64  `json:"collection_nonce"`
}

type Tx struct {
	TxId          int64       `json:"tx_id"`
	TxHash        string      `json:"tx_hash"`
	TxType        int64       `json:"tx_type"`
	GasFee        string      `json:"gas_fee"`
	GasFeeAssetId int64       `json:"gas_fee_asset_id"`
	TxStatus      int64       `json:"tx_status"`
	BlockHeight   int64       `json:"block_height"`
	BlockId       int64       `json:"block_id"`
	StateRoot     string      `json:"state_root"`
	NftIndex      int64       `json:"nft_index"`
	PairIndex     int64       `json:"pair_index"`
	AssetId       int64       `json:"asset_id"`
	TxAmount      string      `json:"tx_amount"`
	NativeAddress string      `json:"native_address"`
	TxInfo        string      `json:"tx_info"`
	TxDetails     []*TxDetail `json:"tx_details"`
	ExtraInfo     string      `json:"extra_info"`
	Memo          string      `json:"memo"`
	AccountIndex  int64       `json:"account_index"`
	Nonce         int64       `json:"nonce"`
	ExpiredAt     int64       `json:"expired_at"`
}

type RespGetTxsListByBlockHeight struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type Block struct {
	BlockCommitment                 string `json:"block_commitment"`
	BlockHeight                     int64  `json:"block_height"`
	StateRoot                       string `json:"state_root"`
	PriorityOperations              int64  `json:"priority_operations"`
	PendingOnChainOperationsHash    string `json:"pending_on_chain_operations_hash"`
	PendingOnChainOperationsPubData string `json:"pending_on_chain_operations_hub_data"`
	CommittedTxHash                 string `json:"committed_tx_hash"`
	CommittedAt                     int64  `json:"committed_at"`
	VerifiedTxHash                  string `json:"verified_tx_hash"`
	VerifiedAt                      int64  `json:"verified_at"`
	Txs                             []*Tx  `json:"txs"`
	BlockStatus                     int64  `json:"block_status"`
}

type RespGetBlocks struct {
	Total  uint32   `json:"total"`
	Blocks []*Block `json:"blocks"`
}

type RespGetBlockByBlockHeight struct {
	Block *Block `json:"block"`
}

type RespGetMaxOfferId struct {
	OfferId uint64 `json:"offer_id"`
}

type RespSendTx struct {
	TxId string `json:"tx_id"`
}

type RespGetNextNonce struct {
	Nonce int64 `json:"nonce"`
}

type RespGetmempoolTxsByAccountName struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"mempool_txs"`
}

type RespGetMempoolTxs struct {
	Total      uint32 `json:"total"`
	MempoolTxs []*Tx  `json:"mempool_txs"`
}

type RespGetTxByHash struct {
	Tx          Tx    `json:"result"`
	CommittedAt int64 `json:"committed_at"`
	VerifiedAt  int64 `json:"verified_at"`
	ExecutedAt  int64 `json:"executed_at"`
	AssetAId    int64 `json:"asset_a_id"`
	AssetBId    int64 `json:"asset_b_id"`
}

type RespGetAccountStatusByAccountPk struct {
	AccountStatus int64  `json:"account_status"`
	AccountIndex  int64  `json:"account_index"`
	AccountName   string `form:"account_name"`
}

type RespGetAccountStatusByAccountName struct {
	AccountStatus uint32 `json:"account_status"`
	AccountIndex  uint32 `json:"account_index"`
	AccountPk     string `json:"account_pk"`
}

type RespGetBalanceInfoByAssetIdAndAccountName struct {
	Balance string `json:"balance_enc"`
}

type RespGetBlockByCommitment struct {
	Block Block `json:"block"`
}

type RespGetLayer2BasicInfo struct {
	BlockCommitted             int64    `json:"block_committed"`
	BlockVerified              int64    `json:"block_verified"`
	TotalTransactions          int64    `json:"total_transactions"`
	TransactionsCountYesterday int64    `json:"transactions_count_yesterday"`
	TransactionsCountToday     int64    `json:"transactions_count_today"`
	DauYesterday               int64    `json:"dau_yesterday"`
	DauToday                   int64    `json:"dau_today"`
	ContractAddresses          []string `json:"contract_addresses"`
}

type RespGetAssetsList struct {
	Assets []*AssetInfo `json:"assets"`
}

type AssetInfo struct {
	AssetId       uint32 `json:"asset_id"`
	AssetName     string `json:"asset_name"`
	AssetDecimals uint32 `json:"asset_decimals"`
	AssetSymbol   string `json:"asset_symbol"`
	AssetAddress  string `json:"asset_address"`
	IsGasAsset    uint32 `json:"is_gas_asset"`
}

type RespGetCurrencyPriceBySymbol struct {
	AssetId uint32 `json:"assetId"`
	Price   uint64 `json:"price"`
}

type RespGetCurrencyPrices struct {
	Data []*DataCurrencyPrices `json:"data"`
}

type DataCurrencyPrices struct {
	Pair    string `json:"pair"`
	AssetId uint32 `json:"assetId"`
	Price   uint64 `json:"price"`
}

type RespGetGasFee struct {
	GasFee int64 `json:"gas_fee"`
}

type RespGetGasFeeAssetList struct {
	Assets []AssetInfo `json:"assets"`
}

type RespGetAccounts struct {
	Total    uint32      `json:"total"`
	Accounts []*Accounts `json:"accounts"`
}

type Accounts struct {
	AccountIndex uint32 `json:"account_index"`
	AccountName  string `json:"account_name"`
	PublicKey    string `json:"public_key"`
}

type RespSearch struct {
	DataType int32 `json:"data_type"`
}

type RespGetTxsListByAccountIndex struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetTxsByAccountIndexAndTxType struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetTxsByAccountName struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetTxsByPubKey struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type RespGetAccountInfoByPubKey struct {
	AccountStatus uint32          `json:"account_status"`
	AccountName   string          `json:"account_name"`
	AccountIndex  int64           `json:"account_index"`
	Assets        []*AccountAsset `json:"assets"`
}

type AccountAsset struct {
	AssetId                  uint32 `json:"asset_id"`
	Balance                  string `json:"balance"`
	LpAmount                 string `json:"lp_amount"`
	OfferCanceledOrFinalized string `json:"offer_canceled_or_finalized"`
}

type RespGetAccountInfoByAccountIndex struct {
	AccountStatus uint32          `json:"account_status"`
	AccountName   string          `json:"account_name"`
	AccountPk     string          `form:"account_pk"`
	Assets        []*AccountAsset `json:"assets"`
}

type RespSendCreateCollectionTx struct {
	CollectionId int64 `json:"collection_id"`
}

type RespSendMintNftTx struct {
	TxId string `json:"tx_id"`
}
type RespCurrentBlockHeight struct {
	Height int64 `json:"height"`
}

type RespSendTransferNftTx struct {
	TxId string `json:"tx_id"`
}

type RespSendWithdrawNftTx struct {
	TxId string `json:"tx_id"`
}

type RespSendCancelOfferTx struct {
	TxId string `json:"tx_id"`
}

type RespSendAtomicMatchTx struct {
	TxId string `json:"tx_id"`
}

//======================================== nft ===============================
type ReqGetStatus struct {
}

type RespGetStatus struct {
	Status        uint32 `json:"status"`
	NetworkId     uint32 `json:"network_id"`
	ServerVersion string `json:"server_version"`
}

type InputBody struct {
	URL string `json:"url"`
}

type ActionBody struct {
	Name string `json:"name,optional"`
}

type SessionVariablesBody struct {
	XHasuraUserId string `json:"x-hasura-user-id,optional"`
	XHasuraRole   string `json:"x-hasura-role,optional"`
}

type ReqGetActionStatus struct {
	Input            InputBody            `json:"input"`
	Action           ActionBody           `json:"action,optional"`
	SessionVariables SessionVariablesBody `json:"session_variables,optional"`
	RequestQuery     string               `json:"request_query"`
}

type RespGetActionStatus struct {
	Status        uint32 `json:"status"`
	NetworkId     uint32 `json:"network_id"`
	ServerVersion string `json:"server_version"`
}

type ReqGetAccountNextNonce struct {
	AccountId int64 `form:"account_id"`
}

type RespGetAccountNextNonce struct {
	Nonce int64 `json:"nonce"`
}

type Categorie struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type ReqGetCollectionCategories struct {
}

type RespGetCollectionCategories struct {
	Categories []*Categorie `json:"categories"`
}

type ReqCheckShortName struct {
	ShortName string `form:"short_name"`
}

type RespCheckShortName struct {
	Valid bool `json:"valid"`
}

type Collection struct {
	Id                 int64  `json:"id"`
	L2CollectionId     int64  `json:"l2_collection_id"`
	AccountName        string `json:"account_name"`
	Name               string `json:"name"`
	ShortName          string `json:"short_name"`
	Description        string `json:"description"`
	CategoryId         int64  `json:"category_id"`
	CollectionUrl      string `json:"collection_url"`
	ExternalLink       string `json:"external_link"`
	TwitterLink        string `json:"twitter_link"`
	InstagramLink      string `json:"instagram_link"`
	TelegramLink       string `json:"telegram_link"`
	DiscordLink        string `json:"discord_link"`
	LogoImage          string `json:"logo_image"`
	LogoThumb          string `json:"logo_thumb"`
	FeaturedImage      string `json:"featured_image"`
	FeaturedThumb      string `json:"featured_Thumb"`
	BannerImage        string `json:"banner_image"`
	BannerThumb        string `json:"banner_thumb"`
	CreatorEarningRate int64  `json:"creator_earning_rate"`
	Status             string `json:"status"`
	ExpiredAt          int64  `json:"expired_at"`
	CreatedAt          int64  `json:"created_at"`
	ItemCount          int64  `json:"item_count"`
	BrowseCount        int64  `json:"browse_count"`
	FloorPrice         int64  `json:"floor_price"`
	OneDayTradeVolume  int64  `json:"one_day_trade_volume"`
	TotalTradeVolume   int64  `json:"total_trade_volume"`
}

type ReqCreateCollection struct {
	ShortName          string `form:"short_name"`
	CategoryId         int64  `form:"category_id"`
	CollectionUrl      string `form:"collection_url"`
	ExternalLink       string `form:"external_link"`
	TwitterLink        string `form:"twitter_link"`
	InstagramLink      string `form:"instagram_link"`
	TelegramLink       string `form:"telegram_link"`
	DiscordLink        string `form:"discord_link"`
	LogoImage          string `form:"logo_image"`
	FeaturedImage      string `form:"featured_image"`
	BannerImage        string `form:"banner_image"`
	CreatorEarningRate int64  `form:"creator_earning_rate"`
	PaymentAssetIds    string `form:"payment_asset_ids"`
	Transaction        string `form:"transaction"`
}

type RespCreateCollection struct {
	Collection Collection `json:"collection"`
}

type ReqUpdateCollection struct {
	Id            int64  `form:"id"`
	AccountName   string `form:"account_name"`
	Name          string `form:"name"`
	CollectionUrl string `form:"collection_url"`
	Description   string `form:"description"`
	CategoryId    int64  `form:"category_id"`
	ExternalLink  string `form:"external_link"`
	TwitterLink   string `form:"twitter_link"`
	InstagramLink string `form:"instagram_link"`
	TelegramLink  string `form:"telegram_link"`
	DiscordLink   string `form:"discord_link"`
	LogoImage     string `form:"logo_image"`
	FeaturedImage string `form:"featured_image"`
	BannerImage   string `form:"banner_image"`
	Timestamp     int64  `form:"timestamp"`
	Signature     string `form:"signature"`
}

type RespUpdateCollection struct {
	Collection Collection `json:"collection"`
}

type ReqGetAllCollectionByCategoryId struct {
	CategoryId int64 `form:"category_id"`
}

type RespGetAllCollectionByCategoryId struct {
	Total       int64        `json:"total"`
	Collections []Collection `json:"collections"`
}

type ReqCategoryById struct {
	CategoryId int64 `form:"category_id"`
}

type RespCategoryById struct {
	Category Categorie `json:"categorie"`
}

type ReqSearchCollection struct {
	AccountNames  []string `form:"account_names"`
	CollectionIds []int64  `form:"Collection_ids"`
	Statuses      []string `form:"statuses"`
	Sort          string   `form:"sort"`
	Offset        int64    `form:"offset"`
	Limit         int64    `form:"limit"`
}

type RespSearchCollection struct {
	Total       int64        `json:"total"`
	Collections []Collection `json:"collections"`
}

type ReqGetContentHash struct {
	AccountName  string `form:"account_name"`
	CollectionId int64  `form:"collection_id"`
	Name         string `form:"name"`
	Properties   string `form:"properties"`
	Levels       string `form:"levels"`
	Stats        string `form:"stats"`
}

type RespGetContentHash struct {
	ContentHash string `json:"content_hash"`
}

type Propertie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Level struct {
	Name     string `json:"name"`
	Value    int64  `json:"value"`
	MaxValue int64  `json:"maxValue"`
}

type Stat struct {
	Name     string `json:"name"`
	Value    int64  `json:"value"`
	MaxValue int64  `json:"maxValue"`
}

type NftInfo struct {
	Id                 int64       `json:"id"`
	AccountName        string      `json:"account_name"`
	NftIndex           int64       `json:"nft_index"`
	CollectionId       int64       `json:"collection_id"`
	CreatorEarningRate int64       `json:"creator_earning_rate"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	Media              string      `json:"media"`
	ImageThumb         string      `json:"image_thumb"`
	VideoThumb         string      `json:"video_thumb"`
	AudioThumb         string      `json:"audio_thumb"`
	Status             string      `json:"status"`
	ContentHash        string      `json:"content_hash"`
	NftUrl             string      `json:"nft_url"`
	ExpiredAt          int64       `json:"expired_at"`
	CreatedAt          int64       `json:"created_at"`
	Properties         []Propertie `json:"properties"`
	Levels             []Level     `json:"levels"`
	Stats              []Stat      `json:"stats"`
}

type ReqCreateAsset struct {
	CollectionId int64  `form:"collection_id"`
	NftUrl       string `form:"nft_url"`
	Name         string `form:"name"`
	Description  string `form:"description"`
	Media        string `form:"media"`
	Properties   string `form:"properties"`
	Levels       string `form:"levels"`
	Stats        string `form:"stats"`
	Transaction  string `form:"transaction"`
}

type RespCreateAsset struct {
	Asset NftInfo `json:"asset"`
}

type ReqSearchAsset struct {
	AccountNames  []string    `json:"account_names"`
	CollectionIds []int64     `json:"Collection_ids"`
	Statuses      []string    `json:"statuses"`
	Keyword       string      `json:"keyword"`
	Properties    []Propertie `json:"properties"`
	Levels        []Level     `json:"levels"`
	Stats         []Stat      `json:"stats"`
	Sort          string      `json:"sort"`
	Offset        int64       `json:"offset"`
	Limit         int64       `json:"limit"`
}

type RespSearchAsset struct {
	Total  int64    `json:"total"`
	Assets []*Asset `json:"assets"`
}

type ReqGetAllAssetByCollectionId struct {
	CollectionId int64 `form:"collection_id"`
}

type RespGetAllAssetByCollectionId struct {
	Total  int64    `json:"total"`
	Assets []*Asset `json:"assets"`
}

type ReqMediaUpload struct {
}

type RespMediaUpload struct {
	PublicId string `json:"public_id"`
	Url      string `json:"url,omitempty"`
}

type ReqGetCollectionOwnerNum struct {
	CollectionId int64 `form:"collection_id"`
}

type ResqGetCollectionOwnerNum struct {
	OwnerNum int64 `json:"owner_num"`
}

type ReqGetActivitytxsByAccountId struct {
	AccountId int64 `form:"account_id"`
}

type ResqGetActivitytxsByAccountId struct {
	Activitytxs []Activitytx `form:"activitytxs"`
}

type Activitytx struct {
	TxType   int64  `form:"tx_type"`
	AssetId  int64  `form:"asset_id"`
	Price    int64  `form:"price"`
	Quantity string `form:"quantity"`
	From     string `form:"from"`
	To       string `form:"to"`
	Time     int64  `form:"time"`
}

type ReqSendTransferNft struct {
	AssetId     int64  `form:"asset_id"`
	Transaction string `form:"transaction"`
}

type ResqSendTransferNft struct {
	Success bool `json:"success"`
}

type ReqSendWithdrawNft struct {
	AssetId     int64  `form:"asset_id"`
	Transaction string `form:"transaction"`
}

type ResqSendWithdrawNft struct {
	Success bool `json:"success"`
}

type ReqGetNextOfferId struct {
	AccountName string `form:"account_name"`
}

type RespGetNextOfferId struct {
	Id int64 `json:"id"`
}

type Offer struct {
	Id                 int64  `json:"id"`
	AccountName        string `json:"account_name"`
	Direction          string `json:"direction"`
	AssetId            int64  `json:"asset_id"`
	PaymentAssetId     int64  `json:"payment_asset_id"`
	PaymentAssetAmount string `json:"payment_asset_amount"`
	Status             string `json:"status"`
	Signature          string `json:"signature"`
	ExpiredAt          int64  `json:"expired_at"`
	CreatedAt          int64  `json:"created_at"`
}

type ReqGetOfferByOfferId struct {
	OfferId int64 `form:"offer_id"`
}

type RespGetOfferByOfferId struct {
	Offer Offer `json:"offer"`
}

type ReqGetOfferByAccountNameAndAssetId struct {
	AccountName string `form:"account_name"`
	AssetId     int64  `form:"asset_id"`
	Transaction string `form:"transaction"`
}

type RespGetOfferByAccountNameAndAssetId struct {
	Offer Offer `json:"offer"`
}

type ReqAcceptOffer struct {
	Id          int64  `form:"id"`
	Transaction string `form:"transaction"`
}

type RespAcceptOffer struct {
	Offer Offer `json:"offer"`
}

type ReqCancelOffer struct {
	Id          int64  `form:"id"`
	Transaction string `form:"transaction"`
}

type RespCancelOffer struct {
	Offer Offer `json:"offer"`
}

type ReqSearchOffer struct {
	AccountNames []string `form:"account_names"`
	AssetIds     []int64  `form:"asset_ids"`
	Statuses     []string `form:"statuses"`
	Directions   []string `form:"directions"`
	Sort         string   `form:"sort"`
	Offset       int64    `form:"offset"`
	Limit        int64    `form:"limit"`
}

type RespSearchOffer struct {
	Offer []Offer `json:"offers"`
}

type ReqListOffer struct {
	AccountName *string `form:"accountName"`
	Transaction *string `form:"transaction"`
}

type RespListOffer struct {
	Offer Offer `json:"offer"`
}

type InputAssetActionBody struct {
	AccountIndex int64 `json:"account_index"`
}

type ReqGetAccountAssets struct {
	Input            InputAssetActionBody `json:"input"`
	Action           ActionBody           `json:"action,optional"`
	SessionVariables SessionVariablesBody `json:"session_variables,optional"`
	RequestQuery     string               `json:"request_query"`
}

type RespGetAccountAssets struct {
	ConfirmedAssetIdList []int64  `json:"confirmedAssetIdList"`
	PendingAssets        []*Asset `json:"pendingAssets"`
}

type InputGetAccountCollectionsActionBody struct {
	AccountIndex int64 `json:"account_index"`
}

type ReqGetAccountCollections struct {
	Input            InputGetAccountCollectionsActionBody `json:"input"`
	Action           ActionBody                           `json:"action,optional"`
	SessionVariables SessionVariablesBody                 `json:"session_variables,optional"`
	RequestQuery     string                               `json:"request_query"`
}

type RespGetAccountCollections struct {
	ConfirmedCollectionIdList []int64      `json:"confirmedCollectionIdList"`
	PendingCollections        []Collection `json:"pendingCollections"`
}

type InputCollectionActionBody struct {
	OpCode       string `json:"opCode"`
	CollectionId int64  `json:"collection_id"`
}

type ReqGetCollectionAction struct {
	Input            InputCollectionActionBody `json:"input"`
	Action           ActionBody                `json:"action,optional"`
	SessionVariables SessionVariablesBody      `json:"session_variables,optional"`
	RequestQuery     string                    `json:"request_query"`
}

type RespGetCollectionAction struct {
	OpCode     string     `json:"opCode"`
	Collection Collection `json:"collection"`
}

type InputCollectionAssetActionBody struct {
	CollectionId int64 `json:"collection_id"`
}

type ReqGetCollectionAssets struct {
	Input            InputCollectionAssetActionBody `json:"input"`
	Action           ActionBody                     `json:"action,optional"`
	SessionVariables SessionVariablesBody           `json:"session_variables,optional"`
	RequestQuery     string                         `json:"request_query"`
}

type ResqGetCollectionAssets struct {
	ConfirmedAssetIdList []int64  `json:"confirmedAssetIdList"`
	PendingAssets        []*Asset `json:"pendingAssets"`
}

type InputGetAccountOffersActionBody struct {
	AccountIndex int64 `json:"account_index"`
}

type ReqGetAccountOffers struct {
	Input            InputGetAccountOffersActionBody `json:"input"`
	Action           ActionBody                      `json:"action,optional"`
	SessionVariables SessionVariablesBody            `json:"session_variables,optional"`
	RequestQuery     string                          `json:"request_query"`
}

type RespGetAccountOffers struct {
	ConfirmedOfferIdList []int64 `json:"confirmedOfferIdList"`
	PendingOffers        []Offer `json:"pendingOffers"`
}

type InputGetAssetOffersActionBody struct {
	AssetId int64 `json:"asset_id"`
}

type ReqGetAssetOffers struct {
	Input            InputGetAssetOffersActionBody `json:"input"`
	Action           ActionBody                    `json:"action,optional"`
	SessionVariables SessionVariablesBody          `json:"session_variables,optional"`
	RequestQuery     string                        `json:"request_query"`
}

type RespGetAssetOffers struct {
	ConfirmedOfferIdList []int64 `json:"confirmedOfferIdList"`
	PendingOffers        []Offer `json:"pendingOffers"`
}

type InputCollectionByIdActionBody struct {
	CollectionId int64 `json:"collection_id"`
}

type ReqGetCollectionById struct {
	Input            InputCollectionByIdActionBody `json:"input"`
	Action           ActionBody                    `json:"action,optional"`
	SessionVariables SessionVariablesBody          `json:"session_variables,optional"`
	RequestQuery     string                        `json:"request_query"`
}

type RespGetCollectionByCollectionId struct {
	Collection Collection `json:"collection"`
}

type InputGetAssetByIdActionBody struct {
	AssetId int64 `json:"asset_id"`
}

type ReqGetAssetById struct {
	Input            InputGetAssetByIdActionBody `json:"input"`
	Action           ActionBody                  `json:"action,optional"`
	SessionVariables SessionVariablesBody        `json:"session_variables,optional"`
	RequestQuery     string                      `json:"request_query"`
}

type RespetAssetByAssetId struct {
	Asset *NftInfo `json:"asset"`
}

type InputOfferActionBody struct {
	OfferId int64 `json:"offer_id"`
}

type ReqGetOfferById struct {
	Input            InputOfferActionBody `json:"input"`
	Action           ActionBody           `json:"action,optional"`
	SessionVariables SessionVariablesBody `json:"session_variables,optional"`
	RequestQuery     string               `json:"request_query"`
}

type RespGetOfferById struct {
	OfferId Offer `json:"offer"`
}
