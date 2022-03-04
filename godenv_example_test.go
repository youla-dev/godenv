package godenv_test

import (
	"bytes"
	"fmt"
	"os"

	"github.com/youla-dev/godenv"
)

func ExampleParse() {
	f, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	vars, err := godenv.Parse(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(vars)

	envContent := `VARIABLE_1='This is variable 1'
# This is comment
VARIABLE_2="This is variable 2"
VARIABLE_TAB_1='Tab is not escaped\t'
VARIABLE_TAB_2="Tab is escaped\t"`

	buf := bytes.NewBufferString(envContent)

	result, err := godenv.Parse(buf)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	// Output:
	// map[VARIABLE_1:This is variable 1 VARIABLE_2:This is variable 2 VARIABLE_TAB_1:Tab is not escaped\t VARIABLE_TAB_2:Tab is escaped	]
}
