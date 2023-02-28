package singleAccountTest

import (
	"fmt"
	"math/rand"
	"time"
)

//MintNftCorrectOnce nft media just once chance
func MintNftCorrectOnce(index int) {
	if index == 1 {
		for j := 0; j < index*PerMinute; j++ {
			go mintNftCorrect(index)
			time.Sleep(time.Millisecond)
		}
	}
}
func mintNftCorrect(index int) {
	Name := fmt.Sprintf("nftName%s%d", Cfg.NftName, rand.Int())
	Description := fmt.Sprintf("nft Description%s%d", Cfg.NftDescription, rand.Int())
	_, err := Client.MintNft(
		Cfg.CollectionId,
		"amber1.zec",
		Cfg.NftUrl, Name,
		Description, Cfg.NftMedia,
		Cfg.Properties, Cfg.Levels, Cfg.Stats)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftCorrect", index, "MintNft", err.Error()))
	} else {
		fmt.Println(fmt.Sprintf("success! txType=%s,index=%d,func=%s,result=success", "mintNftCorrect", index, "MintNft"))
	}

}
