package ethers_helper

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"strings"
)

func MakeTxDataByAbi(abiJson, method string, args ...interface{}) ([]byte, error) {
	parsed, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, err
	}
	pack, err2 := parsed.Pack(method, args...)
	if err2 != nil {
		return nil, err
	}
	return pack, nil
}

func DecodeTxData(abiString string, data string) (map[string]interface{}, error) {
	contract, err := abi.JSON(strings.NewReader(abiString))
	if err != nil {
		return nil, err
	}
	methodBD, err := contract.MethodById(common.Hex2Bytes(data[2:10]))
	if err != nil {
		return nil, err
	}
	unpack := make(map[string]interface{})
	data2, decodeErr := hex.DecodeString(data[10:])
	if decodeErr != nil {
		return nil, decodeErr
	}
	unpackErr := methodBD.Inputs.UnpackIntoMap(unpack, data2)
	return unpack, unpackErr
}
