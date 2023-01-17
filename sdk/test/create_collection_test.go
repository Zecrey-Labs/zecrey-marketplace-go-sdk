package test

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk/model"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/stretchr/testify/assert"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"testing"
)

type marketCreateCollectionTxInfo struct {
	CollectionId       int64
	ShortName          string
	CategoryId         string
	CreatorEarningRate string
	ops                []model.CollectionOption
}

var createCollectionTestCase = []struct {
	txinfo   *marketCreateCollectionTxInfo
	expected bool
}{
	{
		txinfo: &marketCreateCollectionTxInfo{
			//AccountIndex:      0,
			//CollectionId:      0,
			//Name:              ";DROP TABLE account;",
			//Introduction:      ";DROP TABLE account;",
			//GasAccountIndex:   0,
			//GasFeeAssetId:     0,
			//GasFeeAssetAmount: big.NewInt(0),
			//ExpiredAt:         0,
			//Nonce:             0,
		},
		expected: false,
	},
	{
		txinfo: &marketCreateCollectionTxInfo{
			//AccountIndex:      0,
			//CollectionId:      0,
			//Name:              string([]byte{math.MaxUint8, math.MaxUint8}),
			//Introduction:      ";DROP TABLE account;",
			//GasAccountIndex:   0,
			//GasFeeAssetId:     0,
			//GasFeeAssetAmount: big.NewInt(0),
			//ExpiredAt:         0,
			//Nonce:             0,
		},
		expected: false,
	},
	{
		txinfo: &marketCreateCollectionTxInfo{
			//AccountIndex:      -1,
			//CollectionId:      -1,
			//Name:              string([]byte{math.MaxUint8, math.MaxUint8}),
			//Introduction:      ";DROP TABLE account;",
			//GasAccountIndex:   0,
			//GasFeeAssetId:     0,
			//GasFeeAssetAmount: big.NewInt(0).Mul(big.NewInt(math.MaxInt64), big.NewInt(1000000000000000000)),
			//ExpiredAt:         0,
			//Nonce:             -1,
		},
		expected: false,
	},
	{
		txinfo: &marketCreateCollectionTxInfo{
			//AccountIndex:      math.MaxInt64,
			//CollectionId:      -1,
			//Name:              string([]byte{math.MaxUint8, math.MaxUint8}),
			//Introduction:      ";DROP TABLE account;",
			//GasAccountIndex:   0,
			//GasFeeAssetId:     0,
			//GasFeeAssetAmount: big.NewInt(0).Mul(big.NewInt(math.MaxInt64), big.NewInt(1000000000000000000)),
			//ExpiredAt:         0,
			//Nonce:             -1,
		},
		expected: false,
	},
}

func TestCreateCollection(t *testing.T) {
	// create collection
	tc := getTestingAccountClient(t)
	oAccountClient := tc.oAccountClient
	accountName, _, _ := oAccountClient.GetMyInfo()
	//assert.Greater(t, nonce, int64(0), "nonce should be greater than 0")
	txInfoSdk, err := getPreCollectionTx(accountName)
	assert.Nil(t, err, "SignAndSendCreateCollectionTx should not return an error, err: %v", err)
	txInfo := &sdk.CreateCollectionTxInfo{}
	err = json.Unmarshal([]byte(txInfoSdk.Transtion), txInfo)
	if err != nil {
		t.Fatal(err)
	}
	//reset
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	t.Log(txInfo)
	for _, test := range createCollectionTestCase {
		_, err := SignAndSendCreateCollectionTx(oAccountClient.GetKeyManager(), txInfo, test.txinfo.ShortName, test.txinfo.CategoryId, test.txinfo.CreatorEarningRate, test.txinfo.ops...)
		if test.expected {
			assert.Nil(t, err, "SignAndSendCreateCollectionTx should not return an error, err: %v", err)
		} else {
			assert.NotNil(t, err, "SignAndSendCreateCollectionTx should return an error")
		}
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
