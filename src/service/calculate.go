package service

type Calculate interface {
	Service(loopNumber int) int
}

func NewCalculate() Calculate {
	return &calculate{}
}

type calculate struct{}

func (c *calculate) Service(loopNumber int) int {
	result := 0
	for i := 1; i <= loopNumber; i++ {
		result += i
	}
	return result
}
