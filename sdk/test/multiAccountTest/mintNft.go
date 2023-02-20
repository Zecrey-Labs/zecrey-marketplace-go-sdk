package multiAccountTest

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
		t.RandomNftUrl = true
		t.RandomName = true
		t.RandomDescription = true
		t.Properties = "[]"
		t.Levels = "[]"
		t.Stats = "[]"
		t.Medias = []string{}
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

var MediaIndex int //from media dir

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
	bytes, err := ioutil.ReadFile(fmt.Sprintf("%s%s/key%d", NftTestTmp, Collection2Nft, ctx.Index))
	if err != nil {
		return fmt.Errorf("ignore it not have collection")
	}
	var _collectionInfo []collection
	err = json.Unmarshal(bytes, &_collectionInfo)
	if err != nil || len(_collectionInfo) == 0 {
		return fmt.Errorf("ignore it not have collection")
	}

	CollectionId := _collectionInfo[0].CollectionId
	MediaIndex++
	bytes, _ = ioutil.ReadFile(fmt.Sprintf("%smedias/key%d", NftTestTmp, MediaIndex))
	var media []string
	json.Unmarshal(bytes, &media)
	now := time.Now()
	var nftinfo []NftInfo
	for idx := 0; idx < c.Repeat; idx++ {
		if len(media) == 0 {
			return fmt.Errorf("index out of range [0] with length 0 key%d", ctx.Index)
		}
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
	Duration := time.Now().Sub(now)
	if len(nftinfo) > 0 {
		bytes, _ = json.Marshal(nftinfo)
		ioutil.WriteFile(fmt.Sprintf("%s%s/key%d", NftTestTmp, NftDir, ctx.Index), bytes, 0644)

	}

	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}

	if len(failedTx) > 0 {
		writeInfo(ctx.Index, fmt.Sprintf("%v", Duration), fmt.Sprintf(" %v", failedTx))
		return fmt.Errorf("mintnft failed,failNums:%d index=%d time=%v tx: %v", len(failedTx), ctx.Index, Duration, failedTx)
	}
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", ctx.Index, c.Repeat, Duration))

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
