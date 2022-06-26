package core

import (
	yamlParser "automato/yaml_parser"

	"github.com/ethereum/go-ethereum/core/types"
)

type OnEvent struct {
	EventSignature    bool
	executionFunction func() bool
}

func (o OnEvent) EvaluateAndExecute(block *types.Block) {

}

func newOnEvent(onEvent string, astActions *yamlParser.Actions) OnEvent {

	return OnEvent{}

}
