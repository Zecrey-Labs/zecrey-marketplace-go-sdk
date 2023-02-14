package ModuleTest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

type NftRandomOptionParam struct {
	RandomCollectionId bool
	RandomNftUrl       bool
	RandomName         bool
	RandomDescription  bool
	RandomMedia        bool
	RandomProperties   bool
	RandomLevels       bool
	RandomStats        bool

	Properties string
	Levels     string
	Stats      string
	Medias     []string
	Repeat     int
}
type NftInfo struct {
	AccountKeyIndex int
	PrivateKey      string
	NftId           int64
}
type NftRandomOption func(t *NftRandomOptionParam)
type MintNftProcessor struct {
	Repeat           int
	RandomOptions    []NftRandomOption
	RandomOptionNext NftRandomOptionParam

	NftUrl      string
	Name        string
	Description string
	Media       string
	Properties  string
	Levels      string
	Stats       string
}

func GetDefaultNftOption() NftRandomOption {
	r := func(t *NftRandomOptionParam) {
		var medias []string
		bytes, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s", Media2Nft))
		json.Unmarshal(bytes, &medias)
		t.RandomNftUrl = true
		t.RandomName = true
		t.RandomDescription = true
		t.Properties = "[]"
		t.Levels = "[]"
		t.Stats = "[]"
		t.Medias = medias
	}
	return r
}

func NewMintNftProcessor(RandomOptions ...NftRandomOption) *MintNftProcessor {
	RandomOptions = append(RandomOptions, GetDefaultNftOption())
	r := &MintNftProcessor{
		RandomOptions: RandomOptions,
	}
	option := NftRandomOptionParam{}
	for _, op := range r.RandomOptions {
		op(&option)
	}
	r.RandomOptionNext = option
	r.randomTxInfo(option)
	return r
}

//Process  to get medias from TestUploadMediaRepeat
func (c *MintNftProcessor) Process(ctx *Ctx) error {
	option := NftRandomOptionParam{}
	for _, op := range c.RandomOptions {
		op(&option)
	}
	//pre get
	res := make([]struct {
		Success bool
		Err     string
	}, c.Repeat)
	bytes, _ := ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/collection2Nft/key%d", ctx.Index))
	var _collectionInfo []collection
	json.Unmarshal(bytes, &_collectionInfo)

	CollectionId := _collectionInfo[0].CollectionId

	bytes, _ = ioutil.ReadFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/medias/key%d", ctx.Index))
	var media []string
	json.Unmarshal(bytes, &media)
	now := time.Now()
	var nftinfo []NftInfo
	for idx := 0; idx < c.Repeat; idx++ {
		c.randomTxInfo(option)
		resp, err := ctx.Client.MintNft(CollectionId, c.NftUrl, c.Name, c.Description, media[idx],
			c.Properties, c.Levels, c.Stats)
		if err != nil {
			res[idx].Success = false
			res[idx].Err = err.Error()
			continue
		}
		res[idx].Success = true
		nftinfo = append(nftinfo, NftInfo{ctx.Index, ctx.PrivateKey, resp.Asset.Id})

		//fmt.Println(fmt.Sprintf("Index=%d,nftId=%d", ctx.Index, resp.Asset.Id))
	}
	bytes, _ = json.Marshal(nftinfo)
	ioutil.WriteFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s/key%d", NftDir, ctx.Index), bytes, 0644)
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, c.Repeat, time.Now().Sub(now)))

	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}

	if len(failedTx) > 0 {
		return fmt.Errorf("mintnft failed,failNums:%d index=%d time=%v tx: %v", len(failedTx), ctx.Index, time.Now().Sub(now), failedTx)
	}
	return nil
}

func (c *MintNftProcessor) randomTxInfo(option NftRandomOptionParam) *MintNftProcessor {
	c.Properties = option.Properties
	c.Levels = option.Levels
	c.Stats = option.Stats
	c.Repeat = option.Repeat
	rand.Seed(time.Now().UnixNano())
	if option.RandomCollectionId {

	}
	if option.RandomNftUrl {
		c.NftUrl = fmt.Sprintf("mintNftUrlTest%d", rand.Int())
	}
	if option.RandomName {
		c.Name = fmt.Sprintf("mintNftTest%d", rand.Int())
	}
	if option.RandomDescription {
		c.Description = fmt.Sprintf("mintNftDescriptionTest%d", rand.Int())
	}
	if option.RandomMedia {
	}
	if option.RandomProperties {

	}
	if option.RandomLevels {

	}
	if option.RandomStats {

	}

	return c
}
func (c *MintNftProcessor) End() {

}