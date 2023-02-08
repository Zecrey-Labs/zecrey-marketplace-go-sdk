package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
)

func acceptOfferWrongBatch(index int) {
	for j := 0; j < index*10000; j++ {
		go acceptOfferWrong(index)
	}
}

func acceptOfferWrong(index int) {
	result, err := client.AcceptOffer(rand.Int63n(10000)+1000000, false, big.NewInt(1230000))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeOfferCorrect", index, "MintNft", err.Error()))
		return
	}
	data, err := json.Marshal(result)
	fmt.Println("AcceptOffer:", string(data))
}
