package service

type Calculate struct{}

func (c *Calculate) Service(loopNumber int) int {
	result := 0
	for i := 1; i <= loopNumber; i++ {
		result += i
	}
	return result
}
