package main

import (
	"fmt"
	"math/rand"
)

//nft media just once chance
func mintNftCorrectOnce(index int) {
	if index == 1 {
		for j := 0; j < index*PerMinute;; j++ {
			go mintNftCorrect(index)
		}
	}
}
func mintNftCorrect(index int) {
	Name := fmt.Sprintf("nftName%s%d", cfg.NftName, rand.Int())
	Description := fmt.Sprintf("nft Description%s%d", cfg.NftDescription, rand.Int())
	_, err := client.MintNft(
		cfg.CollectionId,
		cfg.NftUrl, Name,
		Description, cfg.NftMedia,
		cfg.Properties, cfg.Levels, cfg.Stats)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftCorrect", index, "MintNft", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,result=success", "mintNftCorrect", index, "MintNft"))
	}

}
