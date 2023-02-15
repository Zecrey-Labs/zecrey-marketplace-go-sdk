package singleAccountTest

import (
	"flag"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
)

var ConfigFile = flag.String("f",
	"/Users/user0/work/zecrey-marketplace-go-sdk/sdk/test/config.yaml", "the config file")

var Cfg Config
var Client *sdk.Client
