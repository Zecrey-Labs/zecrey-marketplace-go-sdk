package test

import (
	"encoding/json"
	"fmt"
	"github.com/Zecrey-Labs/zecrey-marketplace-go-sdk/sdk"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/stretchr/testify/assert"
	"github.com/zecrey-labs/zecrey-crypto/wasm/zecrey-legend/legendTxTypes"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"testing"
)

type TransferNftTxInfo struct {
	AccountName, ToAccountName, AssetId string
}

var transferNftTestCase = []struct {
	txinfo   *TransferNftTxInfo
	expected bool
}{
	{
		txinfo: &TransferNftTxInfo{
			//FromAccountIndex:  0,
			//ToAccountIndex:    0,
			//ToAccountNameHash: "0x0000000000000000000000000000000000000000000000000000000000000000",
			//NftIndex:          0,
			//GasAccountIndex:   0,
			//GasFeeAssetId:     0,
			//GasFeeAssetAmount: big.NewInt(0),
			//CallData:          ";DROP TABLE account;",
			//ExpiredAt:         time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:             0,
		},
		expected: false,
	},
	{
		txinfo: &TransferNftTxInfo{
			//FromAccountIndex:  -1,
			//ToAccountIndex:    -1,
			//ToAccountNameHash: ";DROP TABLE account;",
			//NftIndex:          -1,
			//GasAccountIndex:   -1,
			//GasFeeAssetId:     -1,
			//GasFeeAssetAmount: big.NewInt(-1),
			//CallData:          ";DROP TABLE account;",
			//ExpiredAt:         time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:             0,
		},
		expected: false,
	},
	{
		txinfo: &TransferNftTxInfo{
			//FromAccountIndex:  math.MaxInt64,
			//ToAccountIndex:    math.MaxInt64,
			//ToAccountNameHash: string([]byte{math.MaxUint8}),
			//NftIndex:          math.MaxInt64,
			//GasAccountIndex:   math.MaxInt64,
			//GasFeeAssetId:     math.MaxInt64,
			//GasFeeAssetAmount: big.NewInt(0).Mul(big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64)),
			//CallData:          ";DROP TABLE account;",
			//ExpiredAt:         time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:             0,
		},
		expected: false,
	},
	{
		txinfo: &TransferNftTxInfo{
			//FromAccountIndex:  math.MinInt64,
			//ToAccountIndex:    math.MinInt64,
			//ToAccountNameHash: uuid.New().String(),
			//NftIndex:          math.MinInt64,
			//GasAccountIndex:   math.MinInt64,
			//GasFeeAssetId:     math.MinInt64,
			//GasFeeAssetAmount: big.NewInt(0),
			//CallData:          ";DROP TABLE account;",
			//ExpiredAt:         time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:             0,
		},
		expected: false,
	},
	{
		txinfo: &TransferNftTxInfo{
			//FromAccountIndex:  143363,
			//ToAccountIndex:    58701,
			//ToAccountNameHash: "18e0b84767aa85d642f191710f662b4e8e4a3586502a97345c24bcc92f530a9a",
			//NftIndex:          -1,
			//GasAccountIndex:   -1,
			//GasFeeAssetId:     -1,
			//GasFeeAssetAmount: big.NewInt(-1),
			//CallData:          ";DROP TABLE account;",
			//ExpiredAt:         time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:             0,
		},
		expected: false,
	},
	{
		txinfo: &TransferNftTxInfo{
			//FromAccountIndex:  56706,
			//ToAccountIndex:    14332,
			//ToAccountNameHash: "<<<<<<<" + string([]byte{math.MaxUint8}),
			//NftIndex:          -1,
			//GasAccountIndex:   -1,
			//GasFeeAssetId:     -1,
			//GasFeeAssetAmount: big.NewInt(-1),
			//CallData:          ";DROP TABLE account;",
			//ExpiredAt:         time.Now().Add(24 * time.Hour).UnixMilli(),
			//Nonce:             0,
		},
		expected: false,
	},
}

func TestTransferNft(t *testing.T) {
	// transfer nft
	tc := getTestingAccountClient(t)
	oAccountClient := tc.oAccountClient

	//nonce, err := oAccountClient.GetNextNonce(oAccountInfo.AccountIndex)
	//assert.Nil(t, err, "GetNextNonce should not return an error, err: %v", err)
	//assert.Greater(t, nonce, int64(0), "nonce should be greater than 0")
	//gasFee, err := oAccountClient.GetGasFee(0, sdk.TxTypeCreateCollection)
	//assert.Nil(t, err, "GetGasFeeByAssetIdAndAccountIndex should not return an error, err: %v", err)
	//assert.Greater(t, gasFee, int64(0), "gasFee should be greater than 0")

	//nfts, err = oAccountClient.GetAccountNftList(uint32(nAccountInfo.AccountIndex), 0, 100)
	//assert.Nil(t, err, "GetAccountNftList failed")
	//found := false
	//for _, nft := range nfts.Nfts {
	//	if nft.NftIndex == transferNft.NftIndex {
	//		found = true
	//		break
	//	}
	//}
	//if !found && nfts.Total > 100 {
	//	nfts, err = oAccountClient.GetAccountNftList(uint32(nAccountInfo.AccountIndex), 100, uint32(nfts.Total))
	//	assert.Nil(t, err, "GetAccountNftList failed")
	//	for _, nft := range nfts.Nfts {
	//		if nft.NftIndex == transferNft.NftIndex {
	//			found = true
	//			break
	//		}
	//	}
	//}
	//assert.Equal(t, true, found, "nft should be found")

	for _, test := range transferNftTestCase {
		//nonce, err = oAccountClient.GetNextNonce(oAccountInfo.AccountIndex)
		//assert.Nil(t, err, "GetNextNonce should not return an error, err: %v", err)
		//assert.Greater(t, nonce, int64(0), "nonce should be greater than 0")
		//gasFee, err = oAccountClient.GetGasFee(0, sdk.TxTypeTransferNft)
		//assert.Nil(t, err, "GetGasFeeByAssetIdAndAccountIndex should not return an error, err: %v", err)
		//assert.Greater(t, gasFee, int64(0), "gasFee should be greater than 0")
		resultSdk, err := getPreTransferNftTx(test.txinfo.AccountName, test.txinfo.ToAccountName, test.txinfo.AssetId)
		_, err = SignAndSendTransferNftTx(oAccountClient.GetKeyManager(), test.txinfo.AssetId, resultSdk.Transtion)
		if test.expected {
			assert.Nil(t, err, "SignAndSendTransferNftTx should not return an error")
		} else {
			assert.NotNil(t, err, "SignAndSendTransferNftTx should return an error")
		}
	}
}

func SignAndSendTransferNftTx(keyManager sdk.KeyManager, txInfoSdk, AssetId string) (*sdk.ResqSendTransferNft, error) {
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
	result := &sdk.ResqSendTransferNft{}
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
