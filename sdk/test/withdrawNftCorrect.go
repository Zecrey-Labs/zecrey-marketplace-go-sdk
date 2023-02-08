package main

import (
	"fmt"
	"time"
)

func withdrawNftCorrectOnce(index int) {
	if index == 1 {
		for j := 0; j < index*PerMinute; j++ {
			go withdrawNftCorrect(index)
			time.Sleep(time.Millisecond)
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
