package core

import (
	rpcClient "automato/rpc_client"
	yamlParser "automato/yaml_parser"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type AutomationTask interface {
	EvaluateAndExecute(block *types.Block)
}

func GenerateAutomationTasks(ast *yamlParser.YamlFile) []AutomationTask {

	//create new automation task
	automationTasks := []AutomationTask{}
	//add to automation task list
	for _, at := range ast.AutomationTasks {

		newAutomationTask := newAutomationTaskFromASTTrigger(at.Trigger)
		automationTasks = append(automationTasks, newAutomationTask)
	}

	return automationTasks
}

func newAutomationTaskFromASTTrigger(astTrigger *yamlParser.Trigger) AutomationTask {
	if astTrigger.BlockInterval != 0 {
		return newBlockInterval(big.NewInt(int64(astTrigger.BlockInterval)))
	} else if astTrigger.WhenBlock != 0 {
		return newWhenBlock(big.NewInt(int64(astTrigger.WhenBlock)))
	} else if astTrigger.OnEvent != "" {
		return newOnEvent(astTrigger.OnEvent)
	} else {
		return newBlockInterval(big.NewInt(0))
	}
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

		// trigger := automationTask.Trigger
		// action := automationTask.Actions
		// messageContent := MessageContent{}
		// at.Actions.Actions
	}

}
