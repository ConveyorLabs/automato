package yamlParser

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
	"github.com/alecthomas/repr"
)

func ParseAutomationYaml() {

	ast := &Arg{}

	// fileContents, err := os.ReadFile("../automation.yaml")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	f, err := ioutil.ReadFile("automation.yaml")
	if err != nil {
		panic(err)
	}

	stringF := string(f)

	//TODO: change to parse and read in yaml file and parse
	err = parser.ParseString("filename", stringF, ast)
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

type Action struct {
	Call *Call
	//TODO: how to do "or"
	Tx *Tx `|`
}

type Call struct {
	Address string `Call Colon @Address`
	Args    []*Arg
}

type Arg struct {
	Uint256 int    `@Number`
	Address string `| @Address`
}

type Trigger struct {
	//TODO: how to do ors
	// WhenBlock int    `"WHEN" "BLOCK" Eq @Number Colon`
	OnEvent string `@Address`
	// EveryXInterval *EveryXInterval
}

type Tx struct {
	Address string `"TX" "TO" @Address Colon`
	Args    []*Arg
}

type EveryXInterval struct {
	BlockInterval   int `(("EVERY" @Number "BLOCKS") | ("EVERY" "BLOCK")) Colon`
	SecondsInterval int `| (("EVERY" @Number "SECOND") | ("EVERY" "SECONDS")) Colon`

	//Add more interval options
}

//Lexer / Parser
var (
	yamlLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"EventSignature", `^(0x|0X)?[a-fA-F0-9]+`},
		{"Comment", `(?:#|//)[^\n]*\n?`},

		{"Address", `0[xX][0-9a-fA-F]{40}`},

		{"Identifier", `[a-zA-Z]\w*`},
		{"Number", `(?:\d*\.)?\d+`},
		{"Whitespace", `[ \t\n\r]+`},
		{"Eq", `==`},
		{"Colon", `:`},

		{"Underscore", "_"},

		//

		//
		//Triggers
		// {"WhenBlock", `WHEN Underscore Block`},
		// {"OnEvent", `On Underscore Event`},
		// {"EveryXBlock", `Every Underscore Number Underscore Block`},

		//TODO:
		{"Indent", `four spaces or a tab`},
	})

	parser = participle.MustBuild(&Arg{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)
