package yamlParser

import (
	"fmt"
	"os"

	"github.com/alecthomas/repr"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

func ParseYamlFile() {

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

type YamlFile struct {
	// WhenBlock *WhenBlock
	Automation []*Automation "@@*"
}

type Automation struct {
	Trigger *Trigger "@@*"
	Actions *Actions "@@*"
}

type Trigger struct {
	//TODO: how to do ors
	WhenBlock      int    `When Block Eq @Number Colon`
	OnEvent        string `On Event Colon @EventSignature`
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
	BlockInterval   int `Every Underscore @Number? Underscore Block`
	SecondsInterval int `Every Underscore @Number Underscore Second`
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
		{"Block", `(BLOCK)(S)?`},
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

	parser = participle.MustBuild(&YamlFile{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)
)
