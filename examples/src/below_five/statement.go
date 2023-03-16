package below_four

type row struct {
	property []string
}

func (r row) withProperty(s string) row {
	r.property = append(r.property, s)
	return r
}

func assignStatement() {
	r := row{}
	if len(r.property) > 0 {
		if len(r.property) > 0 {
			if len(r.property) > 0 {
				r.property = []string{
					"one",
					"two",
				}
			}
		}
	}
}

func exprStatement() {
	r := row{}
	r.withProperty("one")
	if len(r.property) > 0 {
		if len(r.property) > 0 {
			if len(r.property) > 0 {
				r.withProperty("two").
					withProperty("three").
					withProperty("four").
					withProperty("five").
					withProperty("six")
			}
		}
	}
}
