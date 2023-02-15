package multiAccountTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type TransferNftRandomOptionParam struct {
	ToAccountName string
	Repeat        int
}

type TransferNftRandomOption func(t *TransferNftRandomOptionParam)

type TransferProcessor struct {
	Repeat           int
	ToAccountName    string
	RandomOptions    []TransferNftRandomOption
	RandomNextOption TransferNftRandomOptionParam
}

func NewTransferNftProcessor(RandomOptions ...TransferNftRandomOption) *TransferProcessor {
	op := TransferNftRandomOptionParam{}
	for _, option := range RandomOptions {
		option(&op)
	}
	r := &TransferProcessor{
		ToAccountName:    op.ToAccountName,
		Repeat:           op.Repeat,
		RandomOptions:    RandomOptions,
		RandomNextOption: op,
	}
	return r
}

func (t *TransferProcessor) Process(ctx *Ctx) error {
	option := TransferNftRandomOptionParam{
		ToAccountName: "amber1",
	}
	for _, op := range t.RandomOptions {
		op(&option)
	}
	var nftinfo []NftInfo
	data, err := ioutil.ReadFile(fmt.Sprintf("%s%s/key%d", NftTestTmp, NftDir, ctx.Index))
	if err != nil {
		return fmt.Errorf("ignore")
	}

	err = json.Unmarshal(data, &nftinfo)
	if err != nil {
		return err
	}
	res := make([]struct {
		Success bool
		Err     string
	}, t.Repeat)
	now := time.Now()
	for idx := 0; idx < t.Repeat; idx++ {
		nftId := nftinfo[idx].NftId
		_, err := ctx.Client.TransferNft(nftId, option.ToAccountName)
		if err != nil {
			res[idx].Success = false
			res[idx].Err = err.Error()
			continue
		}
		res[idx].Success = true
	}
	Duration := time.Now().Sub(now)
	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}
	if len(failedTx) > 0 {
		err := fmt.Errorf("TransferNft failed, failNum=%d,time=%v tx: %v", len(failedTx), Duration, failedTx)
		writeInfo(ctx.Index, fmt.Sprintf("%v", Duration), fmt.Sprintf(" %v", failedTx))
		return err
	}
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, Duration))
	return nil
}

func (c *TransferProcessor) End() {

}
