package transferNft

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
一个账号即可搞定
*/
func InitCtx(_client *sdk.Client, _l1Addr common.Address) *ClientCtx {
	return &ClientCtx{_client, _l1Addr}
}

type ClientCtx struct {
	Client *sdk.Client
	L1Addr common.Address
}
type RandomOption func(t *RandomOptionParam)

type RandomOptionParam struct {
	ToAccountName string
}

func (c *ClientCtx) TransferNftTest(ops ...RandomOption) error {
	option := RandomOptionParam{
		ToAccountName: "sher",
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
			_, err := c.Client.TransferNft(nftId, option.ToAccountName)
			if err != nil {
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
		return fmt.Errorf("TransferNft failed, failedTx: %v", failedTx)
	}
	return nil
}