package singleAccountTest

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func MakeOfferWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go makeSellOfferSellAssetIdWrong(index)
		time.Sleep(time.Millisecond)
		go makeSellOfferAssetTypeWrong(index)
		time.Sleep(time.Millisecond)
		go makeSellOfferAssetAmountWrong(index)
		time.Sleep(time.Millisecond)
	}
}

func makeSellOfferSellAssetIdWrong(index int) {
	result, err := Client.CreateSellOffer(rand.Int63()+1000000000000, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferSellAssetIdWrong", index, "CreateSellOffer", err.Error()))
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferSellAssetIdWrong", index, " json.Marshal", err.Error()))
	}
}

// can success
func makeSellOfferAssetTypeWrong(index int) {
	_, err := Client.CreateSellOffer(Cfg.SellAssetId, rand.Int63n(10)+10, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetTypeWrong", index, "CreateSellOffer", err.Error()))
	}
}

// can success
func makeSellOfferAssetAmountWrong(index int) {
	_, err := Client.CreateSellOffer(Cfg.SellAssetId, 0, big.NewInt(rand.Int63n(100000000000000)))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetAmountWrong", index, "CreateSellOffer", err.Error()))
	}

}