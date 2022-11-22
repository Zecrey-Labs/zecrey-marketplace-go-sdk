package sdk

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strings"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zecrey-labs/zecrey-crypto/zecrey/twistededwards/tebn254/zecrey"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	AccountNameSuffix = ".zec"
	PkBytes           = 64
)

var Modulus = fr.Modulus()

func ComputeAccountNameHashInBytes(accountName string) (res []byte, err error) {
	words := strings.Split(accountName, ".")
	if len(words) != 2 {
		return nil, errors.New("[AccountNameHash] invalid account name")
	}
	rootNode := make([]byte, 32)
	hashOfBaseNode := KeccakHash(append(rootNode, KeccakHash([]byte(words[1]))...))

	baseNode := big.NewInt(0).Mod(big.NewInt(0).SetBytes(hashOfBaseNode), Modulus)
	baseNodeBytes := make([]byte, 32)
	baseNode.FillBytes(baseNodeBytes)

	nameHash := KeccakHash([]byte(words[0]))
	subNameHash := KeccakHash(append(baseNodeBytes, nameHash...))

	subNode := big.NewInt(0).Mod(big.NewInt(0).SetBytes(subNameHash), Modulus)
	subNodeBytes := make([]byte, 32)
	subNode.FillBytes(subNodeBytes)

	return subNodeBytes, nil
}

func SetFixed32Bytes(buf []byte) [32]byte {
	newBuf := new(big.Int).SetBytes(buf).FillBytes(make([]byte, zecrey.PointSize))
	var res [zecrey.PointSize]byte
	copy(res[:], newBuf[:])
	return res
}

func PubKeyStrToPxAndPy(pkStr string) (px [32]byte, py [32]byte, err error) {
	pkBytes, err := PubKeyStrToBytes64(pkStr)
	if err != nil {
		return px, py, err
	}
	px = SetFixed32Bytes(pkBytes[:32])
	py = SetFixed32Bytes(pkBytes[32:])
	return px, py, nil
}

func PubKeyStrToBytes64(pkStr string) (pkBytes []byte, err error) {
	pkBytes, err = hex.DecodeString(pkStr)
	if err != nil {
		logx.Errorf("[PubKeyStrToBytes64] unable to decode pk str: %s", err.Error())
		return nil, err
	}
	if len(pkBytes) != PkBytes {
		logx.Errorf("[PubKeyStrToBytes64] invalid pk")
		return nil, errors.New("[PubKeyStrToBytes64] invalid pk")
	}
	return pkBytes, nil
}

func KeccakHash(value []byte) []byte {
	hashVal := crypto.Keccak256Hash(value)
	return hashVal[:]
}
