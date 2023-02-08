package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

func makeOfferWrongBatch(index int) {
	for j := 0; j < index*10000; j++ {
		go makeSellOffeSellAssetIdWrong(index)
		go makeSellOfferAssetTypeWrong(index)
		go makeSellOfferAssetAmountWrong(index)
	}
}

func makeSellOffeSellAssetIdWrong(index int) {
	result, err := client.CreateSellOffer(rand.Int63()+1000000000000, 0, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOffeSellAssetIdWrong", index, "CreateSellOffer", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOffeSellAssetIdWrong", index, " json.Marshal", err.Error()))
		return
	}
}

func makeSellOfferAssetTypeWrong(index int) {
	result, err := client.CreateSellOffer(cfg.SellAssetId, rand.Int63n(10)+10, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetTypeWrong", index, "CreateSellOffer", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetTypeWrong", index, " json.Marshal", err.Error()))
		return
	}
}

func makeSellOfferAssetAmountWrong(index int) {
	result, err := client.CreateSellOffer(cfg.SellAssetId, 0, big.NewInt(-rand.Int63n(100000000000000)))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetAmountWrong", index, "CreateSellOffer", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetAmountWrong", index, " json.Marshal", err.Error()))
		return
	}
}
