package core

<<<<<<< HEAD
import "math/big"

type WhenBlock struct {
	executionBlock *big.Int
=======
import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type WhenBlock struct {
	executionBlock    *big.Int
	executionFunction func() bool
>>>>>>> 53652d3a7b15f88be5d7f81da49f2031a9900d63
}

func (w *WhenBlock) EvaluateAndExecute(block *types.Block) {
	if block.Number().Cmp(w.executionBlock) == 0 {
		w.executionFunction()
	}

}
