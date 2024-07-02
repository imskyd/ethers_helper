package ethers_helper

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func FilterBlockLogs(rpc string, fromBlock, toBlock, perSearchRange int64, address, topic string, filterFunc func(types.Log)) {
	client, _ := ethclient.Dial(rpc)

	var from = fromBlock
	var innerTo = from + perSearchRange

	var to = toBlock
	if toBlock == 0 {
		nowBlock, _ := client.BlockNumber(context.Background())
		to = int64(nowBlock)
	}

	for from <= to {
		q := ethereum.FilterQuery{
			Addresses: []common.Address{common.HexToAddress(address)},
			FromBlock: big.NewInt(from),
			ToBlock:   big.NewInt(innerTo),
		}
		q.Topics = append(q.Topics, []common.Hash{common.HexToHash(topic)})

		fLogs, err := client.FilterLogs(context.Background(), q)
		if err != nil {
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
