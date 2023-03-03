package singleAccountTest

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func MakeOfferCorrectBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go makeSellOfferCorrect(index)
		time.Sleep(5 * time.Millisecond)
		//go makeBuyOfferCorrect(index)
		//time.Sleep(time.Millisecond)
	}
}

func makeSellOfferCorrect(index int) {
	_, err := Client.CreateSellOffer(Cfg.SellAssetId, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferCorrect", index, "CreateSellOffer", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s", "makeSellOfferCorrect", index, "CreateSellOffer"))
	}
}
func MakeBuyOfferCorrect(index int) {
	_, err := Client.CreateBuyOffer(Cfg.BuyAssetId, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "makeBuyOfferCorrect", index, "CreateBuyOffer", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s", "makeBuyOfferCorrect", index, "CreateBuyOffer"))
	}
}
