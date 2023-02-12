package createCollection

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"math/rand"
	"time"
)

type RandomOptionParam struct {
	RandomShortName          bool
	RandomCategoryId         bool
	RandomCreatorEarningRate bool
	RandomOps                bool

	Ops                []model.CollectionOption
	CategoryId         int64
	CreatorEarningRate int64
}

/*
一个账号即可完成
*/
var (
	log, _ = zap.NewDevelopment()
)

type ClientCtx struct {
	Client *sdk.Client
	L1Addr common.Address
}
type RandomOption func(t *RandomOptionParam)
type TxInfo struct {
	ShortName          string
	CategoryId         string
	CreatorEarningRate string
	Ops                []model.CollectionOption
}

func InitCtx(_client *sdk.Client, _l1Addr common.Address) *ClientCtx {
	return &ClientCtx{_client, _l1Addr}
}
func (c *ClientCtx) CreateCollectionTest(repeat int, ops ...RandomOption) error {
	option := RandomOptionParam{}
	for _, op := range ops {
		op(&option)
	}
	//pre get
	txInfo := TxInfo{}

	res := make([]struct {
		Success bool
		Err     string
	}, repeat)

	for i := 0; i < repeat; i++ {
		randTxInfo := c.randomTxInfo(txInfo, option)

		go func(idx int) {
			_, err := c.Client.CreateCollection(randTxInfo.ShortName, randTxInfo.CategoryId, randTxInfo.CreatorEarningRate, randTxInfo.Ops...)
			if err != nil {
				log.Error("CreateCollection failed", zap.Error(err))
				res[idx].Success = false
				res[idx].Err = err.Error()
				return
			}
			res[idx].Success = true
		}(i)
	}
	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}

	if len(failedTx) > 0 {
		return fmt.Errorf("CreateCollection failed, tx: %v", failedTx)
	}
	return nil
}

func (c *ClientCtx) randomTxInfo(txInfo TxInfo, option RandomOptionParam) TxInfo {
	rand.Seed(time.Now().UnixNano())
	if option.RandomShortName {
		txInfo.ShortName = fmt.Sprintf("createCollectionTest%d", rand.Int())
	}
	if option.RandomCreatorEarningRate {

	}
	if option.RandomCategoryId {

	}
	if option.RandomOps {

	}
	return txInfo
}
