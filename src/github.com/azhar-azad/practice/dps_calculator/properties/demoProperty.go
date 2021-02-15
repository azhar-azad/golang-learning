package properties

type demoProperty struct {
	property
}

func (dp *demoProperty) init() {
	var deposits []float64
	deposits = append(deposits, 500.0)
	deposits = append(deposits, 1000.0)
	deposits = append(deposits, 5000.0)
	deposits = append(deposits, 10000.0)

	var yearRate map[int]float64
	yearRate[3] = 7.00
	yearRate[5] = 7.05
	yearRate[10] = 7.17

	p := property{
		Deposits:      deposits,
		YearRate:      yearRate,
		TableCellSize: 15,
	}

	dp = &demoProperty{
		property: p,
	}
}

func (dp *demoProperty) GetProperty() property {
	dp.init()
	return dp.property
}
