package ethers_helper

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func ListenMemPool(wssRpc string) (chan *types.Transaction, *rpc.ClientSubscription, error) {
	rpcCli, err := rpc.Dial(wssRpc)
	if err != nil {
		return nil, nil, err
	}
	gCli := gethclient.New(rpcCli)
	txChan := make(chan *types.Transaction, 1000)
	sub, subscribeErr := gCli.SubscribeFullPendingTransactions(context.Background(), txChan)
	return txChan, sub, subscribeErr
}

func ListenBlock(wssRpc string) (chan *types.Header, ethereum.Subscription, error) {
	client, dialErr := ethclient.Dial(wssRpc)
	if dialErr != nil {
		return nil, nil, dialErr
	}
	headers := make(chan *types.Header)
	sub, subErr := client.SubscribeNewHead(context.Background(), headers)
	return headers, sub, subErr
}

func ListenEvent(wssRpc string, condition FilterLogCondition) (chan types.Log, ethereum.Subscription, error) {
	client, dialErr := ethclient.Dial(wssRpc)
	if dialErr != nil {
		return nil, nil, dialErr
	}
	q := ethereum.FilterQuery{
		Addresses: condition.FilterQueryAddresses(),
	}
	q.Topics = append(q.Topics, condition.FilterQueryGetTopics())
	chann := make(chan types.Log)
	sub, subErr := client.SubscribeFilterLogs(context.Background(), q, chann)
	return chann, sub, subErr
}
