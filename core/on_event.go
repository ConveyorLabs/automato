package core

import (
	"automato/wallet"
	yamlParser "automato/yaml_parser"
	"math/big"
	"strings"

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
	for _, action := range o.Actions {

		gas := uint64(0)
		gasTipCap := big.NewInt(0)
		gasFeeCap := big.NewInt(0)

		wallet.Wallet.SignAndSendTx(action.ToAddress, action.Calldata, big.NewInt(0), gas, gasTipCap, gasFeeCap)

	}

}

func newOnEvent(onEvent string, astActions *yamlParser.Actions) OnEvent {

	return OnEvent{}

}
