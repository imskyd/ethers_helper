package ethers_helper

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"time"
)

type FilterLogCondition struct {
	Addresses []string
	Topics    []string
}

func (c *FilterLogCondition) FilterQueryAddresses() []common.Address {
	var ads []common.Address
	for _, address := range c.Addresses {
		ads = append(ads, common.HexToAddress(address))
	}
	return ads
}

func (c *FilterLogCondition) FilterQueryGetTopics() []common.Hash {
	var topics []common.Hash
	for _, topic := range c.Topics {
		topics = append(topics, common.HexToHash(topic))
	}
	return topics
}

func FilterBlockLogs(rpc string, fromBlock, toBlock, perSearchRange int64, condition FilterLogCondition, filterFunc func(types.Log)) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatalf("ethclient.Dial err: %s", err.Error())
	}
	var from = fromBlock
	var innerTo = from + perSearchRange

	var to = toBlock
	if toBlock == 0 {
		nowBlock, _ := client.BlockNumber(context.Background())
		to = int64(nowBlock)
	}

	for from <= to {
		q := ethereum.FilterQuery{
			Addresses: condition.FilterQueryAddresses(),
			FromBlock: big.NewInt(from),
			ToBlock:   big.NewInt(innerTo),
		}
		q.Topics = append(q.Topics, condition.FilterQueryGetTopics())

		fLogs, err := client.FilterLogs(context.Background(), q)
		if err != nil {
			log.Printf("client.FilterLogs err: %s", err.Error())
			time.Sleep(time.Second)
			continue
		}

		for _, l := range fLogs {
			filterFunc(l)
		}
		if from+perSearchRange > to {
			from += perSearchRange + 1
			innerTo = to
		} else {
			from += perSearchRange + 1
			innerTo = from + perSearchRange
		}
	}
}
