package sdk

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zecrey-labs/zecrey-crypto/util/ecdsaHelper"
	"github.com/zecrey-labs/zecrey-crypto/util/eddsaHelper"
	"github.com/zecrey-labs/zecrey-eth-rpc/_rpc"
	zecreyLegendRpc "github.com/zecrey-labs/zecrey-eth-rpc/zecrey/core/zecrey-legend"
	zecreyLegendUtil "github.com/zecrey-labs/zecrey-legend/common/util"
	"github.com/zeromicro/go-zero/core/logx"
)

func GetAccountL1Address(accountName string) (common.Address, error) {
	providerClient, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		return bytesToAddress([]byte{}), fmt.Errorf(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	res, err := zecreyLegendUtil.ComputeAccountNameHashInBytes(accountName + NameSuffix)
	if err != nil {
		logx.Error(err)
		return bytesToAddress([]byte{}), err
	}
	//get base contract address
	resp, err := GetLayer2BasicInfo()
	if err != nil {
		return bytesToAddress([]byte{}), err
	}
	ZecreyLegendContract := resp.ContractAddresses[0]
	//ZnsPriceOracle := resp.ContractAddresses[1]

	resBytes := zecreyLegendUtil.SetFixed32Bytes(res)
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(providerClient, ZecreyLegendContract)
	if err != nil {
		return bytesToAddress([]byte{}), err
	}
	// fetch by accountNameHash
	addr, err := zecreyInstance.GetAddressByAccountNameHash(zecreyLegendRpc.EmptyCallOpts(), resBytes)
	if err != nil {
		logx.Error(err)
		return bytesToAddress([]byte{}), err
	}
	if bytes.Equal(addr.Bytes(), bytesToAddress([]byte{}).Bytes()) {
		return bytesToAddress([]byte{}), fmt.Errorf("null address")
	}
	return addr, nil
}

func GetLayer2BasicInfo() (*RespGetLayer2BasicInfo, error) {
	resp, err := http.Get(legendUrl + "/api/v1/info/getLayer2BasicInfo")
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
	result := &RespGetLayer2BasicInfo{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetAccountByAccountName(accountName string) (*RespGetAccountByAccountName, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/account/getAccountByAccountName?account_name=%s", accountName))
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
	result := &RespGetAccountByAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetNextNonce(accountIdx int64) (int64, error) {
	resp, err := http.Get(legendUrl +
		fmt.Sprintf("/api/v1/tx/getNextNonce?account_index=%d", accountIdx))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetNextNonce{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.Nonce, nil
}

func GetAccountIndex(accountName string) (int64, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/account/getAccountByAccountName?account_name=%s", accountName))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(string(body))
	}
	result := &RespGetAccountByAccountName{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.Account.AccountIndex, nil
}

func IfAccountRegistered(accountName string) (bool, error) {
	c, err := newZecreyMarketplaceClientDefault(accountName)
	if err != nil {
		logx.Error(err)
		return false, err
	}
	res, err := zecreyLegendUtil.ComputeAccountNameHashInBytes(accountName + NameSuffix)
	if err != nil {
		logx.Error(err)
		return false, err
	}
	//get base contract address
	resp, err := GetLayer2BasicInfo()
	if err != nil {
		return false, err
	}
	ZecreyLegendContract := resp.ContractAddresses[0]
	//ZnsPriceOracle := resp.ContractAddresses[1]

	resBytes := zecreyLegendUtil.SetFixed32Bytes(res)
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return false, err
	}
	// fetch by accountNameHash
	addr, err := zecreyInstance.GetAddressByAccountNameHash(zecreyLegendRpc.EmptyCallOpts(), resBytes)
	if err != nil {
		logx.Error(err)
		return false, err
	}
	return !bytes.Equal(addr.Bytes(), bytesToAddress([]byte{}).Bytes()), nil
}

func GetCategories() (*RespGetCollectionCategories, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/collection/getCollectionCategories"))
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
	result := &RespGetCollectionCategories{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetCollectionById(collectionId int64) (*RespGetSdkCollectionById, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkCollectionById?collection_id=%d", collectionId))
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
	result := &RespGetSdkCollectionById{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetCollectionsByAccountIndex(AccountIndex int64) (*RespGetSdkAccountCollections, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAccountCollections?account_index=%d", AccountIndex))
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
	result := &RespGetSdkAccountCollections{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	//get collections status = 1
	return result, nil
}

func GetAccountNFTs(AccountIndex int64) (*RespGetSdkAccountAssets, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAccountAssets?account_index=%d", AccountIndex))
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
	result := &RespGetSdkAccountAssets{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	//get collections status = 1
	return result, nil
}

func GetAccountOffers(AccountIndex int64) (*RespGetSdkAccountOffers, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAccountOffers?account_index=%d", AccountIndex))
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
	result := &RespGetSdkAccountOffers{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	//get collections status = 1
	return result, nil
}

func GetNftOffers(NftId int64) (*RespGetSdkAssetOffers, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAssetOffers?asset_id=%d", NftId))
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
	result := &RespGetSdkAssetOffers{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	//get collections status = 1
	return result, nil
}

func GetNftById(nftId int64) (*RespGetSdkAssetById, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/sdk/getSdkAssetByAssetId?asset_id=%d", nftId))
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
	result := &RespGetSdkAssetById{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetNextOfferId(AccountName string) (*RespGetNextOfferId, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/offer/getNextOfferId?account_name=%s", AccountName))
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
	result := &RespGetNextOfferId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetOfferById(OfferId int64) (*RespGetOfferByOfferId, error) {
	resp, err := http.Get(nftMarketUrl + fmt.Sprintf("/api/v1/offer/getOfferByOfferId?offer_id=%d", OfferId))
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
	result := &RespGetOfferByOfferId{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func GetListingOffers(isSell int64) (*RespGetNftBeingSell, error) {
	queryStr := fmt.Sprintf(`
{"query":"query MyQuery {\n  offer(where: {status: {_eq: \"%d\"}, direction: {_eq: \"1\"}}) {\n    id\n    l2_offer_id\n    asset_id\n    counterpart_id\n    payment_asset_id\n    payment_asset_amount\n    signature\n    status\n    direction\n    expired_at\n   created_at\n    asset {\n      id\n      nft_index\n      name\n      collection_id\n      content_hash\n      create_tx_hash\n      creator_earning_rate\n      description\n      expired_at\n      image_thumb\n      l1_token_id\n      last_payment_asset_amount\n      last_payment_asset_id\n      media_detail {\n        url\n      }\n      nft_url\n      status\n      video_thumb\n      asset_stats {\n        max_value\n        key\n      }\n      asset_properties {\n        key\n        value\n      }\n      asset_levels {\n        key\n        max_value\n        value\n      }\n    }\n  }\n}\n","variables":{}}
`, isSell)

	var data = []byte(queryStr)
	body, err := post2Hasura(data)
	if err != nil {
		return nil, err
	}

	result := &RespGetNftBeingSell{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func post2Hasura(data []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, hasuraUrl, bytes.NewReader(data))
	if err != nil {
		return []byte(""), err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-access-key", hasuraAdminKey)
	hc := http.DefaultClient
	hc.Timeout = time.Second * hasuraTimeDeadline
	resp, err := hc.Do(req)
	if err != nil {
		return []byte(""), err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}
	if resp.StatusCode != http.StatusOK {
		return []byte(""), fmt.Errorf(string(body))
	}
	return body, nil
}

func UploadMedia(filePath string) (*RespMediaUpload, error) {
	uri := fmt.Sprintf(nftMarketUrl+"%s", "/api/v1/asset/media")
	paramName := "image"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filePath)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		panic(err)
	}
	err = writer.Close()
	if err != nil {
		panic(err)
	}
	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	clt := http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
	}
	defer clt.CloseIdleConnections()
	res, err := clt.Do(request)
	defer func() {
		res.Body.Close()
	}()
	if err != nil {
		panic(fmt.Errorf("err is %s", err.Error()))
	}
	body1, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil {
		panic(fmt.Errorf("ioutil.ReadAll  err is %s", err1.Error()))
	}
	result := &RespMediaUpload{}
	if err := json.Unmarshal(body1, &result); err != nil {
		return nil, err
	}
	return result, nil
}

//newZecreyMarketplaceClientDefault private
func newZecreyMarketplaceClientWithSeed(accountName, seed string) (*Client, error) {
	keyManager, err := NewSeedKeyManager(seed)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong seed:%s", seed))
	}
	l2pk := eddsaHelper.GetEddsaPublicKey(seed[2:])
	connEth, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	return &Client{
		accountName:    fmt.Sprintf("%s%s", accountName, NameSuffix),
		seed:           seed,
		l2pk:           l2pk,
		nftMarketUrl:   nftMarketUrl,
		legendUrl:      legendUrl,
		providerClient: connEth,
		keyManager:     keyManager,
	}, nil
}

//newZecreyMarketplaceClientDefault private
func newZecreyMarketplaceClientDefault(accountName string) (*Client, error) {
	connEth, err := _rpc.NewClient(chainRpcUrl)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("wrong rpc url:%s", chainRpcUrl))
	}
	return &Client{
		accountName:    fmt.Sprintf("%s%s", accountName, NameSuffix),
		nftMarketUrl:   nftMarketUrl,
		legendUrl:      legendUrl,
		providerClient: connEth,
	}, nil
}

func CreateL1Account() (l1Addr, privateKeyStr, l2pk, seed string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		logx.Errorf("[CreateL1Account] GenerateKey err: %s", err)
		return "", "", "", "", err
	}
	privateKeyStr = hex.EncodeToString(crypto.FromECDSA(privateKey))
	l1Addr, err = ecdsaHelper.GenerateL1Address(privateKey)
	if err != nil {
		logx.Errorf("[CreateL1Account] GenerateL1Address err: %s", err)
		return "", "", "", "", err
	}
	seed, err = eddsaHelper.GetEddsaSeed(privateKey)
	if err != nil {
		logx.Errorf("[CreateL1Account] GetEddsaSeed err: %s", err)
		return "", "", "", "", err
	}
	l2pk = eddsaHelper.GetEddsaPublicKey(seed[2:])
	return
}

func GetSeedAndL2Pk(privateKeyStr string) (l2pk, seed string, err error) {
	privECDSA, err := crypto.ToECDSA(common.FromHex(privateKeyStr))
	seed, err = eddsaHelper.GetEddsaSeed(privECDSA)
	if err != nil {
		logx.Errorf("[CreateL1Account] GetEddsaSeed err: %s", err)
		return "", "", err
	}
	l2pk = eddsaHelper.GetEddsaPublicKey(seed[2:])
	return
}

func RegisterAccountWithPrivateKey(accountName, l1Addr, privateKey string) (*Client, error) {
	l2pk, seed, err := GetSeedAndL2Pk(privateKey)
	if err != nil {
		return nil, err
	}
	c, err := newZecreyMarketplaceClientWithSeed(accountName, seed)
	if err != nil {
		return nil, err
	}
	if ok, err := IfAccountRegistered(accountName); ok {
		if err != nil {
			return nil, err
		}
		return NewClient(accountName, seed)
	}
	var chainId *big.Int
	chainId, err = c.providerClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	authCli, err := _rpc.NewAuthClient(c.providerClient, privateKey, chainId)
	if err != nil {
		return nil, err
	}
	px, py, err := zecreyLegendUtil.PubKeyStrToPxAndPy(l2pk)
	if err != nil {
		return nil, err
	}
	//get base contract address
	resp, err := GetLayer2BasicInfo()
	if err != nil {
		return nil, err
	}
	ZecreyLegendContract := resp.ContractAddresses[0]
	ZnsPriceOracle := resp.ContractAddresses[1]

	gasPrice, err := c.providerClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	zecreyInstance, err := zecreyLegendRpc.LoadZecreyLegendInstance(c.providerClient, ZecreyLegendContract)
	if err != nil {
		return nil, err
	}
	priceOracleInstance, err := zecreyLegendRpc.LoadStablePriceOracleInstance(c.providerClient, ZnsPriceOracle)
	if err != nil {
		return nil, err
	}
	_, err = zecreyLegendRpc.RegisterZNS(c.providerClient, authCli,
		zecreyInstance, priceOracleInstance,
		gasPrice, DefaultGasLimit, accountName,
		common.HexToAddress(l1Addr), px, py)
	if err != nil {
		return nil, err
	}
	return NewClient(accountName, seed)
}

func ApplyRegisterHost(
	accountName string, l2Pk string, OwnerAddr string) (*RespApplyRegisterHost, error) {
	resp, err := http.PostForm(legendUrl+"/api/v1/register/applyRegisterHost",
		url.Values{
			"account_name": {accountName},
			"l2_pk":        {l2Pk},
			"owner_addr":   {OwnerAddr}})
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
	result := &RespApplyRegisterHost{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func bytesToAddress(b []byte) common.Address {
	var a common.Address
	a.SetBytes(b)
	return a
}

func SignMessage(key KeyManager, message string) string {
	sig, err := key.Sign([]byte(message), mimc.NewMiMC())
	if err != nil {
		panic("failed to sign message, err: " + err.Error())
	}
	signed := hex.EncodeToString(sig[:])
	return signed
}
