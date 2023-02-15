package singleAccountTest

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type MintNftTxInfo struct {
	CollectionId string
	NftUrl       string
	Name         string
	Description  string
	Media        string
	Properties   string
	Levels       string
	Stats        string
}

var mintNftTestCase = []string{
	"CollectionInfo", "Media", "NftUrl", "Description", "CreatorTreasuryRate", "Stats", "Levels", "Properties",
}

func MintNftCorrectWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go mintNftWrong(index)
		time.Sleep(time.Millisecond)
	}
}
func mintNftWrong(index int) {
	txDefaultinfo := &MintNftTxInfo{
		CollectionId: fmt.Sprintf("%d", Cfg.CollectionId),
		NftUrl:       Cfg.NftUrl,
		Name:         Cfg.NftName,
		Description:  Cfg.Description,
		Media:        Cfg.NftMediaWrong,
		Properties:   Cfg.Properties,
		Levels:       Cfg.Levels,
		Stats:        Cfg.Stats,
	}
	accountName, _, _ := Client.GetMyInfo()
	for _, testCase := range mintNftTestCase {
		txinfo := *txDefaultinfo
		txinfo.Name = fmt.Sprintf("%s %d ", txinfo.Name, rand.Int())
		resultSdk, err := getPreMintNftTx(accountName, txinfo.CollectionId, txinfo.Name, "txinfo.ContentHash")
		if err != nil {
			fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftCorrect", index, "getPreMintNftTx", err.Error()))
			return
		}
		txInfo := &sdk.MintNftTxInfo{}
		err = json.Unmarshal([]byte(resultSdk.Transtion), txInfo)
		if err != nil {
			fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftCorrect", index, "MintNft.json.Marshal", err.Error()))
			return
		}
		txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)

		_PropertiesByte, err := json.Marshal(txinfo.Properties)
		_LevelsByte, err := json.Marshal(txinfo.Levels)
		_StatsByte, err := json.Marshal(txinfo.Stats)
		switch testCase {
		case "CollectionInfo":
			r := rand.Intn(1000000) + 1000000000
			txinfo.CollectionId = fmt.Sprintf("%d", r)
		case "Media":
			txinfo.Media = Cfg.BoundaryStr2
		case "NftUrl":
			txinfo.NftUrl = Cfg.BoundaryStr2
		case "Description":
			r := rand.Intn(10)
			if r < 5 {
				txinfo.Description = Cfg.BoundaryStr2
			} else {
				txinfo.Description = Cfg.BoundaryStr3
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
				assetProperty.Name = Cfg.BoundaryStr2
			}
			_Properties := []sdk.Propertie{assetProperty}
			_PropertiesByte, err = json.Marshal(_Properties)
			Properties := string(_PropertiesByte)
			if r >= 5 {
				Properties = Cfg.BoundaryStr3
			}
			txinfo.Properties = Properties
		case "Levels":
			assetLevel := sdk.Level{
				Name:     fmt.Sprintf("assetLevel%d", rand.Intn(1000000)),
				Value:    int64(rand.Intn(1000000)),
				MaxValue: int64(rand.Intn(1000000)),
			}
			r := rand.Intn(10)
			if r < 5 {
				assetLevel.Name = Cfg.BoundaryStr2
			}
			_Levels := []sdk.Level{assetLevel}
			_LevelsByte, err = json.Marshal(_Levels)

			Levels := string(_LevelsByte)
			if r >= 5 {
				Levels = Cfg.BoundaryStr3
			}
			txinfo.Levels = Levels
		case "Stats":
			assetStats := sdk.Stat{
				Name:     fmt.Sprintf("assetStats%d", rand.Intn(1000000)),
				Value:    int64(rand.Intn(1000000)),
				MaxValue: int64(rand.Intn(1000000)),
			}
			r := rand.Intn(10)
			if r < 5 {
				assetStats.Name = Cfg.BoundaryStr2
			}
			_Stats := []sdk.Stat{assetStats}
			_StatsByte, err = json.Marshal(_Stats)

			Stats := string(_StatsByte)
			if r >= 5 {
				Stats = Cfg.BoundaryStr3
			}
			txinfo.Stats = Stats
		case "CreatorTreasuryRate":
			r := rand.Intn(100000000) + 100000
			txInfo.CreatorTreasuryRate = int64(r) //65535
		}

		_, err = SignAndSendMintNftTx(txinfo.CollectionId, txinfo.NftUrl, txinfo.Name, txinfo.Description, txinfo.Media, txinfo.Properties, txinfo.Levels, txinfo.Stats, txInfo)
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "mintNftWrong", index, "MintNft.SignAndSendMintNftTx", err.Error()))
	}
}

func SignAndSendMintNftTx(CollectionId, NftUrl, Name, Description, Media, _PropertiesByte, _LevelsByte, _StatsByte string, txInfo *sdk.MintNftTxInfo) (*sdk.RespCreateAsset, error) {
	tx, err := constructMintNftTx(Client.GetKeyManager(), txInfo)

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
		return nil, fmt.Errorf(string(body))
	}

	result := &sdk.RespCreateAsset{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func getPreMintNftTx(accountName, collectionId, name, contentHash string) (*sdk.RespetSdktxInfo, error) {
	respSdkTx, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkMintNftTxInfo?treasury_rate=20&account_name=%s&collection_id=%s&name=%s&content_hash=%s", accountName, collectionId, name, contentHash))
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

func constructMintNftTx(key sdk.KeyManager, tx *sdk.MintNftTxInfo) (string, error) {
	convertedTx := convertMintNftTxInfo(tx)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeMintNftMsgHash(convertedTx, hFunc)
	if err != nil {
		return "", err
	}
	hFunc.Reset()
	signature, err := key.Sign(msgHash, hFunc)
	if err != nil {
		return "", err
	}
	convertedTx.Sig = signature
	txInfoBytes, err := json.Marshal(convertedTx)
	if err != nil {
		return "", err
	}
	return string(txInfoBytes), nil
}
func convertMintNftTxInfo(tx *sdk.MintNftTxInfo) *legendTxTypes.MintNftTxInfo {
	return &legendTxTypes.MintNftTxInfo{
		CreatorAccountIndex: tx.CreatorAccountIndex,
		ToAccountIndex:      tx.ToAccountIndex,
		ToAccountNameHash:   tx.ToAccountNameHash,
		NftIndex:            tx.NftIndex,
		NftContentHash:      tx.NftContentHash,
		NftCollectionId:     tx.NftCollectionId,
		CreatorTreasuryRate: tx.CreatorTreasuryRate,
		GasAccountIndex:     tx.GasAccountIndex,
		GasFeeAssetId:       tx.GasFeeAssetId,
		GasFeeAssetAmount:   tx.GasFeeAssetAmount,
		ExpiredAt:           tx.ExpiredAt,
		Nonce:               tx.Nonce,
		Sig:                 tx.Sig,
	}
}
