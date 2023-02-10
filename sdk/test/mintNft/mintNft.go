package mintNft

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
	"math/rand"
	"sync"
	"time"
)

type RandomOptionParam struct {
	RandomCollectionId bool
	RandomNftUrl       bool
	RandomName         bool
	RandomDescription  bool
	RandomMedia        bool
	RandomProperties   bool
	RandomLevels       bool
	RandomStats        bool

	CollectionId int64
}

var (
	log, _ = zap.NewDevelopment()
)

type ClientCtx struct {
	Client sdk.Client
	L1Addr common.Address
}
type RandomOption func(t *RandomOptionParam)
type MitNftTxInfo struct {
	CollectionId int64
	NftUrl       string
	Name         string
	Description  string
	Media        string
	Properties   string
	Levels       string
	Stats        string
}

func (c *ClientCtx) MitNftTest(repeat int, ops ...RandomOption) error {
	option := RandomOptionParam{}
	for _, op := range ops {
		op(&option)
	}
	//pre get
	txInfo := MitNftTxInfo{}

	wg := sync.WaitGroup{}
	wg.Add(repeat)
	res := make([]struct {
		Success bool
		Err     string
	}, repeat)

	for i := 0; i < repeat; i++ {
		randTxInfo := c.randomTxInfo(txInfo, option)

		go func(idx int) {
			defer wg.Done()
			_, err := c.Client.MintNft(randTxInfo.CollectionId, randTxInfo.NftUrl, randTxInfo.Name, randTxInfo.Description, randTxInfo.Media,
				randTxInfo.Properties, randTxInfo.Levels, randTxInfo.Stats)
			if err != nil {
				log.Error("MintNft failed", zap.Error(err))
				res[idx].Success = false
				res[idx].Err = err.Error()
				return
			}
			res[idx].Success = true
		}(i)
	}

	wg.Wait()
	var failedTx []string

	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}

	if len(failedTx) > 0 {
		return fmt.Errorf("mintnft failed, tx: %v", failedTx)
	}
	return nil
}

func (c *ClientCtx) randomTxInfo(txInfo MitNftTxInfo, option RandomOptionParam) MitNftTxInfo {
	rand.Seed(time.Now().UnixNano())
	if option.RandomCollectionId {
	}
	if option.RandomNftUrl {

	}
	if option.RandomName {

	}
	if option.RandomDescription {

	}
	if option.RandomMedia {

	}
	if option.RandomProperties {

	}
	if option.RandomLevels {

	}
	if option.RandomStats {

	}
	return txInfo
}
