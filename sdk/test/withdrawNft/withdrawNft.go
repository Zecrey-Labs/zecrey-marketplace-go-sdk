package withdrawNft

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/util"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
	"sync"
)

/*
一个账号即可搞定
*/
var (
	log, _ = zap.NewDevelopment()
)

type ClientCtx struct {
	Client *sdk.Client
	L1Addr common.Address
}
type RandomOption func(t *RandomOptionParam)

type RandomOptionParam struct {
	ToAccountAddress string
}

func (c *ClientCtx) withdrawNftTest(ops ...RandomOption) error {
	option := RandomOptionParam{
		ToAccountAddress: "0x0......",
	}
	for _, op := range ops {
		op(&option)
	}
	data, err := ioutil.ReadFile(filepath.Join(".", util.DefaultDir, util.Nft2Transfer))
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
		go func(idx int) {
			defer wg.Done()
			res[idx].nftId = nftId
			_, err := c.Client.WithdrawNft(nftId, option.ToAccountAddress)
			if err != nil {
				log.Error("WithdrawNft failed", zap.Error(err))
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
		return fmt.Errorf("WithdrawNft failed, failedTx: %v", failedTx)
	}
	return nil
}
