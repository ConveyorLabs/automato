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

// type Trigger struct {
// 	//TODO: how to do ors
// 	WhenBlock      int    `When Block Eq @Number Colon`
// 	OnEvent        string `On Event Colon @EventSignature`
// 	EveryXInterval *EveryXInterval
// }

// type Actions struct {
// 	Actions []*Action "@@*"
// }

// type Action struct {
// 	Call *Call
// 	//TODO: how to do "or"
// 	Tx *Tx
// }

// type Call struct {
// 	Address string `Call Colon @Address`
// 	Args    []*Arg
// }

type Arg struct {
	Uint256 int    `@Number`
	Address string `| @Address`
}

// type Tx struct {
// 	Address string `Tx Colon @Address`
// 	Args    []*Arg
// }

type EveryXInterval struct {
	BlockInterval   int `("EVERY" @Number "BLOCKS") | ("EVERY" "BLOCK")`
	SecondsInterval int `| ("EVERY" @Number "SECOND") | ("EVERY" "SECONDS")`

	//Add more interval options
}

//Lexer / Parser
var (
	yamlLexer = lexer.MustSimple([]lexer.SimpleRule{
		{"Comment", `(?:#|//)[^\n]*\n?`},
		{"Identifier", `[a-zA-Z]\w*`},
		{"Number", `(?:\d*\.)?\d+`},
		{"Whitespace", `[ \t\n\r]+`},
		{"Eq", `==`},
		{"Colon", `:`},
		{"Underscore", "_"},
		{"On", `ON`},
		{"Event", `EVENT`},
		{"When", `WHEN`},
		{"Second", `SECOND(S?)`},
		{"Every", `EVERY`},

		//
		{"Address", `"0x" [0-9A-Fa-f]{64}`},

		//
		{"EventSignature", `"0x" [0-9A-Fa-f]{8}`},
		//
		{"Call", `CALL`},
		{"Tx", `TX`},
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
