// package main

// import (
// 	rpcClient "automato/rpc_client"
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func main() {
// 	//Load the environment variables
// 	err := godotenv.Load()
// 	if err != nil {
// 		fmt.Println("Error loading .env file variables", err)
// 		os.Exit(1)
// 	}

// 	//inititlaize the RPC client
// 	rpcClient.Initialize(os.Getenv("HTTP_NODE_URL"), os.Getenv("WS_NODE_URL"))

// 	//Parse the automation.yaml file

// }

package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/repr"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

type YamlFile struct {
	// WhenBlock *WhenBlock
	Automation []*Automation "@@*"
}

type Automation struct {
	Trigger *Trigger "@@*"
	Actions *Actions "@@*"
}

type Trigger struct {
	WhenBlock      int    `WhenBlock (Eq) @Number Colon`
	OnEvent        string `OnEvent Colon @EventSignature`
	EveryXInterval *EveryXInterval
}

type Actions struct {
	Actions []*Action "@@*"
}

type Action struct {
	Call *Call
	//TODO: how to do "or"
	Tx *Tx
}

type Call struct {
	Address string `Call Colon @Address`
	Args    []*Arg
}

type Arg struct {
	Uint256 int    `@Number`
	Address string `| @Address`
}

type Tx struct {
	Address string `Tx Colon @Address`
	Args    []*Arg
}

type EveryXInterval struct {
	BlockInterval   int `Every @Number? Block`
	SecondsInterval int `Every @Number Second`
	//Add more interval options
}

var (
	yamlLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Comment", `(?:#|//)[^\n]*\n?`},
		{"Identifier", `[a-zA-Z]\w*`},
		{"Number", `(?:\d*\.)?\d+`},
		{"Whitespace", `[ \t\n\r]+`},
		{"Eq", `==`},
		{"Colon", `:`},

		//
		{"Address", `"0x" [0-9A-Fa-f]{64}`},

		//
		{"EventSignature", `"0x" [0-9A-Fa-f]{8}`},

		//Triggers
		{"WhenBlock", `WHEN_BLOCK`},
		{"OnEvent", `ON_EVENT`},
		{"Every", `EVERY`},

		//Intervals
		{"Block", `BLOCK(S?)`},
		{"Second", `SECOND(S?)`},

		//Actions
		{"Call", `CALL`},
		{"Tx", `TX`},
	})

	parser = participle.MustBuild(&YamlFile{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)

func main() {

	ast := &YamlFile{}

	fileContents := `when block 123093`

	//TODO: change to parse and read in yaml file and parse
	err := parser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repr.Println(ast)

}
