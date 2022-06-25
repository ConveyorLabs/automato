package yamlParser

import (
	"fmt"
	"testing"

	"github.com/alecthomas/participle/v2"
)

func TestAddress(t *testing.T) {

	localParser := participle.MustBuild(&Arg{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	ast := &Arg{}

	fileContents := `0x000000000000000000000000000000000000dEaD`

	err := localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

func TestUint256(t *testing.T) {

	localParser := participle.MustBuild(&Arg{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	ast := &Arg{}

	fileContents := "234092340923"

	err := localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

func TestBlockInterval(t *testing.T) {

	localParser := participle.MustBuild(&EveryXInterval{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	ast := &EveryXInterval{}

	fileContents := "EVERY 309324 BLOCKS:"

	//TODO: change to parse and read in yaml file and parse
	err := localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fileContents = "EVERY BLOCK:"

	//TODO: change to parse and read in yaml file and parse
	err = localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

func TestTrigger(t *testing.T) {

	//test when block
	localParser := participle.MustBuild(&Trigger{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	ast := &Trigger{}

	fileContents := "WHEN BLOCK == 234034:"

	err := localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// //test on event

	fileContents = "ON EVENT 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c:"

	err = localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
