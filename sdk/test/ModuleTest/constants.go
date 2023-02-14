package ModuleTest

import "github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"

const (
	OfferDir       = "offerDir"
	NftDir         = "nftDir" //makeSell transfer withdraw
	Collection2Nft = "collection2Nft"
	Media2Nft      = "media2Nft.json"
)

type TxType int

const (
	TxTypeCreateCollection TxType = sdk.TxTypeCreateCollection // 11
	TxTypeMint             TxType = sdk.TxTypeMintNft          // 12
	TxTypeTransfer         TxType = sdk.TxTypeTransferNft      // 13
	TxTypeMatch            TxType = sdk.TxTypeAtomicMatch      // 14
	TxTypeCancelOffer      TxType = sdk.TxTypeCancelOffer      // 15
	TxTypeWithdrawNft      TxType = sdk.TxTypeWithdrawNft      // 16
	TxTypeListOffer        TxType = sdk.TxTypeOffer            // 19
)
