package core

import "github.com/ethereum/go-ethereum/core/types"

type OnEvent struct {
	EventSignature    bool
	executionFunction func() bool
}

func (o *OnEvent) EvaluateAndExecute(block *types.Block) {

}
