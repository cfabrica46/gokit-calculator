package service_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/cfabrica46/gokit-calculator/service"
	"github.com/stretchr/testify/assert"
)

const (
	v1Test = "10"
	v2Test = "2"

	addResult      = 12
	subtractResult = 8

	wrongV = "wrong"
)

var errNotTypeIndicated = errors.New("response is not of the type indicated")

func TestAddAndSubtract(t *testing.T) {
	for index, table := range []struct {
		inV1, inV2        string
		outError          string
		outResultAdd      int
		outResultSubtract int
	}{
		{
			inV1:              v1Test,
			inV2:              v2Test,
			outResultAdd:      addResult,
			outResultSubtract: subtractResult,
			outError:          "",
		},
		{
			inV1:              wrongV,
			inV2:              v2Test,
			outResultAdd:      0,
			outResultSubtract: 0,
			outError:          "invalid syntax",
		},
		{
			inV1:              v1Test,
			inV2:              wrongV,
			outResultAdd:      0,
			outResultSubtract: 0,
			outError:          "invalid syntax",
		},
	} {
		t.Run("Add: "+strconv.Itoa(index), func(t *testing.T) {
			var result int
			var resultErr string

			svc := service.NewService()

			result, err := svc.Add(table.inV1, table.inV2)
			if err != nil {
				resultErr = err.Error()
				assert.Contains(t, resultErr, table.outError)
			} else {
				assert.Zero(t, table.outError)
			}

			assert.Equal(t, table.outResultAdd, result)
		})
		t.Run("Subtract: "+strconv.Itoa(index), func(t *testing.T) {
			var result int
			var resultErr string

			svc := service.NewService()

			result, err := svc.Subtract(table.inV1, table.inV2)
			if err != nil {
				resultErr = err.Error()
				assert.Contains(t, resultErr, table.outError)
			} else {
				assert.Zero(t, table.outError)
			}

			assert.Equal(t, table.outResultSubtract, result)
		})
	}
}
