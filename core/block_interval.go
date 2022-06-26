package core

import (
	"automato/wallet"
	yamlParser "automato/yaml_parser"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

//Type to represent trigger condition that executes every X Blocks
type BlockInterval struct {
	Interval          *big.Int
	LastBlockExecuted *big.Int
	Actions           []TX
}

func (b BlockInterval) EvaluateAndExecute(block *types.Block) {
	if big.NewInt(0).Sub(block.Number(), b.LastBlockExecuted).Cmp(b.Interval) >= 0 {
		for _, action := range b.Actions {

			gas := uint64(0)
			gasTipCap := big.NewInt(0)
			gasFeeCap := big.NewInt(0)

			wallet.Wallet.SignAndSendTx(action.ToAddress, action.Calldata, big.NewInt(0), gas, gasTipCap, gasFeeCap)

		}

	}

}

func newBlockInterval(blockInterval *big.Int, astActions *yamlParser.Actions) BlockInterval {
	//initialize the actions
	actions := []TX{}

	for _, action := range astActions.Actions {

		automationTx := unpackStringToTransaction(action.Tx.Tx)
		actions = append(actions, automationTx)

	}

	return BlockInterval{Interval: blockInterval,
		Actions: actions}
}
