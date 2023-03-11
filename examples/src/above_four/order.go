package above_four

type Order struct {
	Header OrderHeader
	Rows   []OrderRow
}

type OrderHeader struct {
	IsValid bool
}

type OrderRow struct {
	IsValid bool
	Price   int
}

func calculate(order *Order) int { // want "calculated nesting for function calculate is 5, max is 4"
	sum := 0
	if order != nil {
		if order.Header.IsValid {
			for _, row := range order.Rows {
				if row.IsValid {
					sum = sum + row.Price
				}
			}
		}
	}
	return sum
}

func calculate2(order *Order) int {
	sum := 0
	if order == nil {
		return sum
	}

	if order.Header.IsValid {
		sum = getPrice(order.Rows)
	}

	return sum
}

func getPrice(rows []OrderRow) int {
	sum := 0
	for _, row := range rows {
		if row.IsValid {
			sum = sum + row.Price
		}
	}
	return sum
}
