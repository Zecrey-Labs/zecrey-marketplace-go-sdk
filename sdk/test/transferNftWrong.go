package main

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
)

type TransferNftTxInfo struct {
	ToAccountName,
	AssetId string
}

var transferNftTestCase = []struct {
	txinfo   *TransferNftTxInfo
	expected bool
}{
	{
		txinfo: &TransferNftTxInfo{
			ToAccountName: cfg.BoundaryStr2,
			AssetId:       "",
		},
		expected: false,
	},
}

func transferNftWrongBetch(index int) {
	for i := 0; i < index; i++ {
		for j := 0; j < i*10000; j++ {
			go transferNftWrong(index)
		}
	}
}

func transferNftWrong(index int) {
	accountName, _, _ := client.GetMyInfo()
	for _, test := range transferNftTestCase {
		test.txinfo.AssetId = fmt.Sprintf("%d", rand.Intn(10000))
		resultSdk, err := getPreTransferNftTx(accountName, test.txinfo.ToAccountName, test.txinfo.AssetId)
		_, err = SignAndSendTransferNftTx(client.GetKeyManager(), test.txinfo.AssetId, resultSdk.Transtion)
		if test.expected {
			fmt.Println(fmt.Sprintf("fail %t! txType=%s,index=%d,func=%s,err=%s", test.expected, "transferNftWrong", index, "MintNft", err.Error()))
			return
		} else {
			fmt.Println(fmt.Sprintf("fail %t! txType=%s,index=%d,func=%s,err=%s", test.expected, "transferNftWrong", index, "MintNft", err.Error()))
			return
		}
	}
}

func SignAndSendTransferNftTx(keyManager sdk.KeyManager, AssetId, txInfoSdk string) (*sdk.RespSendTransferNft, error) {
	txInfo, err := sdkTransferNftTxInfo(keyManager, txInfoSdk)
	resp, err := http.PostForm(nftMarketUrl+"/api/v1/asset/sendTransferNft",
		url.Values{
			"asset_id":    {fmt.Sprintf("%s", AssetId)},
			"transaction": {txInfo},
		},
	)
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
	result := &sdk.RespSendTransferNft{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func getPreTransferNftTx(accountName, toAccountName, AssetId string) (*sdk.RespetSdktxInfo, error) {
	respSdkTx, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkTransferNftTxInfo?account_name=%s&to_account_name=%s%s&nft_id=%d", accountName, toAccountName, NameSuffix, AssetId))
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
	return nil, err
}
func sdkTransferNftTxInfo(key sdk.KeyManager, txInfoSdk string) (string, error) {
	txInfo := &sdk.TransferNftTxInfo{}
	err := json.Unmarshal([]byte(txInfoSdk), txInfo)
	if err != nil {
		return "", err
	}
	txInfo.GasFeeAssetAmount = big.NewInt(MinGasFee)
	tx, err := constructTransferNftTx(key, txInfo)
	if err != nil {
		return "", err
	}
	return tx, err
}
func constructTransferNftTx(key sdk.KeyManager, tx *sdk.TransferNftTxInfo) (string, error) {
	convertedTx := convertTransferNftTxInfo(tx)
	hFunc := mimc.NewMiMC()
	msgHash, err := legendTxTypes.ComputeTransferNftMsgHash(convertedTx, hFunc)
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
func convertTransferNftTxInfo(tx *sdk.TransferNftTxInfo) *legendTxTypes.TransferNftTxInfo {
	return &legendTxTypes.TransferNftTxInfo{
		FromAccountIndex:  tx.FromAccountIndex,
		ToAccountIndex:    tx.ToAccountIndex,
		ToAccountNameHash: tx.ToAccountNameHash,
		NftIndex:          tx.NftIndex,
		GasAccountIndex:   tx.GasAccountIndex,
		GasFeeAssetId:     tx.GasFeeAssetId,
		GasFeeAssetAmount: tx.GasFeeAssetAmount,
		CallData:          tx.CallData,
		CallDataHash:      tx.CallDataHash,
		ExpiredAt:         tx.ExpiredAt,
		Nonce:             tx.Nonce,
		Sig:               tx.Sig,
	}
}
