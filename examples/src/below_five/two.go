package below_four

import (
	"encoding/json"
	"fmt"
)

func nestedFuncThatIsNested2Levels() {
	terminator := "Arnold Schwarzenegger"
	rambo := "Sylvester Stallone"
	if terminator != rambo {
		fmt.Println("Terminator and Rambo are not the same")
	}
}

func ifNesting2Levels() {
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(
		byt, &dat); err != nil {
		panic(err)
	}
}
