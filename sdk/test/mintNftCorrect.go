package main

import (
	"encoding/json"
	"fmt"
)

func mintNftCorrectBatch(index int) {
	for j := 0; j < index*10000; j++ {
		go mintNftCorrect(index)
	}
}
func mintNftCorrect(index int) {
	Name := fmt.Sprintf("nftName:%s", cfg.AccountName)
	Description := fmt.Sprintf("nft Description %s", cfg.NftDescription)
	ret, err := client.MintNft(
		cfg.CollectionId,
		cfg.NftUrl, Name,
		Description, cfg.NftMedia,
		cfg.Properties, cfg.Levels, cfg.Stats)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftCorrect", index, "MintNft", err.Error()))
		return
	}
	_, err = json.Marshal(ret)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftCorrect", index, "MintNft.json.Marshal", err.Error()))
		return
	}
}
