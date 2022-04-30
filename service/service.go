package service

import (
	"fmt"
	"strconv"
)

type serviceInterface interface {
	Add(string, string) (int, error)
	Subtract(string, string) (int, error)
}

// Service ...
type Service struct{}

// NewService ...
func NewService() *Service {
	return &Service{}
}

// Add ...
func (Service) Add(v1, v2 string) (int, error) {
	vInt1, err := strconv.Atoi(v1)
	if err != nil {
		return 0, fmt.Errorf("error convertation: %w", err)
	}

	vInt2, err := strconv.Atoi(v2)
	if err != nil {
		return 0, fmt.Errorf("error convertation: %w", err)
	}

	return vInt1 + vInt2, nil
}

// Subtract ...
func (Service) Subtract(v1, v2 string) (int, error) {
	vInt1, err := strconv.Atoi(v1)
	if err != nil {
		return 0, fmt.Errorf("error convertation: %w", err)
	}

	vInt2, err := strconv.Atoi(v2)
	if err != nil {
		return 0, fmt.Errorf("error convertation: %w", err)
	}

	return vInt1 - vInt2, nil
}
