package main

import (
	"automato/core"
	yamlParser "automato/yaml_parser"
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	//create a new wait group that waits unitl all tasks are finished
	wg := &sync.WaitGroup{}
	wg.Add(1)

	//Load the environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file variables", err)
		os.Exit(1)
	}

	// //inititlaize the RPC client
	// rpcClient.Initialize(os.Getenv("HTTP_NODE_URL"), os.Getenv("WS_NODE_URL"))

	//Parse the automation.yaml file
	ast := yamlParser.ParseAutomationYaml()

	automationTasks := core.GenerateAutomationTasks(ast)

	//start listening to block headers and automate tasks
	go core.StartAutomation(automationTasks)

	wg.Wait()

}
