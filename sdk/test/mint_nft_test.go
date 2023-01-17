package test

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"testing"
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

var mintNftTestCase = []struct {
	txinfo   *MintNftTxInfo
	expected bool
}{
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: 1,
			//ToAccountIndex:      1,
			//ToAccountNameHash:   "0x0000000000000000000000000000000000000000000000000000000000000000",
			//NftIndex:            0,
			//NftContentHash:      "0x000",
			//NftCollectionId:     1,
			//CreatorTreasuryRate: 1,
			//GasAccountIndex:     1,
			//GasFeeAssetId:       0,
			//GasFeeAssetAmount:   big.NewInt(100),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               1,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: -1,
			//ToAccountIndex:      -1,
			//ToAccountNameHash:   "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
			//NftIndex:            0,
			//NftContentHash:      "0x000",
			//NftCollectionId:     math.MinInt64,
			//CreatorTreasuryRate: math.MinInt64,
			//GasAccountIndex:     1,
			//GasFeeAssetId:       0,
			//GasFeeAssetAmount:   big.NewInt(0).Add(big.NewInt(math.MaxInt64), big.NewInt(1)),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               1,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: math.MaxInt64,
			//ToAccountIndex:      math.MaxInt64,
			//ToAccountNameHash:   "0x@@@@@@@!*#@^%!*@@^%#*~@*&^!@%#(",
			//NftIndex:            -1,
			//NftContentHash:      "0x000",
			//NftCollectionId:     -1,
			//CreatorTreasuryRate: -1,
			//GasAccountIndex:     -1,
			//GasFeeAssetId:       -1,
			//GasFeeAssetAmount:   big.NewInt(-1),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               1,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: 100,
			//ToAccountIndex:      0,
			//ToAccountNameHash:   string([]byte{math.MaxUint8}),
			//NftIndex:            math.MinInt64,
			//NftContentHash:      "0x000",
			//NftCollectionId:     -1,
			//CreatorTreasuryRate: -1,
			//GasAccountIndex:     0,
			//GasFeeAssetId:       0,
			//GasFeeAssetAmount:   big.NewInt(100),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               1,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: -1,
			//ToAccountIndex:      1,
			//ToAccountNameHash:   "",
			//NftIndex:            0,
			//NftContentHash:      "0x000",
			//NftCollectionId:     1,
			//CreatorTreasuryRate: 1,
			//GasAccountIndex:     1,
			//GasFeeAssetId:       0,
			//GasFeeAssetAmount:   big.NewInt(100),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               -1,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: 1,
			//ToAccountIndex:      1,
			//ToAccountNameHash:   string([]byte{math.MaxUint8, math.MaxUint8, math.MaxUint8}),
			//NftIndex:            0,
			//NftContentHash:      "EOF",
			//NftCollectionId:     1,
			//CreatorTreasuryRate: 1,
			//GasAccountIndex:     1,
			//GasFeeAssetId:       0,
			//GasFeeAssetAmount:   big.NewInt(100),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               1,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: 1,
			//ToAccountIndex:      1,
			//ToAccountNameHash:   "qzzzz",
			//NftIndex:            0,
			//NftContentHash:      "0x000",
			//NftCollectionId:     1,
			//CreatorTreasuryRate: 1,
			//GasAccountIndex:     1,
			//GasFeeAssetId:       0,
			//GasFeeAssetAmount:   big.NewInt(100),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               1,
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			//CreatorAccountIndex: int64(math.NaN()),
			//ToAccountIndex:      -1,
			//ToAccountNameHash:   "<<<<<<<<<<<<<<<",
			//NftIndex:            1,
			//NftContentHash:      "0x000102",
			//NftCollectionId:     -1,
			//CreatorTreasuryRate: math.MaxInt64,
			//GasAccountIndex:     1,
			//GasFeeAssetId:       0,
			//GasFeeAssetAmount:   big.NewInt(100),
			//ExpiredAt:           time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:               1,
		},
		expected: false,
	},
}

func TestMintNft(t *testing.T) {
	tc := getTestingAccountClient(t)
	oAccountClient := tc.oAccountClient
	accountName, _, _ := oAccountClient.GetMyInfo()
	//assert.Nil(t, err, "GetNextNonce should not return an error, err: %v", err)
	//assert.Greater(t, nonce, int64(0), "nonce should be greater than 0")
	//assert.Nil(t, err, "GetGasFeeByAssetIdAndAccountIndex should not return an error, err: %v", err)
	//assert.Greater(t, gasFee, int64(0), "gasFee should be greater than 0")

	//assert.Nil(t, err, "SignAndSendCreateCollectionTx should not return an error, err: %v", err)
	//
	//gasFee, err = oAccountClient.GetGasFee(0, sdk.TxTypeMintNft)
	//assert.Nil(t, err, "GetGasFee failed")
	//nonce, err = oAccountClient.GetNextNonce(oAccountInfo.AccountIndex)
	//assert.Nil(t, err, "GetNextNonce failed")

	//assert.Nil(t, err, "SignAndSendMintNftTx failed")
	for _, test := range mintNftTestCase {
		resultSdk, err := getPreMintNftTx(accountName, test.txinfo.CollectionId, test.txinfo.Name, "test.txinfo.ContentHash")
		txInfo := &sdk.MintNftTxInfo{}
		err = json.Unmarshal([]byte(resultSdk.Transtion), txInfo)
		if err != nil {
			t.Fatal(err)
		}
		txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
		//txInfo.CreatorTreasuryRate = -1
		txInfo.CreatorTreasuryRate = 1000000000 //65535
		data, err := json.Marshal(txInfo)
		_, err = SignAndSendMintNftTx(test.txinfo.CollectionId, test.txinfo.NftUrl, test.txinfo.Name, test.txinfo.Description, test.txinfo.Media, test.txinfo.Properties, test.txinfo.Levels, test.txinfo.Stats, string(data))
		if test.expected {
			assert.Nil(t, err, "SignAndSendMintNftTx failed")
		} else {
			assert.NotNil(t, err, "SignAndSendMintNftTx failed")
		}
	}
}

func SignAndSendMintNftTx(CollectionId, NftUrl, Name, Description, Media, _PropertiesByte, _LevelsByte, _StatsByte, tx string) (*sdk.RespCreateAsset, error) {
	resp, err := http.PostForm(nftMarketUrl+"/api/v1/asset/createAsset",
		url.Values{
			"collection_id": {fmt.Sprintf("%s", CollectionId)},
			"nft_url":       {NftUrl},
			"name":          {Name},
			"description":   {Description},
			"media":         {Media},
			"properties":    {string(_PropertiesByte)},
			"levels":        {string(_LevelsByte)},
			"stats":         {string(_StatsByte)},
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
