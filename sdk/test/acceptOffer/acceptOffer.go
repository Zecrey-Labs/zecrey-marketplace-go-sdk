package acceptOffer

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/util"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"sync"
)

/*
先用一个账号sellOffer很多nft
再用other账号进行buyOffer
一共俩账号即可
*/

var (
	log, _ = zap.NewDevelopment()
)

type ClientCtx struct {
	Client sdk.Client
	L1Addr common.Address
}

func (c *ClientCtx) acceptOfferTest() error {
	data, err := ioutil.ReadFile(filepath.Join(".", util.DefaultDir, util.Offer2Accept))
	if err != nil {
		panic(err)
	}
	var offerList []struct {
		offerId     int64
		isSell      bool
		assetAmount *big.Int
	}

	err = json.Unmarshal(data, &offerList)
	if err != nil {
		return err
	}

	repeat := len(offerList)
	wg := sync.WaitGroup{}
	wg.Add(repeat)
	res := make([]struct {
		Success bool
		offerId int64
		err     string
	}, repeat)

	for i := 0; i < repeat; i++ {
		params := offerList[i]
		go func(idx int) {
			defer wg.Done()
			res[idx].offerId = params.offerId
			_, err := c.Client.AcceptOffer(params.offerId, params.offerId == 1, params.assetAmount)
			if err != nil {
				log.Error("AcceptOffer failed", zap.Error(err))
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
		return fmt.Errorf("AcceptOffer failed, failedTx: %v", failedTx)
	}
	return nil
}
