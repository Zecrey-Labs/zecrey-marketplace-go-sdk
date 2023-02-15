package multiAccountTest

import "github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"

const (
	OfferDir       = "offerDir"
	NftDir         = "nftDir" //makeSell transfer withdraw
	Collection2Nft = "collection2Nft"
	NftTestTmp     = "/Users/user0/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/"
	KeyDir         = "test_account_in_dev_count_1000/"
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
