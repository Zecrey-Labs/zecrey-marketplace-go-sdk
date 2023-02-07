package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

func makeOfferCorrectBatch(index int) {
	for i := 0; i < index; i++ {
		for j := 0; j < i*10000; j++ {
			go makeSellOfferCorrect(index)
			go makeBuyOfferCorrect(index)
		}
	}
}

func makeSellOfferCorrect(index int) {
	result, err := client.CreateSellOffer(cfg.SellAssetId, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferCorrect", index, "CreateSellOffer", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferCorrect", index, " json.Marshal", err.Error()))
		return
	}
}
func makeBuyOfferCorrect(index int) {
	result, err := client.CreateBuyOffer(cfg.SellAssetId, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "makeBuyOfferCorrect", index, "CreateBuyOffer", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "makeBuyOfferCorrect", index, " json.Marshal", err.Error()))
		return
	}
}
