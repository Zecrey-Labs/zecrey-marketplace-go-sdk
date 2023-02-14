package mintNft

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/ethereum/go-ethereum/common"
	"math/rand"
	"time"
)

/*
 一个账号mint 1000个
*/
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
	Properties   string
	Levels       string
	Stats        string
	Medias       []string
}

var ()

type ClientCtx struct {
	Client *sdk.Client
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

func InitCtx(_client *sdk.Client, _l1Addr common.Address) *ClientCtx {
	return &ClientCtx{_client, _l1Addr}
}

func (c *ClientCtx) MitNftTest(repeat, Index int, ops ...RandomOption) error {
	option := RandomOptionParam{}
	for _, op := range ops {
		op(&option)
	}
	//pre get
	txInfo := MitNftTxInfo{}
	res := make([]struct {
		Success bool
		Err     string
	}, repeat)

	for idx := 0; idx < repeat; idx++ {
		randTxInfo := c.randomTxInfo(txInfo, option)
		resp, err := c.Client.MintNft(randTxInfo.CollectionId, randTxInfo.NftUrl, randTxInfo.Name, randTxInfo.Description, option.Medias[idx],
			randTxInfo.Properties, randTxInfo.Levels, randTxInfo.Stats)
		if err != nil {
			res[idx].Success = false
			res[idx].Err = err.Error()
			continue
		}
		res[idx].Success = true
		fmt.Println(fmt.Sprintf("Index=%d,nftId=%d", Index, resp.Asset.Id))
	}

	var failedTx []string

	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}

	if len(failedTx) > 0 {
		return fmt.Errorf("mintnft failed,failNums:%d tx: %v", len(failedTx), failedTx)
	}
	return nil
}

func (c *ClientCtx) randomTxInfo(txInfo MitNftTxInfo, option RandomOptionParam) MitNftTxInfo {
	txInfo.CollectionId = option.CollectionId
	txInfo.Properties = option.Properties
	txInfo.Levels = option.Levels
	txInfo.Stats = option.Stats

	rand.Seed(time.Now().UnixNano())
	if option.RandomCollectionId {

	}
	if option.RandomNftUrl {
		txInfo.NftUrl = fmt.Sprintf("mintNftUrlTest%d", rand.Int())
	}
	if option.RandomName {
		txInfo.Name = fmt.Sprintf("mintNftTest%d", rand.Int())
	}
	if option.RandomDescription {
		txInfo.Description = fmt.Sprintf("mintNftDescriptionTest%d", rand.Int())
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
