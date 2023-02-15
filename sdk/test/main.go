package main

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/multiAccountTest"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/singleAccountTest"
	"github.com/zeromicro/go-zero/core/conf"
	"time"
)

func testAll() {
	conf.MustLoad(*singleAccountTest.ConfigFile, &singleAccountTest.Cfg)
	_client, err := sdk.NewClient(singleAccountTest.Cfg.AccountName, singleAccountTest.Cfg.Seed)
	singleAccountTest.Client = _client
	if err != nil {
		panic(err)
	}

	for i := 1; i < 30; i++ {
		singleAccountTest.CreateCollectionCorrectBatch(i)
		singleAccountTest.CreateCollectionWrongBatch(i)
		singleAccountTest.MintNftCorrectOnce(i)
		singleAccountTest.MintNftCorrectWrongBatch(i)
		singleAccountTest.MakeOfferCorrectBatch(i)
		singleAccountTest.MakeOfferWrongBatch(i)
		singleAccountTest.TransferNftCorrectOnce(i)
		singleAccountTest.TransferNftWrongBatch(i)
		singleAccountTest.WithdrawNftCorrectOnce(i)
		singleAccountTest.WithdrawNftWrongBatch(i)
		singleAccountTest.AcceptOfferWrongBatch(i)
		time.Sleep(60 * time.Second)
	}

	time.Sleep(10 * time.Minute)
	fmt.Println("==== test over !!!")
}

func main() {
	multiAccountTest.StartTest(1000, multiAccountTest.TxTypeMatch)
}
