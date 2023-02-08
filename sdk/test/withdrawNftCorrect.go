package main

import (
	"encoding/json"
	"fmt"
)

func withdrawNftCorrectBatch(index int) {
	for j := 0; j < index*10000; j++ {
		go withdrawNftCorrect(index)
	}
}

func withdrawNftCorrect(index int) {
	result, err := client.WithdrawNft(cfg.WithdrawAssetId, cfg.Tol1Address)
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "withdrawNftAssetIdWrong", index, "WithdrawNft", err.Error()))
		return
	}
	_, err = json.Marshal(result)
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "withdrawNftL1AddressWrong", index, " json.Marshal", err.Error()))
		return
	}
}
