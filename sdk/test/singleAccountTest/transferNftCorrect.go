package singleAccountTest

import (
	"fmt"
	"time"
)

func TransferNftCorrectOnce(index int) {
	if index == 1 {
		for j := 0; j < index*PerMinute; j++ {
			go transferNftCorrect(index)
			time.Sleep(time.Millisecond)
		}
	}
}

func transferNftCorrect(index int) {
	accountName, _, _ := Client.GetMyInfo()
	resultSdk, err := getPreTransferNftTx(accountName, Cfg.ToAccountName, fmt.Sprintf("%d", Cfg.TransferAssetId))
	if err != nil {
		fmt.Println(fmt.Sprintf("success ! txType=%s,index=%d,func=%s,err=%s", "transferNftCorrect", index, "getPreTransferNftTx", err.Error()))
		return
	}
	_, err = SignAndSendTransferNftTx(Client.GetKeyManager(), fmt.Sprintf("%d", Cfg.TransferAssetId), resultSdk.Transtion)
	if err != nil {
		fmt.Println(fmt.Sprintf("success ! txType=%s,index=%d,func=%s,err=%s", "transferNftCorrect", index, "MintNft", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("fail ! txType=%s,index=%d,func=%s,err=%s", "transferNftCorrect", index, "MintNft", err.Error()))
	}

}
