package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func withdrawNftWrongBatch(index int) {
	for i := 0; i < index; i++ {
		for j := 0; j < i*10000; j++ {
			go withdrawNftAssetIdWrong(index)
			go withdrawNftL1AddressWrong(index)
		}
	}
}

func withdrawNftAssetIdWrong(index int) {
	result, err := client.WithdrawNft(rand.Int63n(10000000)+100000000, cfg.Tol1Address)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftAssetIdWrong", index, "WithdrawNft", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftL1AddressWrong", index, " json.Marshal", err.Error()))
		return
	}
}

func withdrawNftL1AddressWrong(index int) {
	result, err := client.WithdrawNft(cfg.WithdrawAssetId, cfg.BoundaryStr2)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftL1AddressWrong", index, "WithdrawNft", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftL1AddressWrong", index, "json.Marshal", err.Error()))
		return
	}
}
