package above_four

import (
	"fmt"
	"unicode/utf8"
)

func nestedFuncThatIsNested5Levels() { // want "calculated nesting for function nestedFuncThatIsNested5Levels is 5, max is 4"
	terminator := "Arnold Schwarzenegger"
	rambo := "Sylvester Stallone"
	if utf8.RuneCountInString(terminator) > 0 {
		if len(terminator) > 0 {
			if len(rambo) > 0 {
				if terminator != rambo {
					fmt.Println("Terminator and Rambo are not the same")
				}
			}
		}
	}
}

func nestedFuncThatIsNested5LevelsWithElse() { // want "calculated nesting for function nestedFuncThatIsNested5LevelsWithElse is 5, max is 4"
	terminator := "Arnold Schwarzenegger"
	rambo := "Sylvester Stallone"
	if utf8.RuneCountInString(terminator) > 0 {
		if len(terminator) > 0 {
			if len(rambo) > 0 {
				fmt.Println("Terminator and Rambo are not the same")
			} else {
				if utf8.RuneCountInString(rambo) > 0 {
					fmt.Println("Rambo is runes long")
				}
				fmt.Println("Rambo is bytes long")
			}
		} else {
			fmt.Println("Terminator is bytes long")
		}
	} else {
		fmt.Println("Terminator is runes long")
	}
}
