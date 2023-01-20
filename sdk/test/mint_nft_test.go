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
	Properties   []sdk.Propertie
	Levels       []sdk.Level
	Stats        []sdk.Stat
}

var mintNftTestCase = []struct {
	txinfo   *MintNftTxInfo
	expected bool
}{
	{
		txinfo: &MintNftTxInfo{
			CollectionId: "123456789",
			NftUrl:       "-",
			Name:         fmt.Sprintf("nftName2:%s", "accountName"),
			Description:  fmt.Sprintf("%s `s nft", "accountName"),
			Media:        "collection/dz5hwqaszpwtflg0sfz4",
			Properties:   []sdk.Propertie{},
			Levels:       []sdk.Level{},
			Stats:        []sdk.Stat{},
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			CollectionId: "1",
			NftUrl:       "-",
			Name:         fmt.Sprintf("nftName2:%s", "accountName"),
			Description:  fmt.Sprintf("%s `s nft", "accountName"),
			Media:        "xxxxxxxxxxxxxxxx",
			Properties:   []sdk.Propertie{},
			Levels:       []sdk.Level{},
			Stats:        []sdk.Stat{},
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			CollectionId: "1",
			NftUrl:       boundaryStr,
			Name:         fmt.Sprintf("nftName2:%s", "accountName"),
			Description:  fmt.Sprintf("%s `s nft", "accountName"),
			Media:        "collection/dz5hwqaszpwtflg0sfz4",
			Properties:   []sdk.Propertie{},
			Levels:       []sdk.Level{},
			Stats:        []sdk.Stat{},
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			CollectionId: "1",
			NftUrl:       "-",
			Name:         fmt.Sprintf("nftName2:%s", "accountName"),
			Description:  boundaryStr2,
			Media:        "collection/dz5hwqaszpwtflg0sfz4",
			Properties:   []sdk.Propertie{},
			Levels:       []sdk.Level{},
			Stats:        []sdk.Stat{},
		},
		expected: false,
	},
	{
		txinfo: &MintNftTxInfo{
			CollectionId: "1",
			NftUrl:       boundaryStr2,
			Name:         fmt.Sprintf("nftName2:%s", "accountName"),
			Description:  fmt.Sprintf("%s `s nft", "accountName"),
			Media:        "collection/dz5hwqaszpwtflg0sfz4",
			Properties:   []sdk.Propertie{},
			Levels:       []sdk.Level{},
			Stats:        []sdk.Stat{},
		},
		expected: false,
	},
}

func TestMintNft(t *testing.T) {
	tc := getTestingAccountClient(t)
	oAccountClient := tc.oAccountClient
	accountName, _, _ := oAccountClient.GetMyInfo()
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
		_PropertiesByte, err := json.Marshal(test.txinfo.Properties)
		_LevelsByte, err := json.Marshal(test.txinfo.Levels)
		_StatsByte, err := json.Marshal(test.txinfo.Stats)
		_, err = SignAndSendMintNftTx(test.txinfo.CollectionId, test.txinfo.NftUrl, test.txinfo.Name, test.txinfo.Description, test.txinfo.Media, string(_PropertiesByte), string(_LevelsByte), string(_StatsByte), string(data))
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
