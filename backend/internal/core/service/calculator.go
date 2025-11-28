package service

import "errors"

// let's have a calculator service that will define the interface for our business logic

type CalculatorService interface {
	Add(a, b int32) int32
	Sub(a, b int32) int32
	Mul(a, b int32) int32
	Div(a, b int32) (int32, error) // error if the result isn't a 32 bits integer
	Mod(a, b int32) (int32, error)
}

type CalculatorServiceImpl struct {
}

func NewCalculatorService() CalculatorService {
	return &CalculatorServiceImpl{}
}

func (c *CalculatorServiceImpl) Add(a, b int32) int32 {
	return a + b
}
func (c *CalculatorServiceImpl) Sub(a, b int32) int32 {
	return a - b
}
func (c *CalculatorServiceImpl) Mul(a, b int32) int32 {
	return a * b
}
func (c *CalculatorServiceImpl) Div(a, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	} else {
		return a / b, nil
	}
}
func (c *CalculatorServiceImpl) Mod(a, b int32) (int32, error) {
	if b == 0 {
		return 0, errors.New("cannot mod zero")
	} else {
		return a % b, nil
	}
}
