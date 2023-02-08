package main

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
)

func createCollectionCorrectBatch(index int) {
	for j := 0; j < index*10000; j++ {
		go createCollectionCorrect(index)
	}
}

func createCollectionCorrect(index int) {
	ret, err := client.CreateCollection(cfg.ShortName, cfg.CategoryId, cfg.CreatorEarningRate,
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
		fmt.Println(fmt.Sprintf("fail! txType=%s,testType=%s,index=%d,func=%s,err=%s", "createCollectionCorrect", index, "CreateCollection", err.Error()))
		return
	}
	data, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,testType=%s,index=%d,func=%s,err=%s", "createCollectionCorrect", index, "CreateCollection.json.Marshal", err.Error()))
		return
	}
	fmt.Println(fmt.Sprintf("success! txType=%s,testType=%s,index=%d,func=%s,result=%s", "createCollectionCorrect", index, "CreateCollection.json.Marshal", string(data)))
}
