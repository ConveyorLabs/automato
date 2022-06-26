package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

//Type to represent trigger condition that executes every X Blocks
type BlockInterval struct {
	Interval          *big.Int
	LastBlockExecuted *big.Int
}

func (b BlockInterval) EvaluateAndExecute(block *types.Block) {
	if big.NewInt(0).Sub(block.Number(), b.LastBlockExecuted).Cmp(b.Interval) >= 0 {

	}

}

func newBlockInterval(blockInterval *big.Int) BlockInterval {
	return BlockInterval{Interval: blockInterval}
}
