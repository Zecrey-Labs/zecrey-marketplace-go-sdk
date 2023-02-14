package ModuleTest

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
	now := time.Now()
	option := TransferNftRandomOptionParam{
		ToAccountName: "sher",
	}
	for _, op := range t.RandomOptions {
		op(&option)
	}
	var nftinfo []NftInfo
	data, err := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s/key%d", NftDir, ctx.Index))
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &nftinfo)
	if err != nil {
		return err
	}
	res := make([]struct {
		Success bool
		Err     string
	}, t.Repeat)

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
	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}
	if len(failedTx) > 0 {
		err := fmt.Errorf("TransferNft failed, failNum=%d,time=%v tx: %v", len(failedTx), time.Now().Sub(now), failedTx)
		return err
	}
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, time.Now().Sub(now)))
	return nil
}
func (c *TransferProcessor) End() {

}
