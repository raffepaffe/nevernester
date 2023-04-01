package above_four

import (
	"fmt"
	"net/http"
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

func nestedFuncWithReturnNesting6Levels() http.HandlerFunc { // want "calculated nesting for function nestedFuncWithReturnNesting6Levels is 6, max is 4"
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if ctx == nil {
				w.WriteHeader(500)
				if r == nil {
					w.WriteHeader(404)
				} else {
					if w == nil {
						w.WriteHeader(402)
					}
					w.WriteHeader(401)
				}
				return
			}

		})
}
