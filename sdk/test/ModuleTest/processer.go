package ModuleTest

type Processor interface {
	Process(tx *Ctx) error
	End()
}
type Processors struct {
	processorsMap map[TxType]Processor
}

func GetProcessors() *Processors {
	processors := &Processors{make(map[TxType]Processor)}
	processors.processorsMap[TxTypeCreateCollection] = NewCreateCollectionProcessor(func(t *RandomOptionParam) { t.Repeat = 1 })
	processors.processorsMap[TxTypeMint] = NewMintNftProcessor(func(t *NftRandomOptionParam) { t.Repeat = 1 })
	processors.processorsMap[TxTypeTransfer] = NewTransferNftProcessor(func(t *TransferNftRandomOptionParam) {
		t.Repeat = 1
		t.ToAccountName = "sher"
	})
	processors.processorsMap[TxTypeMatch] = NewAcceptOfferProcessor(func(t *AcceptOfferRandomOptionParam) {
		t.Repeat = 1
	})
	processors.processorsMap[TxTypeCancelOffer] = NewCancelOfferProcessor(func(t *CancelOfferRandomOptionParam) { t.Repeat = 1 })
	processors.processorsMap[TxTypeWithdrawNft] = NewWithdrawNftProcessor(func(t *WithdrawNftRandomOptionParam) { t.Repeat = 1 })
	processors.processorsMap[TxTypeListOffer] = NewlistOfferProcessor(func(t *ListOfferRandomOptionParam) { t.Repeat = 1 })
	return processors
}
