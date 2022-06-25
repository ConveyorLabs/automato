package main

import (
	rpcClient "automato/rpc_client"
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

	//inititlaize the RPC client
	rpcClient.Initialize(os.Getenv("HTTP_NODE_URL"), os.Getenv("WS_NODE_URL"))

	//Parse the automation.yaml file

}
