package below_four

import (
	"fmt"
	"unicode/utf8"
)

func nestedFuncThatIsNested4Levels() {
	terminator := "Arnold Schwarzenegger"
	rambo := "Sylvester Stallone"
	if utf8.RuneCountInString(terminator) > 0 {
		if len(terminator) > 0 {
			if terminator != rambo {
				fmt.Println("Terminator and Rambo are not the same")
			}
			if rambo != terminator {
				fmt.Println("Rambo and Terminator are not the same")
			}
		}
	}
}

func nestedFuncThatIsNested4LevelsWithElse() {
	terminator := "Arnold Schwarzenegger"
	rambo := "Sylvester Stallone"
	if utf8.RuneCountInString(terminator) > 0 {
		if len(terminator) > 0 {
			if len(rambo) > 0 {
				fmt.Println("Terminator and Rambo are not the same")
			} else {
				fmt.Println("Rambo is bytes long")
			}
		} else {
			fmt.Println("Terminator is bytes long")
		}
	} else {
		fmt.Println("Terminator is runes long")
	}
}
