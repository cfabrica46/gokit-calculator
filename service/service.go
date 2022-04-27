package service

import (
	"fmt"
	"strconv"
)

type serviceInterface interface {
	/* Add(string, string) (int, error)
	Subtract(string, string) (int, error) */
	Operation(string, string, State) (int, error)
}

// Service ...
type Service struct{}

// NewService ...
func NewService() *Service {
	return &Service{}
}

// Operation...
func (Service) Operation(v1, v2 string, st State) (result int, err error) {
	vInt1, err := strconv.Atoi(v1)
	if err != nil {
		return result, fmt.Errorf("error convertation: %w", err)
	}

	vInt2, err := strconv.Atoi(v2)
	if err != nil {
		return result, fmt.Errorf("error convertation: %w", err)
	}

	return st.Operation(vInt1, vInt2), err
}
