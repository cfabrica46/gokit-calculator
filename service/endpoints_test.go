package service_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cfabrica46/gokit-calculator/service"
	"github.com/stretchr/testify/assert"
)

func TestAddEndpoint(t *testing.T) {
	for index, table := range []struct {
		outErr    string
		in        service.AddRequest
		outResult int
		isError   bool
	}{
		{
			in: service.AddRequest{
				V1: v1Test,
				V2: v2Test,
			},
			outResult: addResult,
			outErr:    "",
			isError:   false,
		},
		{
			in:        service.AddRequest{},
			outResult: 0,
			outErr:    "invalid syntax",
			isError:   true,
		},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			svc := service.NewService()

			r, err := service.MakeAddEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.AddResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outResult, result.Result)
		})
	}
}

func TestSubtractEndpoint(t *testing.T) {
	for index, table := range []struct {
		outErr    string
		in        service.SubtractRequest
		outResult int
		isError   bool
	}{
		{
			in: service.SubtractRequest{
				V1: v1Test,
				V2: v2Test,
			},
			outResult: subtractResult,
			outErr:    "",
			isError:   false,
		},
		{
			in:        service.SubtractRequest{},
			outResult: 0,
			outErr:    "invalid syntax",
			isError:   true,
		},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			svc := service.NewService()

			r, err := service.MakeSubtractEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.SubtractResponse)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outResult, result.Result)
		})
	}
}
