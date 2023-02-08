package main

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type marketCreateCollectionTxInfo struct {
	ShortName          string
	CategoryId         string
	CreatorEarningRate string
	ops                []model.CollectionOption
}

var createCollectionTestCase = []string{
	"ShortName", "Link", "CreatorEarningRate", "CategoryId",
}

func createCollectionWrongBatch(index int) {
	for j := 0; j < index*PerMinute; j++ {
		go createCollectionWrong(index)
		time.Sleep(time.Millisecond)
	}
}

func createCollectionWrong(index int) {
	accountName, _, _ := client.GetMyInfo()
	txInfoSdk, err := getPreCollectionTx(accountName)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "createCollectionWrong", index, "createCollectionWrong", err.Error()))
		return
	}
	txInfo := &sdk.CreateCollectionTxInfo{}
	err = json.Unmarshal([]byte(txInfoSdk.Transtion), txInfo)
	if err != nil {
		fmt.Println(fmt.Sprintf("fail! txType=%s,index=%d,func=%s,err=%s", "createCollectionWrong", index, "createCollectionWrong.json.Unmarshal", err.Error()))
		return
	}
	//reset
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	txCaseDefaultInfo := &marketCreateCollectionTxInfo{
		ShortName:          cfg.ShortName,
		CategoryId:         cfg.CategoryId,
		CreatorEarningRate: cfg.CreatorEarningRate,
		ops: []model.CollectionOption{
			model.WithCollectionUrl(cfg.CollectionUrl),
			model.WithExternalLink(cfg.ExternalLink),
			model.WithTwitterLink(cfg.TwitterLink),
			model.WithInstagramLink(cfg.InstagramLink),
			model.WithTelegramLink(cfg.TelegramLink),
			model.WithDiscordLink(cfg.DiscordLink),
			model.WithLogoImage(cfg.LogoImage),
			model.WithFeaturedImage(cfg.FeaturedImage),
			model.WithBannerImage(cfg.BannerImage),
			model.WithDescription(cfg.Description)},
	}
	for _, testCase := range createCollectionTestCase {
		txCaseInfo := *txCaseDefaultInfo
		txCaseInfo.ShortName = fmt.Sprintf("txCaseInfo.ShortName#%d", rand.Int())
		switch testCase {
		case "ShortName":
			txCaseInfo.ShortName = cfg.BoundaryStr
		case "Link":
			l := len(txCaseInfo.ops)
			r := rand.Intn(l)
			for _index, do := range []model.CollectionOption{
				model.WithCollectionUrl(cfg.BoundaryStr2),
				model.WithExternalLink(cfg.BoundaryStr2),
				model.WithTwitterLink(cfg.BoundaryStr2),
				model.WithInstagramLink(cfg.BoundaryStr2),
				model.WithTelegramLink(cfg.BoundaryStr3),
				model.WithDiscordLink(cfg.BoundaryStr3),
				model.WithLogoImage(cfg.BoundaryStr3),
				model.WithFeaturedImage(cfg.BoundaryStr3),
				model.WithBannerImage(cfg.BoundaryStr3),
				model.WithDescription(cfg.BoundaryStr3),
			} {
				if _index == r {
					txCaseInfo.ops = append(txCaseInfo.ops, do)
					break
				}
			}
		case "CreatorEarningRate":
			r := rand.Intn(100000000) + 100000
			txCaseInfo.CreatorEarningRate = fmt.Sprintf("%d", r)
		case "CategoryId":
			r := rand.Intn(10000) + 10
			txCaseInfo.CategoryId = fmt.Sprintf("%d", r)
		}

		_, err := SignAndSendCreateCollectionTx(client.GetKeyManager(), txInfo, txCaseInfo.ShortName, txCaseInfo.CategoryId, txCaseInfo.CreatorEarningRate, txCaseInfo.ops...)

		fmt.Println(fmt.Sprintf("Hope to fail ,txType=%s,testType=%s,index=%d,func=%s,err=%s", "CreateCollection", "Wrong", index, "SignAndSendCreateCollectionTx2", err.Error()))

	}
}
func getPreCollectionTx(accountName string) (*sdk.RespetSdktxInfo, error) {
	respSdkTx, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCreateCollectionTxInfo?account_name=%s", accountName))
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
		return nil, fmt.Errorf(string(body))
	}
	return resultSdk, nil
}

func SignAndSendCreateCollectionTx(keyManager sdk.KeyManager, tx *sdk.CreateCollectionTxInfo, ShortName string, CategoryId string, CreatorEarningRate string, ops ...model.CollectionOption) (*sdk.RespCreateCollection, error) {
	cp := model.GetDefaultCollection()
	for _, do := range ops {
		do.F(cp)
	}
	txStr, err := sdkCreateCollectionTxInfo(keyManager, tx, cp.Description, ShortName)
	if err != nil {
		return nil, err
	}
	resp, err := http.PostForm(nftMarketUrl+"/api/v1/collection/createCollection",
		url.Values{
			"short_name":           {ShortName},
			"category_id":          {CategoryId},
			"collection_url":       {cp.CollectionUrl},
			"external_link":        {cp.ExternalLink},
			"twitter_link":         {cp.TwitterLink},
			"instagram_link":       {cp.TelegramLink},
			"discord_link":         {cp.InstagramLink},
			"telegram_link":        {cp.DiscordLink},
			"logo_image":           {cp.LogoImage},
			"featured_image":       {cp.FeaturedImage},
			"banner_image":         {cp.BannerImage},
			"creator_earning_rate": {CreatorEarningRate},
			"payment_asset_ids":    {cp.PaymentAssetIds},
			"transaction":          {txStr}})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}
	result := &sdk.RespCreateCollection{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func sdkCreateCollectionTxInfo(key sdk.KeyManager, tx *sdk.CreateCollectionTxInfo, Description, ShortName string) (string, error) {
	convertedTx := convertCreateCollectionTxInfo(tx)
	convertedTx.Name = ShortName
	convertedTx.Introduction = Description
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeCreateCollectionMsgHash(convertedTx, hFunc)
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
	return string(txInfoBytes), nil
}
func convertCreateCollectionTxInfo(tx *sdk.CreateCollectionTxInfo) *legendTxTypes.CreateCollectionTxInfo {
	return &legendTxTypes.CreateCollectionTxInfo{
		AccountIndex:      tx.AccountIndex,
		CollectionId:      tx.CollectionId,
		Name:              tx.Name,
		Introduction:      tx.Introduction,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}
