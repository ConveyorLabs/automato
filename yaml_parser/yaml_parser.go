package yamlParser

import (
	"fmt"
	"os"

	"github.com/alecthomas/repr"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func ParseAutomationYaml() {

	ast := &Arg{}

	fileContents := "234092340923"

	//TODO: change to parse and read in yaml file and parse
	err := parser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	repr.Println(ast)

}

// type YamlFile struct {
// 	// WhenBlock *WhenBlock
// 	AutomationTasks []*AutomationTask "@@*"
// }

// type AutomationTask struct {
// 	Trigger *Trigger "@@*"
// 	Actions *Actions "@@*"
// }

// type Actions struct {
// 	Actions []*Action "@@*"
// }

type Trigger struct {
	//TODO: how to do ors
	WhenBlock       int    `"WHEN" "BLOCK" Eq @Number Colon`
	OnEvent         string `| "ON" "EVENT" @EventSignature Colon`
	BlockInterval   int    `| (("EVERY" @Number "BLOCKS") | ("EVERY" "BLOCK")) Colon`
	SecondsInterval int    `| (("EVERY" @Number "SECOND") | ("EVERY" "SECONDS")) Colon`
}

type Action struct {
	Call *Call
	//TODO: how to do "or"
	// Tx *Tx `|`
}

type Arg struct {
	Uint256 int    `@Number`
	Address string `| @Address`
}

type Call struct {
	Call string `"CALL" @FunctionCall`
}

type Tx struct {
	Tx string `"TX" @FunctionCall`
}

//Lexer / Parser
var (
	yamlLexer = lexer.MustSimple([]lexer.SimpleRule{
		//If you got here, please dont look at this regex
		{"FunctionCall", `0[xX][0-9a-fA-F]{40}\([a-zA-Z]+\([a-zA-Z0-9]*(,[a-zA-Z0-9]*)*\)(\s*,\s*[a-zA-Z0-9]+\s*)*\)`},
		{"Comment", `(?:#|//)[^\n]*\n?`},
		{"EventSignature", `0[xX][0-9a-fA-F]{64}`},
		{"Address", `0[xX][0-9a-fA-F]{40}`},
		{"Identifier", `[a-zA-Z]\w*`},
		{"Number", `(?:\d*\.)?\d+`},
		{"Whitespace", `[ \t\n\r]+`},
		{"Eq", `==`},
		{"Colon", `:`},
		{"Underscore", "_"},

		//TODO:
		{"Indent", `four spaces or a tab`},
	})

	parser = participle.MustBuild(&Arg{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)
