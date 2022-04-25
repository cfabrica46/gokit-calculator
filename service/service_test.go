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

// errWebServer        = errors.New("error from web server")
var errNotTypeIndicated = errors.New("response is not of the type indicated")

func TestAdd(t *testing.T) {
	for index, table := range []struct {
		inV1, inV2 string
		outError   string
		outResult  int
	}{
		{
			inV1:      v1Test,
			inV2:      v2Test,
			outResult: addResult,
			outError:  "",
		},
		{
			inV1:      wrongV,
			inV2:      v2Test,
			outResult: 0,
			outError:  "invalid syntax",
		},
		{
			inV1:      v1Test,
			inV2:      wrongV,
			outResult: 0,
			outError:  "invalid syntax",
		},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
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

			assert.Equal(t, table.outResult, result)
		})
	}
}

func TestSubtract(t *testing.T) {
	for index, table := range []struct {
		inV1, inV2 string
		outError   string
		outResult  int
	}{
		{
			inV1:      v1Test,
			inV2:      v2Test,
			outResult: subtractResult,
			outError:  "",
		},
		{
			inV1:      wrongV,
			inV2:      v2Test,
			outResult: 0,
			outError:  "invalid syntax",
		},
		{
			inV1:      v1Test,
			inV2:      wrongV,
			outResult: 0,
			outError:  "invalid syntax",
		},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
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

			assert.Equal(t, table.outResult, result)
		})
	}
}
