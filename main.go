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
	Triggers []*Trigger "@@*"
}

type Trigger struct {
	WhenBlock      int    `WhenBlock (Eq) @Number`
	OnEvent        string `OnEvent @EventSignature`
	EveryXInterval *EveryXInterval

	// String      string ` | @Ident`
	// Number int `| @Number`
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

		//
		{"EventSignature", `"0x" [0-9A-Fa-f]{8}`},

		//Triggers
		{"WhenBlock", `when block`},
		{"OnEvent", `on event`},
		{"Every", `every`},

		//Intervals
		{"Block", `block(s?)`},
		{"Second", `second(s?)`},
	})

	parser = participle.MustBuild(&YamlFile{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)

func main() {

	ast := &YamlFile{}

	fileContents := `when block 123093
	
	`

	err := parser.ParseString("name", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repr.Println(ast)

}
