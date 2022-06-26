package core

import (
	"automato/wallet"
	yamlParser "automato/yaml_parser"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type WhenBlock struct {
	executionBlock    *big.Int
	executionFunction func() bool
	Actions           []TX
}

func (w WhenBlock) EvaluateAndExecute(block *types.Block) {
	if block.Number().Cmp(w.executionBlock) == 0 {
		// w.executionFunction()
		for _, action := range w.Actions {

			wallet.Wallet.SignAndSendTx(action.ToAddress, action.Calldata, big.NewInt(0))

		}
	}

}

func newWhenBlock(executionBlock *big.Int, astActions *yamlParser.Actions) WhenBlock {
	return WhenBlock{executionBlock: executionBlock}
}
