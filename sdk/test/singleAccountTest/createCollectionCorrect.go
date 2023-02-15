package singleAccountTest

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"time"
)

func CreateCollectionCorrectBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go createCollectionCorrect(index)
		time.Sleep(time.Millisecond)
	}
}

func createCollectionCorrect(index int) {
	_, err := Client.CreateCollection(Cfg.ShortName, Cfg.CategoryId, Cfg.CreatorEarningRate,
		model.WithCollectionUrl(Cfg.CollectionUrl),
		model.WithExternalLink(Cfg.ExternalLink),
		model.WithTwitterLink(Cfg.TwitterLink),
		model.WithInstagramLink(Cfg.InstagramLink),
		model.WithTelegramLink(Cfg.TelegramLink),
		model.WithDiscordLink(Cfg.DiscordLink),
		model.WithLogoImage(Cfg.LogoImage),
		model.WithFeaturedImage(Cfg.FeaturedImage),
		model.WithBannerImage(Cfg.BannerImage),
		model.WithDescription(Cfg.Description))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "createCollectionCorrect", index, "CreateCollection", err.Error()))
		return
	} else {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,result=success", "createCollectionCorrect", index, "CreateCollection"))
	}

}
