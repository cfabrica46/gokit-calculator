package service_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cfabrica46/gokit-calculator/service"
	"github.com/stretchr/testify/assert"
)

func TestAddAndSubtractEndpoint(t *testing.T) {
	for index, table := range []struct {
		outErr            string
		in                service.Request
		outResultAdd      int
		outResultSubtract int
		isError           bool
	}{
		{
			in: service.Request{
				V1: v1Test,
				V2: v2Test,
			},
			outResultAdd:      addResult,
			outResultSubtract: subtractResult,
			outErr:            "",
			isError:           false,
		},
		{
			in:                service.Request{},
			outResultAdd:      0,
			outResultSubtract: 0,
			outErr:            "invalid syntax",
			isError:           true,
		},
	} {
		t.Run("Add: "+strconv.Itoa(index), func(t *testing.T) {
			svc := service.NewService()

			r, err := service.MakeAddEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.Response)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outResultAdd, result.Result)
		})
		t.Run("Subtract: "+strconv.Itoa(index), func(t *testing.T) {
			svc := service.NewService()

			r, err := service.MakeSubtractEndpoint(svc)(context.TODO(), table.in)
			if err != nil {
				t.Error(err)
			}

			result, ok := r.(service.Response)
			if !ok {
				t.Error(errNotTypeIndicated)
			}

			if !table.isError {
				assert.Zero(t, result.Err)
			} else {
				assert.Contains(t, result.Err, table.outErr)
			}

			assert.Equal(t, table.outResultSubtract, result.Result)
		})
	}
}
