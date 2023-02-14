package ModuleTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type WithdrawNftRandomOptionParam struct {
	Repeat int
}

type WithdrawNftRandomOption func(t *WithdrawNftRandomOptionParam)

type WithdrawProcessor struct {
	Repeat            int
	RandomOptions     []WithdrawNftRandomOption
	RandomNextOptions WithdrawNftRandomOptionParam
}

func NewWithdrawNftProcessor(RandomOptions ...WithdrawNftRandomOption) *WithdrawProcessor {
	op := WithdrawNftRandomOptionParam{}
	for _, option := range RandomOptions {
		option(&op)
	}
	r := &WithdrawProcessor{
		Repeat:            op.Repeat,
		RandomOptions:     RandomOptions,
		RandomNextOptions: op,
	}
	return r
}

func (t *WithdrawProcessor) Process(ctx *Ctx) error {
	option := WithdrawNftRandomOptionParam{}
	for _, op := range t.RandomOptions {
		op(&option)
	}
	var nftinfo []NftInfo
	data, err := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s/key%d", NftDir, ctx.Index))
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
		_, err := ctx.Client.WithdrawNft(nftId, ctx.L1Addr.String())
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
		err := fmt.Errorf("WithdrawNft failed, failNum=%d,time=%v,tx: %v", len(failedTx), Duration, failedTx)
		writeInfo(ctx.Index, fmt.Sprintf("%v", Duration), fmt.Sprintf(" %v", failedTx))
		return err
	}
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, Duration))
	return nil
}
func (c *WithdrawProcessor) End() {

}
