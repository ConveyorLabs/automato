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
	WhenBlock *WhenBlock
	// Symbols   []*Symbol "@@*"
}

type WhenBlock struct {
	BlockNumber int `"when block" @@Number`
}

type Symbol struct {
	String string ` @Ident`
	Number int    `| @Number`
}

var (
	yamlLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Comment", `(?:#|//)[^\n]*\n?`},
		{"Ident", `[a-zA-Z]\w*`},
		{"Number", `(?:\d*\.)?\d+`},
		{"Whitespace", `[ \t\n\r]+`},
	})

	parser = participle.MustBuild(&YamlFile{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)

func main() {

	ast := &YamlFile{}

	fileContents := `
	//some comment
	//thing

	//230493409
	when block 1230934

	//another thing
	
	`

	err := parser.ParseString("name", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repr.Println(ast)

}
