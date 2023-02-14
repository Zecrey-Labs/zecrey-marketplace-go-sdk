package ModuleTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"
)

type AcceptOfferRandomOptionParam struct {
	Repeat int
}

type AcceptOfferRandomOption func(t *AcceptOfferRandomOptionParam)

type AcceptOfferProcessor struct {
	Repeat                int
	RandomNextOptionParam AcceptOfferRandomOptionParam
	RandomOptions         []AcceptOfferRandomOption
}

func NewAcceptOfferProcessor(RandomOptions ...AcceptOfferRandomOption) *AcceptOfferProcessor {
	op := AcceptOfferRandomOptionParam{}
	for _, option := range RandomOptions {
		option(&op)
	}
	r := &AcceptOfferProcessor{
		Repeat:                op.Repeat,
		RandomNextOptionParam: op,
		RandomOptions:         RandomOptions,
	}
	return r
}

func (t *AcceptOfferProcessor) Process(ctx *Ctx) error {
	now := time.Now()
	var offer2accept []OfferInfo
	data, err := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s/key%d", OfferDir, ctx.Index))
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &offer2accept)
	if err != nil {
		return err
	}

	res := make([]struct {
		Success bool
		Err     string
	}, t.Repeat)

	for idx := 0; idx < t.Repeat; idx++ {
		params := offer2accept[idx]
		assetAmount, _ := big.NewInt(0).SetString(params.AssetAmount, 10)
		_, err := ctx.Client.AcceptOffer(params.OfferId, false, assetAmount)
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
		err := fmt.Errorf("AcceptOffer failed, failNum=%d,time=%v tx: %v", len(failedTx), time.Now().Sub(now), failedTx)
		return err
	}
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, time.Now().Sub(now)))
	return nil
}
func (c *AcceptOfferProcessor) End() {

}
