package properties

type IProperty interface {
	GetDeposits() []float64
	GetYearRate() map[int]float64
	GetTableCellSize() int
	init()
}
