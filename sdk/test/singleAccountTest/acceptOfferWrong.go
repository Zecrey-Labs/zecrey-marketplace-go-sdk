package singleAccountTest

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func AcceptOfferWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go acceptOfferWrong(index)
		time.Sleep(time.Millisecond)
	}
}

func acceptOfferWrong(index int) {
	_, err := Client.AcceptOffer(rand.Int63n(10000)+1000000, false, big.NewInt(1230000))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "makeOfferCorrect", index, "MintNft", err.Error()))
		return
	}
}