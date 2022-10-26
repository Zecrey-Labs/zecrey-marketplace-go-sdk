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

type Categorie struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
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

type RespCreateCollection struct {
	Collection Collection `json:"collection"`
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

type RespCategoryById struct {
	Category Categorie `json:"categorie"`
}

type RespSearchCollection struct {
	Total       int64        `json:"total"`
	Collections []Collection `json:"collections"`
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

type RespCreateAsset struct {
	Asset NftInfo `json:"asset"`
}

type RespSearchAsset struct {
	Total  int64    `json:"total"`
	Assets []*Asset `json:"assets"`
}

type RespGetAllAssetByCollectionId struct {
	Total  int64    `json:"total"`
	Assets []*Asset `json:"assets"`
}

type RespMediaUpload struct {
	PublicId string `json:"public_id"`
	Url      string `json:"url,omitempty"`
}

type ResqGetCollectionOwnerNum struct {
	OwnerNum int64 `json:"owner_num"`
}

type ResqSendTransferNft struct {
	Success bool `json:"success"`
}

type ResqSendWithdrawNft struct {
	Success bool `json:"success"`
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

type RespGetOfferByAccountNameAndAssetId struct {
	Offer Offer `json:"offer"`
}

type RespAcceptOffer struct {
	Offer Offer `json:"offer"`
}

type RespCancelOffer struct {
	Offer Offer `json:"offer"`
}

type RespSearchOffer struct {
	Offer []Offer `json:"offers"`
}

type RespListOffer struct {
	Offer Offer `json:"offer"`
}

type InputAssetActionBody struct {
	AccountIndex int64 `json:"account_index"`
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

type ReqGetPrepareCreateCollectionTxInfo struct {
	AccountName string `form:"account_name"`
}

type ReqGetPrepareTransferNftTxInfo struct {
	AccountName   string `form:"account_name"`
	ToAccountName string `form:"to_account_name"`
	NftIndex      int64  `form:"nft_index"`
}

type ReqGetPrepareWithdrawNftTxInfo struct {
	AccountName string `form:"account_name"`
	NftIndex    int64  `form:"nft_index"`
}

type RespetPreparetxInfo struct {
	TxType    int64  `json:"tx_type"`
	Transtion string `json:"transtion"`
}

type RespApplyRegisterHost struct {
	Ok bool `json:"ok"`
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
