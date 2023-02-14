package ModuleTest

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"io/ioutil"
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
	Repeat             int
}

type RandomOption func(t *RandomOptionParam)

func GetDefaultOption() RandomOption {
	CollectionUrl := "https://res.cloudinary.com/zecrey/image/upload/collection/ahykviwc0suhoyzusb5q.jpg"
	ExternalLink := "https://weibo.com/alice"
	TwitterLink := "https://twitter.com/alice"
	InstagramLink := "https://www.instagram.com/alice/"
	TelegramLink := "https://tgstat.com/channel/@alice"
	DiscordLink := "https://discord.com/api/v10/applications/<aliceid>/commands"
	LogoImage := "collection/j9w9z4dmcd2beufvkxkp"
	FeaturedImage := "collection/j9w9z4dmcd2beufvkxkp"
	BannerImage := "collection/j9w9z4dmcd2beufvkxkp"
	Description := "Description information"
	r := func(t *RandomOptionParam) {
		t.CategoryId = 1
		t.CreatorEarningRate = 200
		t.RandomShortName = true
		t.Ops = []model.CollectionOption{model.WithCollectionUrl(CollectionUrl),
			model.WithExternalLink(ExternalLink),
			model.WithTwitterLink(TwitterLink),
			model.WithInstagramLink(InstagramLink),
			model.WithTelegramLink(TelegramLink),
			model.WithDiscordLink(DiscordLink),
			model.WithLogoImage(LogoImage),
			model.WithFeaturedImage(FeaturedImage),
			model.WithBannerImage(BannerImage),
			model.WithDescription(Description)}
	}
	return r

}

type CreateCProcessor struct {
	Repeat int

	ShortName          string
	CategoryId         string
	CreatorEarningRate string
	Ops                []model.CollectionOption

	RandomOptions    []RandomOption
	RandomOptionNext RandomOptionParam
}

func NewCreateCollectionProcessor(RandomOptions ...RandomOption) *CreateCProcessor {
	RandomOptions = append(RandomOptions, GetDefaultOption())
	r := &CreateCProcessor{
		RandomOptions: RandomOptions,
	}
	option := RandomOptionParam{}
	for _, op := range r.RandomOptions {
		op(&option)
	}
	r.RandomOptionNext = option
	return r.randomTxInfo(option)
}

func (t *CreateCProcessor) Process(ctx *Ctx) error {
	index := ctx.Index
	nftClient := ctx.Client
	now := time.Now()

	res := make([]struct {
		Success bool
		Err     string
	}, t.Repeat)
	var _collectionInfo []collection

	for idx := 0; idx < t.Repeat; idx++ {
		t.randomTxInfo(t.RandomOptionNext)
		resp, err := nftClient.CreateCollection(t.ShortName, t.CategoryId, t.CreatorEarningRate, t.Ops...)
		if err != nil {
			res[idx].Success = false
			res[idx].Err = err.Error()
			continue
		}
		res[idx].Success = true
		_collectionInfo = append(_collectionInfo, collection{index, ctx.PrivateKey, resp.Collection.Id})
	}
	bytes, err := json.Marshal(_collectionInfo)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(fmt.Sprintf("/Users/zhangwei/work/zecrey-marketplace-go-sdk/sdk/test/.nftTestTmp/%s/key%d", Collection2Nft, index), bytes, 0644)
	fmt.Println(fmt.Sprintf("index=%d,successNum=%d,time=%v", index, t.Repeat, time.Now().Sub(now)))

	var failedTx []string
	for _, r := range res {
		if !r.Success {
			failedTx = append(failedTx, r.Err)
		}
	}
	if len(failedTx) > 0 {
		err := fmt.Errorf("CreateCollection failed,index=%d  failNum=%d   time=%v tx: %v", index, len(failedTx), time.Now().Sub(now), failedTx)
		return err
	}
	return nil
}

func (t *CreateCProcessor) randomTxInfo(option RandomOptionParam) *CreateCProcessor {
	rand.Seed(time.Now().UnixNano())
	t.Ops = option.Ops
	t.Repeat = option.Repeat
	t.CategoryId = fmt.Sprintf("%d", option.CategoryId)
	t.CreatorEarningRate = fmt.Sprintf("%d", option.CreatorEarningRate)
	if option.RandomShortName {
		t.ShortName = fmt.Sprintf("createCollectionTest%d", rand.Int())
	}
	if option.RandomCreatorEarningRate {

	}
	if option.RandomCategoryId {

	}
	if option.RandomOps {

	}
	return t
}

type collection struct {
	AccountKeyIndex int
	PrivateKey      string
	CollectionId    int64
}

var _collectionInfo []collection

func (t *CreateCProcessor) set(accountKeyIndex int, privateKey string, collectionId int64) {

}
func (t *CreateCProcessor) End() {
}
