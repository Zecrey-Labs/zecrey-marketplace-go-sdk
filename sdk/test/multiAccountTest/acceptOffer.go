package multiAccountTest

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
	var offer2accept []OfferInfo
	data, err := ioutil.ReadFile(fmt.Sprintf("%s%s/key%d", NftTestTmp, OfferDir, ctx.Index))
	if err != nil {
		return fmt.Errorf("ignore")
	}
	err = json.Unmarshal(data, &offer2accept)
	if err != nil {
		return err
	}

	res := make([]struct {
		Success bool
		Err     string
	}, t.Repeat)
	now := time.Now()
	for idx := 0; idx < t.Repeat; idx++ {
		params := offer2accept[idx]
		n := new(big.Int)
		n, ok := n.SetString(params.AssetAmount, 10)
		if !ok {
			return fmt.Errorf("SetString: error params.AssetAmount=%s", params.AssetAmount)
		}
		fmt.Println("convert ", n.String())
		fmt.Println("params.AssetAmount ", params.AssetAmount)
		_, err := ctx.Client.AcceptOffer(params.OfferId, false, n)
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
		err := fmt.Errorf("AcceptOffer failed, failNum=%d,time=%v tx: %v", len(failedTx), Duration, failedTx)
		writeInfo(ctx.Index, fmt.Sprintf("%v", Duration), fmt.Sprintf(" %v", failedTx))
		return err
	}
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, Duration))
	return nil
}
func (c *AcceptOfferProcessor) End() {

}
