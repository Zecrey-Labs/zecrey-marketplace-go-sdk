package mintNft

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/test/util"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

type RandomOptionParam struct {
	RandomFromAccountIndex bool

	//AmountPerTx *big.Int
}

var (
	log, _    = zap.NewDevelopment()
	txTimeout = 60 * time.Second
)

type ClientCtx struct {
	Client *sdk.Client
	L1Addr common.Address
}
type RandomOption func(t *RandomOptionParam)
type UpdateCollectionTx struct {
	Id   string
	Name string
	ops  []model.CollectionOption
}

func (c *ClientCtx) MitNftTest(repeat int, ops ...RandomOption) error {
	option := RandomOptionParam{}
	for _, op := range ops {
		op(&option)
	}

	data, err := ioutil.ReadFile(filepath.Join(".", util.DefaultDir, util.Nft2Transfer))
	if err != nil {
		panic(err)
	}
	var collectionList []struct {
		collectionId int64
	}

	err = json.Unmarshal(data, &collectionList)
	if err != nil {
		return err
	}
	txInfo := UpdateCollectionTx{}

	wg := sync.WaitGroup{}
	wg.Add(repeat)
	res := make([]struct {
		Success bool
		Err     string
	}, repeat)

	for i := 0; i < repeat; i++ {
		go func(idx int) {
			defer wg.Done()
			randTxInfo := c.randomTxInfo(txInfo, option)
			_, err := c.Client.UpdateCollection(randTxInfo.Id, randTxInfo.Name, randTxInfo.ops...)
			if err != nil {
				log.Error("UpdateCollection failed", zap.Error(err))
				res[idx].Success = false
				res[idx].Err = err.Error()
				return
			}
			res[idx].Success = true
		}(i)
	}

	wg.Wait()
	failedTx := []string{}

	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}

	if len(failedTx) > 0 {
		return fmt.Errorf("UpdateCollection failed, tx hash: %v", failedTx)
	}
	return nil
}

func (c *ClientCtx) randomTxInfo(txInfo UpdateCollectionTx, option RandomOptionParam) UpdateCollectionTx {

	return txInfo
}
