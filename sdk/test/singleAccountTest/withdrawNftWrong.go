package singleAccountTest

import (
	"fmt"
	"math/rand"
	"time"
)

func WithdrawNftWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go withdrawNftAssetIdWrong(index)
		time.Sleep(time.Millisecond)
		go withdrawNftL1AddressWrong(index)
		time.Sleep(time.Millisecond)
	}
}

func withdrawNftAssetIdWrong(index int) {
	_, err := Client.WithdrawNft(rand.Int63n(10000000)+100000000, Cfg.Tol1Address)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftAssetIdWrong", index, "WithdrawNft", err.Error()))
	}
}

func withdrawNftL1AddressWrong(index int) {
	_, err := Client.WithdrawNft(Cfg.WithdrawAssetId, Cfg.BoundaryStr2)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftL1AddressWrong", index, "WithdrawNft", err.Error()))
	}

}
