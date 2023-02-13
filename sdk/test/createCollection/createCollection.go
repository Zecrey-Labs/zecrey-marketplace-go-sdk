package createCollection

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/ethereum/go-ethereum/common"
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
func (c *ClientCtx) CreateCollectionTest(repeat, index int, ops ...RandomOption) error {
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

	for idx := 0; idx < repeat; idx++ {
		randTxInfo := c.randomTxInfo(txInfo, option)
		_, err := c.Client.CreateCollection(randTxInfo.ShortName, randTxInfo.CategoryId, randTxInfo.CreatorEarningRate, randTxInfo.Ops...)
		if err != nil {
			res[idx].Success = false
			res[idx].Err = err.Error()
			continue
		}
		res[idx].Success = true
	}
	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}

	if len(failedTx) > 0 {
		return fmt.Errorf("CreateCollection failed,index=%d failNum=%d tx: %v", index, len(failedTx), failedTx)
	}
	return nil
}

func (c *ClientCtx) randomTxInfo(txInfo TxInfo, option RandomOptionParam) TxInfo {
	rand.Seed(time.Now().UnixNano())
	txInfo.Ops = option.Ops
	txInfo.CategoryId = fmt.Sprintf("%d", option.CategoryId)
	txInfo.CreatorEarningRate = fmt.Sprintf("%d", option.CreatorEarningRate)
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
