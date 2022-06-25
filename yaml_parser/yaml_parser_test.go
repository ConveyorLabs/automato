package yamlParser

import (
	"fmt"
	"testing"
)

func TestAddressArg(t *testing.T) {
	ast := &Arg{}

	fileContents := "234092340923"

	//TODO: change to parse and read in yaml file and parse
	err := parser.ParseString("fileName", fileContents, ast)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

// func TestUint256Arg(t *testing.T) {

// 	ast := &Arg{}

// 	fileContents := "2348923498230492349823498"

// 	err := parser.ParseString("fileName", fileContents, ast)
// 	if err != nil {
// 		t.Fail()
// 	}

// }
