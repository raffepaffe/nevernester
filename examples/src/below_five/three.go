package below_four

import "fmt"

func nestedFuncThatIsNested3Levels() {
	terminator := "Arnold Schwarzenegger"
	rambo := "Sylvester Stallone"
	if terminator != rambo {
		if len(terminator) > 0 {
			fmt.Println("Terminator has bytes")
		}
	}
}

func rowBreak() {
	type row struct {
		row *row
	}

	r1 := row{}
	r2 := row{}
	r3 := row{}
	r4 := row{}

	r1.row = &r2
	r2.row = &r3
	r3.row = &r4

	r5 := r1.
		row.
		row.
		row
	if &r5 != nil {
		r6 := r5.
			row.
			row.
			row.
			row
		fmt.Printf("%v is not nil", r6)
	}
}
