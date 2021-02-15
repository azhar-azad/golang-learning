package properties

type property struct {
	Deposits      []float64
	YearRate      map[int]float64
	TableCellSize int
}

func (p *property) GetDeposits() []float64 {
	return p.Deposits
}

func (p *property) GetYearRate() map[int]float64 {
	return p.YearRate
}

func (p *property) GetTableCellSize() int {
	return p.TableCellSize
}
