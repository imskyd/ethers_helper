package example

import (
	"fmt"
	"github.com/imskyd/ethers_helper"
	"log"
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
