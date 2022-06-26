package core

import "math/big"

//Type to represent trigger condition that executes every X Blocks
type BlockInterval struct {
	Interval          *big.Int
	LastBlockExecuted *big.Int
	ExecutionFunction func() bool
}

func (b *BlockInterval) evaluate(blockNumber *big.Int) bool {
	if big.NewInt(0).Sub(blockNumber, b.LastBlockExecuted).Cmp(b.Interval) >= 0 {
		return true

	} else {
		return false
	}

}

func (b *BlockInterval) execute() bool {
	return b.ExecutionFunction()
}
