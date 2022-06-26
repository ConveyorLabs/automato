package core

import (
	rpcClient "automato/rpc_client"
	"automato/wallet"
	yamlParser "automato/yaml_parser"
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type OnEvent struct {
	EventSignature    bool
	executionFunction func() bool
	Actions           []TX
	TopicAddress      *common.Address
	TopicHash         common.Hash
}

func unpackTopicHash(hash string) (*common.Address, common.Hash) {
	hashSplit := strings.Split(hash, "(")
	address := common.HexToAddress(hashSplit[0])
	topicHash := common.HexToHash(strings.Split(hashSplit[1], ")")[0])
	return &address, topicHash
}

func (o OnEvent) EvaluateAndExecute(block *types.Block) {

	blockNumber, err := rpcClient.HTTPClient.BlockNumber(context.Background())

	if err != nil {
		fmt.Println("error when getting block number", err)
		//graceful error handling incoming
		os.Exit(1)
	}
	//create a filter query on the contract address
	filter := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(blockNumber)),
		Addresses: []common.Address{*o.TopicAddress},
		//add transfer event signature as a topic
		Topics: [][]common.Hash{{o.TopicHash}},
	}

	filterLogs, filterLogError := rpcClient.HTTPClient.FilterLogs(context.Background(), filter)
	if filterLogError != nil {
		fmt.Println("error when getting block number", err)
		//graceful error handling incoming
		os.Exit(1)
	}

	if len(filterLogs) > 0 {

		for _, action := range o.Actions {

			wallet.Wallet.SignAndSendTx(action.ToAddress, action.Calldata, big.NewInt(0))

		}
	}

}

func newOnEvent(onEvent string, astActions *yamlParser.Actions) OnEvent {

	return OnEvent{}

}
