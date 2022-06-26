package core

import "math/big"

//Type to represent trigger condition that executes every X Blocks
type BlockInterval struct {
	TaskType          string
	Interval          *big.Int
	LastBlockExecuted *big.Int
	executionFunction func() bool
}

func (b *BlockInterval) evaluateAndExecute(blockNumber *big.Int) {
	if big.NewInt(0).Sub(blockNumber, b.LastBlockExecuted).Cmp(b.Interval) >= 0 {
		b.executionFunction()
	}

}
