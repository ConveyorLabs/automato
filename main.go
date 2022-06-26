package main

import (
	"automato/core"
	yamlParser "automato/yaml_parser"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Load the environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file variables", err)
		os.Exit(1)
	}

	// //inititlaize the RPC client
	// rpcClient.Initialize(os.Getenv("HTTP_NODE_URL"), os.Getenv("WS_NODE_URL"))

	//Parse the automation.yaml file
	yamlParser.ParseAutomationYaml()

	core.GenerateAutomationTasks()

	core.StartAutomation()

}
