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

	for i := 0; i < 30; i++ {
		go createCollectionCorrectBatch(i)
		go createCollectionWrongBetch(i)
		time.Sleep(60 * time.Second)
	}
	select {}
}
