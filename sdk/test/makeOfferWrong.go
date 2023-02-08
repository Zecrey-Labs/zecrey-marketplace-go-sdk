package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

func makeOfferWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go makeSellOfferSellAssetIdWrong(index)
		go makeSellOfferAssetTypeWrong(index)
		go makeSellOfferAssetAmountWrong(index)
	}
}

func makeSellOfferSellAssetIdWrong(index int) {
	result, err := client.CreateSellOffer(rand.Int63()+1000000000000, 0, big.NewInt(rand.Int63()))
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
	_, err := client.CreateSellOffer(cfg.SellAssetId, rand.Int63n(10)+10, big.NewInt(rand.Int63()))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetTypeWrong", index, "CreateSellOffer", err.Error()))
	}
}

// can success
func makeSellOfferAssetAmountWrong(index int) {
	_, err := client.CreateSellOffer(cfg.SellAssetId, 0, big.NewInt(rand.Int63n(100000000000000)))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeSellOfferAssetAmountWrong", index, "CreateSellOffer", err.Error()))
	}

}
