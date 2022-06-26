package core

import (
	rpcClient "automato/rpc_client"
	yamlParser "automato/yaml_parser"
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
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

		for _, action := range at.Actions.Actions {

			automationTx := unpackStringToTransaction(action.Tx.Tx)

		}

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

type TX struct {
	ToAddress common.Address
	Calldata  []byte
}

func unpackStringToTransaction(transaction string) TX {
	//parse the string for the relevant values
	txStringSplit := strings.Split(transaction, "(")
	//get the to address
	toAddress := common.HexToAddress(txStringSplit[0])
	//split the function sig from the rest of the string
	functionSigSplit := strings.Split(txStringSplit[1], ")")
	//init call data
	calldata := []byte{}

	//extract the function sig
	functionSig := keccak256([]byte(functionSigSplit[0]))
	//append function sig to calldata
	calldata = append(calldata, functionSig[:8]...)

	//parse args
	paramsSplit := strings.Split(functionSigSplit[1], ",")

	for i, param := range paramsSplit {

		//if it is the last arg in the arguments, strip the additional ")"
		if i == len(paramsSplit)-1 {
			lastParam := param[:len(param)-1]
			//determine if the param is an int or an address
			if lastParam[:2] == "0x" {
				toAddress := common.HexToAddress(lastParam)
				//add the address to calldata
				calldata = append(calldata, toAddress.Bytes()...)
			} else {
				//if the arg is a number
				uint256ArgBytes, err := rlp.EncodeToBytes(lastParam)
				if err != nil {
					fmt.Printf("Error when converting uint256 arg to bytes, param: {%s}, err:{%v}", param, err)
				}

				calldata = append(calldata, uint256ArgBytes...)
			}
		}

		//determine if the param is an int or an address
		if param[:2] == "0x" {
			toAddress := common.HexToAddress(param)
			//add the address to calldata
			calldata = append(calldata, toAddress.Bytes()...)
		} else {
			//if the arg is a number
			uint256ArgBytes, err := rlp.EncodeToBytes(param)
			if err != nil {
				fmt.Printf("Error when converting uint256 arg to bytes, param: {%s}, err:{%v}", param, err)
			}

			calldata = append(calldata, uint256ArgBytes...)
		}

	}

	return TX{
		ToAddress: toAddress,
		Calldata:  calldata,
	}
}

func keccak256(buf []byte) []byte {
	h := sha3.NewLegacyKeccak256()
	h.Write(buf)
	b := h.Sum(nil)
	return b
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
