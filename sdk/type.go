package sdk

import "io"

type Asset struct {
	AssetId                  uint32 `json:"asset_id"`
	Balance                  string `json:"balance"`
	LpAmount                 string `json:"lp_amount"`
	OfferCanceledOrFinalized string `json:"offer_canceled_or_finalized"`
}

type AccountInfo struct {
	Index     uint32   `json:"account_index"`
	Name      string   `json:"account_name"`
	Nonce     int64    `json:"nonce"`
	AccountPk string   `json:"account_pk"`
	Assets    []*Asset `json:"assets"`
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

type RespGetCollectionOwnerNum struct {
	OwnerNum int64 `json:"owner_num"`
}

type RespSendTransferNft struct {
	Success bool `json:"success"`
}

type RespSendWithdrawNft struct {
	Success bool `json:"success"`
}

type RespSendWithdrawTx struct {
	TxId string `json:"tx_id"`
}

type RespGetNextOfferId struct {
	Id int64 `json:"id"`
}

type Offer struct {
	Id                 int64  `json:"id"`
	L2OfferId          int64  `json:"l2_offer_id"`
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

type RespGetCollectionAction struct {
	OpCode     string     `json:"opCode"`
	Collection Collection `json:"collection"`
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
	NftOfferTreasuryRate       int64    `json:"nft_offer_treasury_rate"`
}

type RespGetAccountByAccountName struct {
	Account NftAccountInfo `json:"account"`
}

type RespGetNextNonce struct {
	Nonce int64 `json:"nonce"`
}

type NftAccountInfo struct {
	Id            int64  `json:"id"`
	AccountIndex  int64  `json:"account_index"`
	AccountPk     string `json:"account_pk"`
	AccountName   string `json:"account_name"`
	Bio           string `json:"bio"`
	Email         string `json:"email"`
	ExternalLink  string `json:"external_link"`
	TwitterLink   string `json:"twitter_link"`
	InstagramLink string `json:"instagram_link"`
	ProfileImage  string `json:"profile_image"`
	ProfileThumb  string `json:"profile_thumb"`
	BannerImage   string `json:"banner_image"`
	BannerThumb   string `json:"banner_thumb"`
	CreatedAt     int64  `json:"created_at"`
}

type ReqMediaUpload struct {
	image io.Writer `form:"image"`
}

type RespGetCollectionCategories struct {
	Categories []*Categorie `json:"categories"`
}

// =========================  hasura struct ============================
type MediaDetail struct {
	Url string `json:"url"`
}

type HauaraNftInfo struct {
	Id                 int64       `json:"id"`
	NftIndex           int64       `json:"nft_index"`
	CollectionId       int64       `json:"collection_id"`
	CreatorEarningRate int64       `json:"creator_earning_rate"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	Media              MediaDetail `json:"media_detail"`
	ImageThumb         string      `json:"image_thumb"`
	VideoThumb         string      `json:"video_thumb"`
	AudioThumb         string      `json:"audio_thumb"`
	Status             int64       `json:"status"`
	ContentHash        string      `json:"content_hash"`
	NftUrl             string      `json:"nft_url"`
	ExpiredAt          int64       `json:"expired_at"`
	CreatedAt          int64       `json:"created_at"`
	Properties         Propertie   `json:"properties"`
	Levels             Level       `json:"levels"`
	Stats              Stat        `json:"stats"`
}

type HasuraOffer struct {
	Id                 int64          `json:"id"`
	L2OfferId          int64          `json:"l2_offer_id"`
	Direction          int            `json:"direction"`
	AssetId            int64          `json:"asset_id"`
	PaymentAssetId     int64          `json:"payment_asset_id"`
	PaymentAssetAmount float64        `json:"payment_asset_amount"`
	Status             int            `json:"status"`
	Signature          string         `json:"signature"`
	ExpiredAt          int64          `json:"expired_at"`
	CreatedAt          string         `json:"created_at"`
	Asset              *HauaraNftInfo `json:"asset"`
}

type HasuraDataOffer struct {
	Offers []*HasuraOffer `json:"offer"`
}

type RespGetListingOffers struct {
	Data *HasuraDataOffer `json:"data"`
}

type HasuraCollectionId struct {
	Id int64 `json:"id"`
}
type HasuraDataCollectionId struct {
	Collection []*HasuraCollectionId `json:"collection"`
}

type RespGetDefaultCollectionId struct {
	Data *HasuraDataCollectionId `json:"data"`
}
type RespGetNftBeingBuy struct {
	Data *HasuraDataOffer `json:"data"`
}

// =======================   sdk ======================================
type ReqGetSdkCreateCollectionTxInfo struct {
	AccountName string `form:"account_name"`
}

type ReqGetSdkMintNftTxInfo struct {
	AccountName  string `form:"account_name"`
	CollectionId int64  `form:"collection_id"`
	TreasuryRate int64  `form:"treasury_rate"`
	Name         string `form:"name"`
	ContentHash  string `form:"content_hash"`
}

type ReqGetSdkTransferNftTxInfo struct {
	AccountName   string `form:"account_name"`
	ToAccountName string `form:"to_account_name"`
	NftId         int64  `form:"nft_id"`
}

type ReqGetSdkAtomicMatchWithTx struct {
	AccountName string `form:"account_name"`
	IsSell      bool   `form:"is_sell"`
	OfferId     int64  `form:"offer_id"`
	MoneyId     int64  `form:"money_id"`
	MoneyAmount string `form:"money_amount"`
}

type ReqGetSdkWithdrawNftTxInfo struct {
	AccountName string `form:"account_name"`
	NftId       int64  `form:"nft_id"`
}

type ReqGetSdkOfferTxInfo struct {
	AccountName string `form:"account_name"`
	NftId       int64  `form:"nft_id"`
	MoneyId     int64  `form:"money_id"`
	MoneyAmount string `form:"money_amount"`
	IsSell      bool   `form:"is_sell"`
}

type ReqGetSdkCancelOfferTxInfo struct {
	AccountName string `form:"account_name"`
	OfferId     int64  `form:"offer_id"`
}

type RespetSdktxInfo struct {
	TxType    int64  `json:"tx_type"`
	Transtion string `json:"transtion"`
}

type ReqGetSdkAccountAssets struct {
	AccountIndex int64 `json:"account_index"`
}

type RespGetSdkAccountAssets struct {
	SdkAssets []*NftInfo `json:"sdkAssets"`
}

type ReqGetSdkAccountCollections struct {
	AccountIndex int64 `json:"account_index"`
}

type RespGetSdkAccountCollections struct {
	SdkCollections []*Collection `json:"sdkCollections"`
}

type ReqGetSdkCollection struct {
	CollectionId int64 `json:"collection_id"`
}

type ReqGetSdkCollectionAssets struct {
	CollectionId int64 `json:"collection_id"`
}

type ResqGetSdkCollectionAssets struct {
	SdkAssets []*NftInfo `json:"sdkAssets"`
}

type ReqGetSdkAccountOffers struct {
	AccountIndex int64 `json:"account_index"`
}

type RespGetSdkAccountOffers struct {
	SdkOffers []*Offer `json:"sdkOffers"`
}

type ReqGetSdkAssetOffers struct {
	AssetId int64 `json:"asset_id"`
}

type RespGetSdkAssetOffers struct {
	SdkOffers []*Offer `json:"sdkOffers"`
}

type ReqGetSdkCollectionById struct {
	CollectionId int64 `json:"collection_id"`
}

type RespGetSdkCollectionById struct {
	Collection Collection `json:"collection"`
}

type ReqGetSdkAssetById struct {
	AssetId int64 `json:"asset_id"`
}

type RespGetSdkAssetById struct {
	Asset *NftInfo `json:"asset"`
}

type ReqGetSdkOffer struct {
	OfferId int64 `json:"offer_id"`
}

type RespGetSdkOfferById struct {
	OfferId Offer `json:"offer"`
}
