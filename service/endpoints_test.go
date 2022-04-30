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
		inAdd             service.AddRequest
		inSubtract        service.SubtractRequest
		outResultAdd      int
		outResultSubtract int
		isError           bool
	}{
		{
			inAdd: service.AddRequest{
				V1: v1Test,
				V2: v2Test,
			},
			inSubtract: service.SubtractRequest{
				V1: v1Test,
				V2: v2Test,
			},
			outResultAdd:      addResult,
			outResultSubtract: subtractResult,
			outErr:            "",
			isError:           false,
		},
		{
			inAdd:             service.AddRequest{},
			inSubtract:        service.SubtractRequest{},
			outResultAdd:      0,
			outResultSubtract: 0,
			outErr:            "invalid syntax",
			isError:           true,
		},
	} {
		t.Run("Add: "+strconv.Itoa(index), func(t *testing.T) {
			svc := service.NewService()

			r, err := service.MakeAddEndpoint(svc)(context.TODO(), table.inAdd)
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

			assert.Equal(t, table.outResultAdd, result.Result)
		})
		t.Run("Subtract: "+strconv.Itoa(index), func(t *testing.T) {
			svc := service.NewService()

			r, err := service.MakeSubtractEndpoint(svc)(context.TODO(), table.inSubtract)
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

			assert.Equal(t, table.outResultSubtract, result.Result)
		})
	}
}
