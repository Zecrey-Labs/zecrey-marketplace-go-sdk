package ModuleTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type CancelOfferRandomOptionParam struct {
	Repeat int
}

type CancelOfferRandomOption func(t *CancelOfferRandomOptionParam)

type CancelOfferProcessor struct {
	Repeat int
}

func NewCancelOfferProcessor(RandomOptions ...CancelOfferRandomOption) *CancelOfferProcessor {
	r := &CancelOfferProcessor{}
	op := CancelOfferRandomOptionParam{}
	for _, option := range RandomOptions {
		option(&op)
	}
	r.Repeat = op.Repeat
	return r
}

func (t *CancelOfferProcessor) Process(ctx *Ctx) error {
	var offer2cancel []OfferInfo
	data, err := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s/key%d", OfferDir, ctx.Index))
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &offer2cancel)
	if err != nil {
		return err
	}
	res := make([]struct {
		Success bool
		Err     string
	}, t.Repeat)
	now := time.Now()
	for idx := 0; idx < t.Repeat; idx++ {
		res[idx].Success = true
		params := offer2cancel[idx]
		_, err := ctx.Client.CancelOffer(params.OfferId)
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
		err := fmt.Errorf("CancelOffer failed,index=%d, failNum=%d tx: %v", ctx.Index, len(failedTx), failedTx)
		writeInfo(ctx.Index, fmt.Sprintf("%v", Duration), fmt.Sprintf(" %v", failedTx))
		return err
	}
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, Duration))
	return nil
}
func (c *CancelOfferProcessor) End() {

}
