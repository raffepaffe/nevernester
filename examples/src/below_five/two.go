package below_four

import "fmt"

func nestedFuncThatIsNested2Levels() {
	terminator := "Arnold Schwarzenegger"
	rambo := "Sylvester Stallone"
	if terminator != rambo {
		fmt.Println("Terminator and Rambo are not the same")
	}
}
