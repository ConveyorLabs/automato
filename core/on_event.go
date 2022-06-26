package core

import "github.com/ethereum/go-ethereum/core/types"

type OnEvent struct {
<<<<<<< HEAD
	EventSignature    string
	ExecutionFunction func() string
=======
	EventSignature    bool
	executionFunction func() bool
>>>>>>> 53652d3a7b15f88be5d7f81da49f2031a9900d63
}

func (o *OnEvent) EvaluateAndExecute(block *types.Block) {

}
