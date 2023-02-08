package main

import (
	"fmt"
)

func transferNftCorrectOnce(index int) {
	if index == 1 {
		for j := 0; j < index*10000; j++ {
			go transferNftCorrect(index)
		}
	}

}

func transferNftCorrect(index int) {
	accountName, _, _ := client.GetMyInfo()
	resultSdk, err := getPreTransferNftTx(accountName, cfg.ToAccountName, fmt.Sprintf("%d", cfg.TransferAssetId))
	if err != nil {
		fmt.Println(fmt.Sprintf("success ! txType=%s,index=%d,func=%s,err=%s", "transferNftCorrect", index, "getPreTransferNftTx", err.Error()))
	}
	_, err = SignAndSendTransferNftTx(client.GetKeyManager(), fmt.Sprintf("%d", cfg.TransferAssetId), resultSdk.Transtion)
	if err != nil {
		fmt.Println(fmt.Sprintf("success ! txType=%s,index=%d,func=%s,err=%s", "transferNftCorrect", index, "MintNft", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("fail ! txType=%s,index=%d,func=%s,err=%s", "transferNftCorrect", index, "MintNft", err.Error()))
	}

}
