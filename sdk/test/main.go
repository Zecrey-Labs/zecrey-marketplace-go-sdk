package test

import (
	"os"
	"testing"

	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

const (
	// TESTING_SEED is the seed used for testing
	ZECREY_TESTING_A_SEED_KEY_STRING = "ZECREY_TESTING_A_SEED_KEY_STRING"
	ZECREY_TESTING_A_ACCOUNT_NAME    = "ZECREY_TESTING_A_ACCOUNT_NAME"
	ZECREY_TESTING_B_SEED_KEY_STRING = "ZECREY_TESTING_B_SEED_KEY_STRING"
	ZECREY_TESTING_B_ACCOUNT_NAME    = "ZECREY_TESTING_B_ACCOUNT_NAME"

	nftMarketUrl       = "https://dev-legend-nft.zecrey.com"
	legendUrl          = "https://dev-legend-app.zecrey.com"
	hasuraUrl          = "https://legend-market-dev.hasura.app/v1/graphql"
	hasuraAdminKey     = "kqWAsFWVvn61mFuiuQ5yqJkWpu5VS1B5FGTdFzlVlQJ9fMTr9yNIjOnN3hERC9ex"
	hasuraTimeDeadline = 15 //15s
	chainRpcUrl        = "https://data-seed-prebsc-1-s1.binance.org:8545"
	DefaultGasLimit    = 5000000
	NameSuffix         = ".zec"
	MinGasFee          = 100000000000000 // 0.0001BNB

)

type testingClient struct {
	oAccountClient *sdk.Client
	nAccountClient *sdk.Client
}

func getTestingAccountClient(t *testing.T) testingClient {
	godotenv.Load()
	seed := os.Getenv(ZECREY_TESTING_A_SEED_KEY_STRING)
	name := os.Getenv(ZECREY_TESTING_A_ACCOUNT_NAME)
	oAccountClient, err := sdk.NewClient(name, seed)
	assert.Nil(t, err, "NewClient should not return an error, err: %v", err)
	seed = os.Getenv(ZECREY_TESTING_B_SEED_KEY_STRING)
	name = os.Getenv(ZECREY_TESTING_B_ACCOUNT_NAME)
	nAccountClient, err := sdk.NewClient(name, seed)
	assert.Nil(t, err, "NewClient should not return an error, err: %v", err)
	return testingClient{
		oAccountClient: oAccountClient,
		nAccountClient: nAccountClient,
	}
}
