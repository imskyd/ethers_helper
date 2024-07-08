package ethers_helper

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func AddressMustBelongToPrivateKey(address, _pk string) (bool, error, string) {
	privateKey, err := crypto.HexToECDSA(_pk)
	if err != nil {
		return false, err, ""
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey"), ""
	}

	pkToAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return strings.EqualFold(pkToAddress, address), nil, pkToAddress
}

func GetAddressFromPrivateKey(_pk string) (string, error) {
	privateKey, err := crypto.HexToECDSA(_pk)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", nil
	}

	pkToAddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return pkToAddress, nil
}

func GetFromAddressByTransaction(signer types.Signer, tx *types.Transaction) (common.Address, error) {
	if signer == nil {
		signer = types.NewCancunSigner(tx.ChainId())
	}
	from, err := types.Sender(signer, tx)
	return from, err
}
