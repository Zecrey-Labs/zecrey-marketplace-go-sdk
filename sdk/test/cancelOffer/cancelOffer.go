package cancelOffer

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/util"
	"github.com/ethereum/go-ethereum/common"
	"io/ioutil"
	"path/filepath"
	"sync"
)

/*
一个账号去取消先前listOffer的
一个账号即可完成
*/

func InitCtx(_client *sdk.Client, _l1Addr common.Address) *ClientCtx {
	return &ClientCtx{_client, _l1Addr}
}

type ClientCtx struct {
	Client *sdk.Client
	L1Addr common.Address
}

func (c *ClientCtx) CancelOfferTest() error {
	data, err := ioutil.ReadFile(filepath.Join(".", util.DefaultDir, util.Offer2Cancel))
	if err != nil {
		panic(err)
	}
	var offerList []struct {
		offerId int64
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
			_, err := c.Client.CancelOffer(params.offerId)
			if err != nil {
				fmt.Println(fmt.Errorf("CancelOffer failed %s", err.Error()))
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
		return fmt.Errorf("cancelOffer failed, failedTx: %v", failedTx)
	}
	return nil
}
