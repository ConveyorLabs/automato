package yamlParser

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func ParseAutomationYaml() *YamlFile {

	ast := &YamlFile{}

	b, err := ioutil.ReadFile("automation.yaml")
	if err != nil {
		panic(err)
	}

	fileAsString := string(b)
	err = parser.ParseString("filename", fileAsString, ast)
	if err != nil {
		fmt.Println("Issue when parsing the automation.yaml file")
		fmt.Println(err)
		os.Exit(1)
	}

	return ast

}

type YamlFile struct {
	AutomationTasks []*AutomationTask "@@*"
}

type AutomationTask struct {
	Trigger *Trigger "@@"
	Actions *Actions `@@`
}

type Actions struct {
	Actions []*Action "@@*"
}

type Trigger struct {
	//TODO: how to do ors
	WhenBlock       int    `"WHEN" "BLOCK" Eq @Number Colon`
	OnEvent         string `| "ON" "EVENT" @EventSignature Colon`
	BlockInterval   int    `| (("EVERY" @Number "BLOCKS") | ("EVERY" "BLOCK")) Colon`
	SecondsInterval int    `| (("EVERY" @Number "SECOND") | ("EVERY" "SECONDS")) Colon`
}

type Action struct {
	Tx *Tx `@@`
	//TODO: set it up so that it can call and get the
	//return data to pass into a tx
	// Call *Call

}

type Arg struct {
	Uint256 int    `@Number`
	Address string `| @Address`
}

// type Call struct {
// 	Call string `"CALL" Colon  @FunctionCall`
// }

type Tx struct {
	Tx string `"TX" Colon @FunctionCall`
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
	})

	parser = participle.MustBuild(&YamlFile{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)
