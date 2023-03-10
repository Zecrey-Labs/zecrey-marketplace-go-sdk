package multiAccountTest

import (
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/ethereum/go-ethereum/common"
	ethercrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/zecrey-labs/zecrey-crypto/util/ecdsaHelper"
	legendSdk "github.com/zecrey-labs/zecrey-legend-go-sdk/sdk"
	"io/ioutil"
	"os"

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
	privateKey, err := ethercrypto.LoadECDSA(fmt.Sprintf("%s%s%s", NftTestTmp, KeyDir, fmt.Sprintf("key%d", index)))
	if err != nil {
		fmt.Println(fmt.Sprintf("ethercrypto.LoadECDSA failed:%v index:%d", err, index))
		return nil
	}
	privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
	l1Addr, err := ecdsaHelper.GenerateL1Address(privateKey)
	l2PublicKey, seed, err := sdk.GetSeedAndL2Pk(privateKeyString)
	if err != nil {
		fmt.Println(fmt.Sprintf("sdk.GetSeedAndL2Pk failed:%v index:%d", err, index))
		//panic(fmt.Sprintf("GetSeedAndL2Pk failed:%v", err))
		return nil
	}
	legendClient := legendSdk.NewZecreyLegendSDK(TestNetwork)
	AccountInfo, err := legendClient.GetAccountInfoByPubKey(l2PublicKey)
	//fmt.Println("seed[2:]:", seed[2:], "privateKeyString:", privateKeyString, "l2PublicKey:", l2PublicKey, "l1Addr", l1Addr, "name:", AccountInfo.AccountName, "Index", index)
	if err != nil {
		fmt.Println(fmt.Sprintf("legendClient.GetAccountInfoByPubKey failed:%v index:%d", err, index))
		//panic(fmt.Sprintf("GetSeedAndL2Pk failed:%v", err))
		return nil
	}
	fmt.Println(AccountInfo.AccountName, "Index", index, "AccountInfo", AccountInfo.AccountIndex)
	if err != nil {
		//panic(fmt.Sprintf("NewClient failed:%v", err))
		return nil
	}
	client, err := sdk.NewClientNoSuffix(AccountInfo.AccountName, seed)
	if err != nil {
		//panic(err)
		return nil
	}
	return &Ctx{privateKeyString, client, common.HexToAddress(l1Addr), AccountInfo, seed, index}
}
func GetCtxAmber(index int) *Ctx {
	privateKey, err := ethercrypto.LoadECDSA(fmt.Sprintf("%s%s%s", NftTestTmp, KeyDir, fmt.Sprintf("key%d", index)))
	if err != nil {
		fmt.Println(fmt.Sprintf("ethercrypto.LoadECDSA failed:%v index:%d", err, index))
		return nil
	}
	privateKeyString := hex.EncodeToString(ethercrypto.FromECDSA(privateKey))
	l1Addr, err := ecdsaHelper.GenerateL1Address(privateKey)
	l2PublicKey, seed, err := sdk.GetSeedAndL2Pk(privateKeyString)
	if err != nil {
		fmt.Println(fmt.Sprintf("sdk.GetSeedAndL2Pk failed:%v index:%d", err, index))
		//panic(fmt.Sprintf("GetSeedAndL2Pk failed:%v", err))
		return nil
	}
	legendClient := legendSdk.NewZecreyLegendSDK(TestNetwork)
	AccountInfo, err := legendClient.GetAccountInfoByPubKey(l2PublicKey)
	//fmt.Println("seed[2:]:", seed[2:], "privateKeyString:", privateKeyString, "l2PublicKey:", l2PublicKey, "l1Addr", l1Addr, "name:", AccountInfo.AccountName, "Index", index)
	if err != nil {
		fmt.Println(fmt.Sprintf("legendClient.GetAccountInfoByPubKey failed:%v index:%d", err, index))
		//panic(fmt.Sprintf("GetSeedAndL2Pk failed:%v", err))
		return nil
	}
	fmt.Println(AccountInfo.AccountName, "Index", index)
	if err != nil {
		//panic(fmt.Sprintf("NewClient failed:%v", err))
		return nil
	}
	client, err := sdk.NewClientNoSuffix(AccountInfo.AccountName, seed[2:])
	if err != nil {
		//panic(err)
		return nil
	}
	return &Ctx{privateKeyString, client, common.HexToAddress(l1Addr), AccountInfo, seed[2:], index}
}

var xlsFile *os.File

func StartTest(accountNum int, testType TxType) {
	MediaIndex = 0 //mediaIndex
	xlsFile1, _ := initCsv(testType)
	xlsFile = xlsFile1
	defer xlsFile.Close()
	wg := sync.WaitGroup{}
	wg.Add(accountNum)
	now := time.Now()
	Processor := GetProcessors().processorsMap[testType]
	count := 0
	for index := 0; index < accountNum; index++ {
		time.Sleep(5 * time.Millisecond)
		go func(_index int) {
			defer wg.Done()
			if ctx := GetCtx(_index); ctx != nil {
				if err := Processor.Process(ctx); err != nil {
					fmt.Println(err)
				} else {
					count++
				}
			}
		}(index)
	}

	wg.Wait()
	Processor.End()
	bytes, err := json.Marshal(accountNoMoney)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(fmt.Sprintf("/Users/user0/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/noMoney.json"), bytes, 0644)
	fmt.Println(fmt.Sprintf("==== test over all time=%v\n success=%d", time.Now().Sub(now), count))
}

func writeInfo(index int, Duration string, errStr string) {
	wStr := csv.NewWriter(xlsFile)

	s0 := []string{fmt.Sprintf("%d", index), Duration, errStr}
	err := wStr.Write(s0)
	if err != nil {
		fmt.Println(err)
	}
	wStr.Flush()
}

func initCsv(testType TxType) (*os.File, error) {
	strTime := time.Now().Format("20060102150405")
	nameList := map[TxType]string{}
	nameList[TxTypeCreateCollection] = fmt.Sprintf("CreateCollection_%s.csv", strTime)
	nameList[TxTypeMint] = fmt.Sprintf("MintNft_%s.csv", strTime)
	nameList[TxTypeTransfer] = fmt.Sprintf("TransferNft_%s.csv", strTime)
	nameList[TxTypeMatch] = fmt.Sprintf("MatchOffer_%s.csv", strTime)
	nameList[TxTypeCancelOffer] = fmt.Sprintf("CancelOffer_%s.csv", strTime)
	nameList[TxTypeWithdrawNft] = fmt.Sprintf("WithdrawNft_%s.csv", strTime)
	nameList[TxTypeListOffer] = fmt.Sprintf("ListOffer_%s.csv", strTime)
	filename := fmt.Sprintf("%s.csv", nameList[testType])
	xlsFile, fErr := os.OpenFile("./"+filename, os.O_RDWR|os.O_CREATE, 0766)
	if fErr != nil {
		fmt.Println("Export:created excel file failed ==", fErr)
		return nil, fErr
	}
	xlsFile.WriteString("\xEF\xBB\xBF")
	wStr := csv.NewWriter(xlsFile)
	wStr.Write([]string{"index", "Duration", "errStr"})
	wStr.Flush()
	return xlsFile, nil
}
