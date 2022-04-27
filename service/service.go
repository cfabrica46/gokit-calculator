package service

import (
	"strconv"
)

type serviceInterface interface {
	/* Add(string, string) (int, error)
	Subtract(string, string) (int, error) */
	Operation(string, string) (int, error)
}

// Service ...
type Service struct{}

// NewService ...
func NewService() *Service {
	return &Service{}
}

// Operation...
func (Service) Operation(v1, v2 string) (result int, err error) {
	vInt1, err := strconv.Atoi(v1)
	if err != nil {
		return
	}

	vInt2, err := strconv.Atoi(v2)
	if err != nil {
		return
	}

	result = vInt1 + vInt2

	return result, err
}

/* // Add ...
func (Service) Add(v1, v2 string) (result int, err error) {
	vInt1, err := strconv.Atoi(v1)
	if err != nil {
		return
	}

	vInt2, err := strconv.Atoi(v2)
	if err != nil {
		return
	}

	result = vInt1 + vInt2

	return
}

// Subtract ...
func (Service) Subtract(v1, v2 string) (result int, err error) {
	fmt.Println(v1, v2)

	vInt1, err := strconv.Atoi(v1)
	if err != nil {
		return
	}

	vInt2, err := strconv.Atoi(v2)
	if err != nil {
		return
	}

	result = vInt1 - vInt2

	return
} */
