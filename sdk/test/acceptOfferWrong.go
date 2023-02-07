package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

func makeOfferWrongBetch(index int) {
	for i := 0; i < index; i++ {
		for j := 0; j < i*10000; j++ {
			go makeOfferWrong(index)
		}
	}
}
func makeOfferWrong(index int) {
	result, err := client.AcceptOffer(rand.Int63n(10000)+1000000, false, big.NewInt(1230000))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeOfferCorrect", index, "MintNft", err.Error()))
		return
	}
	data, err := json.Marshal(result)
	fmt.Println("AcceptOffer:", string(data))
}
