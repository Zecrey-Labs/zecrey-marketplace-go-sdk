package main

import (
	"flag"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/zeromicro/go-zero/core/conf"
	"time"
)

var configFile = flag.String("f",
	"./config.yaml", "the config file")

var cfg Config
var client *sdk.Client

func main() {
	conf.MustLoad(*configFile, &cfg)
	_client, err := sdk.NewClient(cfg.AccountName, cfg.Seed)
	client = _client
	if err != nil {
		panic(err)
	}

	for i := 1; i < 30; i++ {
		go createCollectionCorrectBatch(i)
		go createCollectionWrongBatch(i)
		go mintNftCorrectOnce(i)
		go mintNftCorrectWrongBatch(i)
		go makeOfferCorrectBatch(i)
		go makeOfferWrongBatch(i)
		go transferNftCorrectOnce(i)
		go transferNftWrongBatch(i)
		go withdrawNftCorrectOnce(i)
		go withdrawNftWrongBatch(i)
		go acceptOfferWrongBatch(i)
		time.Sleep(60 * time.Second)
	}
	time.Sleep(10 * time.Minute)
	panic("==== test over !!!")
}
