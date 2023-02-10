package listOffer

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/util"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"math/rand"
	"path/filepath"
	"sync"
	"time"
)

/*

 */
var (
	log, _ = zap.NewDevelopment()
)

type ClientCtx struct {
	Client sdk.Client
	L1Addr common.Address
}
type RandomOption func(t *RandomOptionParam)

type RandomOptionParam struct {
	RandomAssetAmount bool

	AssetAmountDefault int64
}

func (c *ClientCtx) listOfferTest(ops ...RandomOption) error {
	option := RandomOptionParam{
		AssetAmountDefault: 1000000,
	}
	for _, op := range ops {
		op(&option)
	}
	data, err := ioutil.ReadFile(filepath.Join(".", util.DefaultDir, util.Nft2MakeSell))
	if err != nil {
		panic(err)
	}
	var nftList []struct {
		nftId int64
	}

	err = json.Unmarshal(data, &nftList)
	if err != nil {
		return err
	}

	repeat := len(nftList)
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	res := make([]struct {
		Success bool
		nftId   int64
		err     string
	}, repeat)

	//offer2Cancel := make([]struct {
	//	Success bool
	//	nftId   int64
	//	err     string
	//}, repeat)

	for i := 0; i < repeat; i++ {
		nftId := nftList[i].nftId
		assetAmount := c.randomTxParams(option)
		go func(idx int) {
			defer wg.Done()
			res[idx].nftId = nftId
			_, err := c.Client.CreateSellOffer(nftId, 0, big.NewInt(assetAmount))
			if err != nil {
				log.Error("CancelOffer failed", zap.Error(err))
				res[idx].Success = false
				res[idx].err = err.Error()
				return
			}
			res[idx].Success = true
		}(i)
	}

	wg.Wait()
	var failedTx []string

	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.err)
		}
	}

	if len(failedTx) > 0 {
		return fmt.Errorf("CreateSellOffer failed, failedTx: %v", failedTx)
	}
	return nil
}

func (c *ClientCtx) randomTxParams(option RandomOptionParam) int64 {
	rand.Seed(time.Now().UnixNano())
	if option.RandomAssetAmount {
		return rand.Int63n(1000000000)
	}
	return option.AssetAmountDefault
}
