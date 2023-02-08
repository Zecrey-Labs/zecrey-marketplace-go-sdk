package main

import (
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
)

func createCollectionCorrectBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go createCollectionCorrect(index)
	}
}

func createCollectionCorrect(index int) {
	_, err := client.CreateCollection(cfg.ShortName, cfg.CategoryId, cfg.CreatorEarningRate,
		model.WithCollectionUrl(cfg.CollectionUrl),
		model.WithExternalLink(cfg.ExternalLink),
		model.WithTwitterLink(cfg.TwitterLink),
		model.WithInstagramLink(cfg.InstagramLink),
		model.WithTelegramLink(cfg.TelegramLink),
		model.WithDiscordLink(cfg.DiscordLink),
		model.WithLogoImage(cfg.LogoImage),
		model.WithFeaturedImage(cfg.FeaturedImage),
		model.WithBannerImage(cfg.BannerImage),
		model.WithDescription(cfg.Description))
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "createCollectionCorrect", index, "CreateCollection", err.Error()))
		return
	} else {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,result=success", "createCollectionCorrect", index, "CreateCollection"))
	}

}
