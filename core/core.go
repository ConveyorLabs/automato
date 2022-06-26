package core

import (
	rpcClient "automato/rpc_client"
	yamlParser "automato/yaml_parser"
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

type AutomationTask interface {
	EvaluateAndExecute(*types.Block) bool
}

type Action struct {
	isTX           bool
	messageContent MessageContent
}

type MessageContent struct {
	address           bool
	functionSignature bool
}

func GenerateAutomationTasks(ast *yamlParser.YamlFile) []AutomationTask {

	//create new automation task

	//add to automation task list

	return []AutomationTask{}
}

func StartAutomation(automationTasks []AutomationTask) {
	//create a new block header channel
	blockHeaderChan := make(chan *types.Header)

	//subscribe to block headers
	_, err := rpcClient.WSClient.SubscribeNewHead(context.Background(), blockHeaderChan)
	if err != nil {
		fmt.Println("Error when subscribing to block headers", err)
	}

	//listen for block headers and execute automation tasks
	for {
		blockHeader := <-blockHeaderChan
		block, err := rpcClient.HTTPClient.BlockByHash(context.Background(), blockHeader.Hash())
		if err != nil {
			fmt.Println("Error when getting block by hash", err)
		}

		//loop through automation tasks and execute if the evaluation condition is met
		for _, task := range automationTasks {
			go task.EvaluateAndExecute(block)
		}

	}

}
