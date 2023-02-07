package main

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
)

type MintNftTxInfo struct {
	WrongCase    string
	CollectionId string
	NftUrl       string
	Name         string
	Description  string
	Media        string
	Properties   string
	Levels       string
	Stats        string
}

var mintNftTestCase = []struct {
	txinfo   *MintNftTxInfo
	expected bool
}{
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "CollectionId",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "Media",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "NftUrl",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "Description",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "Properties",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "Levels",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "Stats",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			WrongCase:    "CreatorTreasuryRate",
			CollectionId: fmt.Sprintf("%d", cfg.CollectionId),
			NftUrl:       cfg.NftUrl,
			Name:         cfg.NftName,
			Description:  cfg.Description,
			Media:        cfg.NftMedia,
			Properties:   cfg.Properties,
			Levels:       cfg.Levels,
			Stats:        cfg.Stats,
		},
		expected: false,
	},
}

func mintNftCorrectWrongBatch(index int) {
	for i := 0; i < index; i++ {
		for j := 0; j < i*10000; j++ {
			go mintNftWrong(index)
		}
	}
}
func mintNftWrong(index int) {
	accountName, _, _ := client.GetMyInfo()
	for _, test := range mintNftTestCase {
		resultSdk, err := getPreMintNftTx(accountName, test.txinfo.CollectionId, test.txinfo.Name, "test.txinfo.ContentHash")
		txInfo := &sdk.MintNftTxInfo{}
		err = json.Unmarshal([]byte(resultSdk.Transtion), txInfo)
		if err != nil {
			fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftCorrect", index, "MintNft.json.Marshal", err.Error()))
			return
		}
		txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)

		txDataInfo, err := json.Marshal(txInfo)
		_PropertiesByte, err := json.Marshal(test.txinfo.Properties)
		_LevelsByte, err := json.Marshal(test.txinfo.Levels)
		_StatsByte, err := json.Marshal(test.txinfo.Stats)
		switch test.txinfo.WrongCase {
		case "CollectionId":
			r := rand.Intn(1000000) + 1000000000
			test.txinfo.CollectionId = fmt.Sprintf("%d", r)
		case "Media":
			test.txinfo.Media = cfg.BoundaryStr2
		case "NftUrl":
			test.txinfo.NftUrl = cfg.BoundaryStr2
		case "Description":
			r := rand.Intn(10)
			if r < 5 {
				test.txinfo.Description = cfg.BoundaryStr2
			} else {
				test.txinfo.Description = cfg.BoundaryStr3
			}
		case "Properties":
			key := fmt.Sprintf("color——%d", rand.Intn(1000000))
			value := "red"
			assetProperty := sdk.Propertie{
				Name:  key,
				Value: value,
			}
			r := rand.Intn(10)
			if r < 5 {
				assetProperty.Name = cfg.BoundaryStr2
			}
			_Properties := []sdk.Propertie{assetProperty}
			_PropertiesByte, err = json.Marshal(_Properties)
			Properties := string(_PropertiesByte)
			if r >= 5 {
				Properties = cfg.BoundaryStr3
			}
			test.txinfo.Properties = Properties
		case "Levels":
			assetLevel := sdk.Level{
				Name:     fmt.Sprintf("assetLevel%d", rand.Intn(1000000)),
				Value:    int64(rand.Intn(1000000)),
				MaxValue: int64(rand.Intn(1000000)),
			}
			r := rand.Intn(10)
			if r < 5 {
				assetLevel.Name = cfg.BoundaryStr2
			}
			_Levels := []sdk.Level{assetLevel}
			_LevelsByte, err = json.Marshal(_Levels)

			Levels := string(_LevelsByte)
			if r >= 5 {
				Levels = cfg.BoundaryStr3
			}
			test.txinfo.Levels = Levels
		case "Stats":
			assetStats := sdk.Stat{
				Name:     fmt.Sprintf("assetStats%d", rand.Intn(1000000)),
				Value:    int64(rand.Intn(1000000)),
				MaxValue: int64(rand.Intn(1000000)),
			}
			r := rand.Intn(10)
			if r < 5 {
				assetStats.Name = cfg.BoundaryStr2
			}
			_Stats := []sdk.Stat{assetStats}
			_StatsByte, err = json.Marshal(_Stats)

			Stats := string(_LevelsByte)
			if r >= 5 {
				Stats = cfg.BoundaryStr3
			}
			test.txinfo.Stats = Stats
		case "CreatorTreasuryRate":
			r := rand.Intn(100000000) + 100000
			txInfo.CreatorTreasuryRate = int64(r) //65535
		}

		_, err = SignAndSendMintNftTx(test.txinfo.CollectionId, test.txinfo.NftUrl, test.txinfo.Name, test.txinfo.Description, test.txinfo.Media, string(_PropertiesByte), string(_LevelsByte), string(_StatsByte), string(txDataInfo))
		if test.expected {
			fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftWrong", index, "MintNft.SignAndSendMintNftTx", err.Error()))
			return
		} else {
			fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftWrong", index, "MintNft.SignAndSendMintNftTx", err.Error()))
			return
		}
	}
}

// todo 签名
func SignAndSendMintNftTx(CollectionId, NftUrl, Name, Description, Media, _PropertiesByte, _LevelsByte, _StatsByte, tx string) (*sdk.RespCreateAsset, error) {
	resp, err := http.PostForm(nftMarketUrl+"/api/v1/asset/createAsset",
		url.Values{
			"collection_id": {fmt.Sprintf("%s", CollectionId)},
			"nft_url":       {NftUrl},
			"name":          {Name},
			"description":   {Description},
			"media":         {Media},
			"properties":    {_PropertiesByte},
			"levels":        {_LevelsByte},
			"stats":         {_StatsByte},
			"transaction":   {tx},
		},
	)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		//if assert.Contains(t, string(body), fmt.Errorf("CreatorTreasuryRate should  not be less than 0").Error()) {
		//	fmt.Println("pass")
		//	return
		//}
		//if assert.Contains(t, string(body), fmt.Errorf("CreatorTreasuryRate should not be larger than 65535").Error()) {
		//	fmt.Println("pass")
		//	return
		//}
		//t.Fatal(fmt.Errorf(string(body)))
		return nil, err
	}

	result := &sdk.RespCreateAsset{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func getPreMintNftTx(accountName, collectionId, name, contentHash string) (*sdk.RespetSdktxInfo, error) {
	respSdkTx, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkMintNftTxInfo?treasury_rate=20&account_name=%s&collection_id=%d&name=%s&content_hash=%s", accountName, collectionId, name, contentHash))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(respSdkTx.Body)
	if err != nil {
		return nil, err
	}
	if respSdkTx.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	resultSdk := &sdk.RespetSdktxInfo{}
	if err := json.Unmarshal(body, &resultSdk); err != nil {
		return nil, err
	}
	return resultSdk, err
}
