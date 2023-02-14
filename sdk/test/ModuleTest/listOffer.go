package ModuleTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"time"
)

type ListOfferRandomOptionParam struct {
	RandomAssetAmount bool

	AssetAmountDefault int64
	Repeat             int
}

type ListOfferRandomOption func(t *ListOfferRandomOptionParam)

type ListOfferProcessor struct {
	Repeat        int
	AssetAmount   int64
	RandomOptions []ListOfferRandomOption
}
type OfferInfo struct {
	AccountKeyIndex int
	PrivateKey      string
	AssetAmount     string
	OfferId         int64
}

func NewlistOfferProcessor(RandomOptions ...ListOfferRandomOption) *ListOfferProcessor {
	r := &ListOfferProcessor{
		RandomOptions: RandomOptions,
	}
	option := ListOfferRandomOptionParam{}
	for _, op := range r.RandomOptions {
		op(&option)
	}
	r.randomTxParams(option)
	return r
}

func (t *ListOfferProcessor) Process(ctx *Ctx) error {
	option := ListOfferRandomOptionParam{
		AssetAmountDefault: 1000000,
	}
	for _, op := range t.RandomOptions {
		op(&option)
	}
	now := time.Now()
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

	var offer2cancel []OfferInfo
	for idx := 0; idx < t.Repeat; idx++ {
		nftId := nftinfo[idx].NftId
		t.randomTxParams(option)
		resp, err := ctx.Client.CreateSellOffer(nftId, 0, big.NewInt(t.AssetAmount))
		if err != nil {
			res[idx].Success = false
			res[idx].Err = err.Error()
			continue
		}
		res[idx].Success = true
		offer2cancel = append(offer2cancel, OfferInfo{ctx.Index, ctx.PrivateKey, resp.Offer.PaymentAssetAmount, resp.Offer.Id})
	}

	bytes, _ := json.Marshal(offer2cancel)
	ioutil.WriteFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s/key%d", OfferDir, ctx.Index), bytes, 0644)
	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}
	if len(failedTx) > 0 {
		err := fmt.Errorf("ListOffer failed,index=%d,time=%v,failNum=%d tx: %v \n", ctx.Index, time.Now().Sub(now), len(failedTx), failedTx)
		//fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v errs=%v", ctx.Index, t.Repeat, time.Now().Sub(now), err))
		return err
	}
	fmt.Println(fmt.Sprintf("ListOffer index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, time.Now().Sub(now)))
	return nil
}

func (t *ListOfferProcessor) randomTxParams(option ListOfferRandomOptionParam) *ListOfferProcessor {
	rand.Seed(time.Now().UnixNano())
	t.Repeat = option.Repeat
	if option.RandomAssetAmount {
		t.AssetAmount = rand.Int63n(1000000000)
	}
	return t
}
func (t *ListOfferProcessor) End() {

}
