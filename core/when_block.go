package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type WhenBlock struct {
	executionBlock    *big.Int
	executionFunction func() bool
}

func (w *WhenBlock) EvaluateAndExecute(block *types.Block) {
	if block.Number().Cmp(w.executionBlock) == 0 {
		w.executionFunction()
	}

}
