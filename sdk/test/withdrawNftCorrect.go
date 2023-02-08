package main

import (
	"fmt"
)

func withdrawNftCorrectOnce(index int) {
	if index == 1 {
		for j := 0; j < index*10000; j++ {
			go withdrawNftCorrect(index)
		}
	}
}

func withdrawNftCorrect(index int) {
	_, err := client.WithdrawNft(cfg.WithdrawAssetId, cfg.Tol1Address)
	if err != nil {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,err=%s", "withdrawNftAssetIdWrong", index, "WithdrawNft", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "withdrawNftAssetIdWrong", index, "WithdrawNft", err.Error()))
	}
}
