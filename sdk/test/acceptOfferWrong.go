package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func acceptOfferWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go acceptOfferWrong(index)
	}
}

func acceptOfferWrong(index int) {
	_, err := client.AcceptOffer(rand.Int63n(10000)+1000000, false, big.NewInt(1230000))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeOfferCorrect", index, "MintNft", err.Error()))
		return
	}
}
