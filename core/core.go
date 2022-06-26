package core

import (
	rpcClient "automato/rpc_client"
	yamlParser "automato/yaml_parser"
	"context"
	"fmt"
	"math/big"
	"reflect"

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
	address           string
	functionSignature string
}

func GenerateAutomationTasks(ast *yamlParser.YamlFile) []AutomationTask {

	//create new automation task

	//add to automation task list
	automationTasks := []AutomationTask{}
	for _, at := range ast.AutomationTasks {
		automationTask := at

		if reflect.DeepEqual(automationTask.Trigger, at.Trigger.BlockInterval) {
			newBlockInterval := BlockInterval{}
			newBlockInterval.Interval = big.NewInt(automationTask.Trigger.BlockInterval)
			// for _, action := range automationTask.Actions.Actions {
			// 	if reflect.DeepEqual(action, at.Actions.Actions.Tx) {

			// 	}
			// }

			newAction := Action{}
			newAction.isTX = true
			newMessageContent := MessageContent{}
			newMessageContent.address = automationTask.Actions.Actions[0].Tx.Tx[:32]
			newMessageContent.functionSignature = automationTask.Actions.Actions[0].Tx.Tx[32:]
			newAction.messageContent = newMessageContent

		}

		if reflect.DeepEqual(automationTask.Trigger, at.Trigger.OnEvent) {

		}

		if reflect.DeepEqual(automationTask.Trigger, at.Trigger.WhenBlock) {

		}

	}

	return automationTasks
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
