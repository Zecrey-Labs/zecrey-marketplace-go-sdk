package ModuleTest

import (
	"encoding/hex"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/ethereum/go-ethereum/common"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/zecrey-labs/zecrey-crypto/util/ecdsaHelper"
	legendSdk "github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"

	"sync"
	"time"
)

type Ctx struct {
	PrivateKey  string
	Client      *sdk.Client
	L1Addr      common.Address
	AccountInfo *legendSdk.RespGetAccountInfoByPubKey
	Seed        string
	Index       int
}

func GetCtx(index int) *Ctx {
	privateKey, err := ethercrypto.LoadECDSA(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/localPrivateKeys/%s", fmt.Sprintf("key%d", index)))
	privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
	l1Addr, err := ecdsaHelper.GenerateL1Address(privateKey)
	l2PublicKey, seed, _ := sdk.GetSeedAndL2Pk(privateKeyString)
	legendClient := legendSdk.NewZecreyLegendSDK("https://test-legend-app.zecrey.com")
	AccountInfo, err := legendClient.GetAccountInfoByPubKey(l2PublicKey)
	if err != nil {
		panic(fmt.Sprintf("NewClient failed:%v", err))
	}
	client, err := sdk.NewClientNoSuffix(AccountInfo.AccountName, seed)
	if err != nil {
		panic(err)
	}
	return &Ctx{privateKeyString, client, common.HexToAddress(l1Addr), AccountInfo, seed, index}
}

func StartTest(accountNum int, testType TxType) {
	wg := sync.WaitGroup{}
	wg.Add(accountNum)
	now := time.Now()
	Processor := GetProcessors().processorsMap[testType]
	var errs []error
	for index := 0; index < accountNum; index++ {
		ctx := GetCtx(index)
		go func() {
			defer wg.Done()
			if err := Processor.Process(ctx); err != nil {
				errs = append(errs, err)
			}
		}()
	}

	wg.Wait()
	Processor.End()
	fmt.Println(fmt.Sprintf("==== test over all time=%v\nerrs=%v", time.Now().Sub(now), errs))
}
