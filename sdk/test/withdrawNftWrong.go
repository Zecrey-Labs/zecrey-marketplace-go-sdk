package main

import (
	"fmt"
	"math/rand"
	"time"
)

func withdrawNftWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go withdrawNftAssetIdWrong(index)
		time.Sleep(time.Millisecond)
		go withdrawNftL1AddressWrong(index)
		time.Sleep(time.Millisecond)
	}
}

func withdrawNftAssetIdWrong(index int) {
	_, err := client.WithdrawNft(rand.Int63n(10000000)+100000000, cfg.Tol1Address)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftAssetIdWrong", index, "WithdrawNft", err.Error()))
	}
}

func withdrawNftL1AddressWrong(index int) {
	_, err := client.WithdrawNft(cfg.WithdrawAssetId, cfg.BoundaryStr2)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftL1AddressWrong", index, "WithdrawNft", err.Error()))
	}

}
