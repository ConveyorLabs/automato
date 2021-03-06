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

			wallet.Wallet.SignAndSendTx(action.ToAddress, action.Calldata, big.NewInt(0))

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
		LastBlockExecuted: big.NewInt(0),

		Actions: actions}
}
