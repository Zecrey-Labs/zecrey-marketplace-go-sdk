package multiAccountTest

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
	UseForAccept      bool

	AssetAmountDefault int64
	Repeat             int
}

type ListOfferRandomOption func(t *ListOfferRandomOptionParam)

type ListOfferProcessor struct {
	UseForAccept  bool
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
		AssetAmountDefault: 100,
	}
	for _, op := range t.RandomOptions {
		op(&option)
	}

	var nftInfo []NftInfo
	data, err := ioutil.ReadFile(fmt.Sprintf("%s%s/key%d", NftTestTmp, NftDir, ctx.Index))
	if err != nil {
		return fmt.Errorf("ignore")
	}

	err = json.Unmarshal(data, &nftInfo)
	if err != nil {
		return err
	}
	res := make([]struct {
		Success bool
		Err     string
	}, t.Repeat)
	now := time.Now()
	var offer2cancel []OfferInfo
	for idx := 0; idx < t.Repeat; idx++ {
		nftId := nftInfo[idx].NftId
		t.randomTxParams(option)
		resp, err := ctx.Client.CreateSellOffer(nftId, 0, big.NewInt(0).SetInt64(t.AssetAmount))
		if err != nil {
			res[idx].Success = false
			res[idx].Err = err.Error()
			continue
		}
		res[idx].Success = true
		offer2cancel = append(offer2cancel, OfferInfo{ctx.Index, ctx.PrivateKey, fmt.Sprintf("%d", t.AssetAmount), resp.Offer.Id})
	}
	if len(offer2cancel) > 0 {
		bytes, _ := json.Marshal(offer2cancel)
		if t.UseForAccept {
			ioutil.WriteFile(fmt.Sprintf("%s%s/key%d", NftTestTmp, OfferDir, ctx.Index+2), bytes, 0644)
		} else {
			ioutil.WriteFile(fmt.Sprintf("%s%s/key%d", NftTestTmp, OfferDir, ctx.Index), bytes, 0644)
		}
	}
	Duration := time.Now().Sub(now)
	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}
	if len(failedTx) > 0 {
		err := fmt.Errorf("ListOffer failed,index=%d,time=%v,failNum=%d tx: %v \n", ctx.Index, Duration, len(failedTx), failedTx)
		writeInfo(ctx.Index, fmt.Sprintf("%v", Duration), fmt.Sprintf(" %v", failedTx))
		return err
	}
	fmt.Println(fmt.Sprintf("ListOffer index=%d,successNum=%d,time=%v", ctx.Index, t.Repeat, Duration))
	return nil
}

func (t *ListOfferProcessor) randomTxParams(option ListOfferRandomOptionParam) *ListOfferProcessor {
	rand.Seed(time.Now().UnixNano())
	t.Repeat = option.Repeat
	t.AssetAmount = option.AssetAmountDefault
	t.UseForAccept = option.UseForAccept
	if option.RandomAssetAmount {
		t.AssetAmount = rand.Int63n(100000000000)
	}
	return t
}
func (t *ListOfferProcessor) End() {

}
