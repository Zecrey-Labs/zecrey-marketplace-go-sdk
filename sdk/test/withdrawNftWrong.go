package main

import (
	"fmt"
	"math/rand"
)

func withdrawNftWrongBatch(index int) {
	for j := 0; j < index*10000; j++ {
		go withdrawNftAssetIdWrong(index)
		go withdrawNftL1AddressWrong(index)
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
