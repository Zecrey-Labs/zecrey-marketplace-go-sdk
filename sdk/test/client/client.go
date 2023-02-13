package client

import (
	"encoding/hex"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/ethereum/go-ethereum/common"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	curve "github.com/zecrey-labs/zecrey-crypto/ecc/ztwistededwards/tebn254"
	"github.com/zecrey-labs/zecrey-crypto/util/ecdsaHelper"
	legendSdk "github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"

	"sync"
	"time"
)

type Ctx struct {
	Client      *sdk.Client
	L1Addr      common.Address
	AccountInfo *legendSdk.RespGetAccountInfoByPubKey
	Seed        string
	Index       int
}

func GetCtx(_client *sdk.Client, index int) *Ctx {
	privateKey, err := ethercrypto.LoadECDSA(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/test_account_in_dev_count_1000/%s", fmt.Sprintf("key%d", index)))
	privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
	l1Addr, err := ecdsaHelper.GenerateL1Address(privateKey)
	_, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
	legendClient := legendSdk.NewZecreyLegendSDK("https://dev-legend-app.zecrey.com")
	sk, err := curve.GenerateEddsaPrivateKey(seed)
	AccountInfo, err := legendClient.GetAccountInfoByPubKey(hex.EncodeToString(sk.PublicKey.Bytes()))
	if err != nil {
		fmt.Printf("NewClient failed:%v", err)
	}
	return &Ctx{_client, common.HexToAddress(l1Addr), AccountInfo, seed, index}
}

func StartTest(accountNum, testType int) {
	wg := sync.WaitGroup{}
	wg.Add(accountNum)
	now := time.Now()
	for index := 0; index < accountNum; index++ {

	}
	wg.Wait()
	fmt.Println(fmt.Sprintf("==== test over all time=%v ", time.Now().Sub(now)))
}
