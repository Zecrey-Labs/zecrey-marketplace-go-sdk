package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func makeOfferCorrectBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go makeSellOfferCorrect(index)
		time.Sleep(time.Millisecond)
		go makeBuyOfferCorrect(index)
		time.Sleep(time.Millisecond)
	}
}

func makeSellOfferCorrect(index int) {
	_, err := client.CreateSellOffer(cfg.SellAssetId, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferCorrect", index, "CreateSellOffer", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s", "makeSellOfferCorrect", index, "CreateSellOffer"))
	}
}
func makeBuyOfferCorrect(index int) {
	_, err := client.CreateBuyOffer(cfg.BuyAssetId, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "makeBuyOfferCorrect", index, "CreateBuyOffer", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s", "makeBuyOfferCorrect", index, "CreateBuyOffer"))
	}
}
