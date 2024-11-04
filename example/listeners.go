package example

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/imskyd/ethers_helper"
	"log"
	"strings"
)

func MemPoolListen() {
	memPoolTxCh, sub, err := ethers_helper.ListenMemPool("")
	if err != nil {
		log.Panicf(err.Error())
	}

	for {
		select {
		case subErr := <-sub.Err():
			fmt.Println(subErr)
		case tx := <-memPoolTxCh:
			fmt.Println(tx)
			//tx data filter
			if len(tx.Data()) < 8 {
				continue
			}
			//method filter
			methodB := tx.Data()[:4]
			method := common.Bytes2Hex(methodB)
			if method != "00cb637e" {
				continue
			}

			//from address filter
			from, err := types.Sender(types.NewCancunSigner(tx.ChainId()), tx)
			if err != nil {
				from, err := types.Sender(types.HomesteadSigner{}, tx)
				fmt.Println(from, err, tx.Hash().String())
			}
			if strings.EqualFold(from.String(), "0x000000000000092a51384334D8090DD869a816CB") {
				//method data filter
				data, err := ethers_helper.DecodeTxData("abi", common.Bytes2Hex(tx.Data()))
				fmt.Println(data, err)
			}
		}
	}
}

func BlockListen() {
	headChan, sub, err := ethers_helper.ListenBlock("")
	if err != nil {
		log.Panicf(err.Error())
	}

	for {
		select {
		case subErr := <-sub.Err():
			fmt.Println(subErr)
		case header := <-headChan:
			fmt.Println(header)
		}
	}
}

func EventListen() {
	var condition ethers_helper.FilterLogCondition
	logChan, sub, err := ethers_helper.ListenEvent("", condition)
	if err != nil {
		log.Panicf(err.Error())
	}

	for {
		select {
		case subErr := <-sub.Err():
			fmt.Println(subErr)
		case l := <-logChan:
			fmt.Println(l)
		}
	}
}
