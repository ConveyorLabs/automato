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

	//test on event
	ast = &Trigger{}

	fileContents = "ON EVENT 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c:"

	err = localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	//test every x interval

	ast = &Trigger{}

	fileContents = "EVERY BLOCK:"

	err = localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	//test every x blocks
	ast = &Trigger{}

	fileContents = "EVERY 309324 BLOCKS:"

	err = localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	//test every block

	fileContents = "EVERY BLOCK:"
	ast = &Trigger{}

	err = localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

func TestCall(t *testing.T) {

	localParser := participle.MustBuild(&Call{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	ast := &Call{}

	fileContents := "CALL 0x000000000000000000000000000000000000dEaD(functionSig(arg1,arg2), 0x000000000000000000000000000000000000dEaD, 2093490234)"

	err := localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

func TestTx(t *testing.T) {

	localParser := participle.MustBuild(&Tx{},
		participle.Lexer(yamlLexer),
		participle.Elide("Comment", "Whitespace"),
		participle.UseLookahead(2),
	)

	ast := &Tx{}

	fileContents := "TX 0x000000000000000000000000000000000000dEaD(functionSig())"

	err := localParser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

// func TestAction(t *testing.T) {

// 	localParser := participle.MustBuild(&Action{},
// 		participle.Lexer(yamlLexer),
// 		participle.Elide("Comment", "Whitespace"),
// 		participle.UseLookahead(2),
// 	)

// 	ast := &Action{}

// 	fileContents := "CALL 0x000000000000000000000000000000000000dEaDfunctionSig()"

// 	err := localParser.ParseString("fileName", fileContents, ast)
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Fail()
// 	}

// }
