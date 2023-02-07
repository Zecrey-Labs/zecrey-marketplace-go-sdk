package main

import (
	"fmt"
)

func transferNftCorrectBetch(index int) {
	for i := 0; i < index; i++ {
		for j := 0; j < i*10000; j++ {
			go transferNftCorrect(index)
		}
	}
}

func transferNftCorrect(index int) {
	accountName, _, _ := client.GetMyInfo()
	resultSdk, err := getPreTransferNftTx(accountName, cfg.ToAccountName, fmt.Sprintf("%d", cfg.TransferAssetId))
	_, err = SignAndSendTransferNftTx(client.GetKeyManager(), fmt.Sprintf("%d", cfg.TransferAssetId), resultSdk.Transtion)
	if err != nil {
		fmt.Println(fmt.Sprintf("success ! txType=%s,index=%d,func=%s,err=%s", "transferNftCorrect", index, "MintNft", err.Error()))
		return
	}

}
