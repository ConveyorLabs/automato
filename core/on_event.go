package core

import (
	"automato/wallet"
	yamlParser "automato/yaml_parser"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type OnEvent struct {
	EventSignature    bool
	executionFunction func() bool
	Actions           []TX
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
