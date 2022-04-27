package service_test

import (
	"context"
	"strconv"
	"testing"

	"github.com/cfabrica46/gokit-calculator/service"
	"github.com/stretchr/testify/assert"
)

func TestOperationEndpoint(t *testing.T) {
	for index, table := range []struct {
		inState   service.State
		outErr    string
		in        service.Request
		outResult int
		isError   bool
	}{
		{
			in: service.Request{
				V1: v1Test,
				V2: v2Test,
			},
			inState:   service.NewAddState(),
			outResult: addResult,
			outErr:    "",
			isError:   false,
		},
		{
			in: service.Request{
				V1: v1Test,
				V2: v2Test,
			},
			inState:   service.NewSubtractState(),
			outResult: subtractResult,
			outErr:    "",
			isError:   false,
		},
		{
			in:        service.Request{},
			outResult: 0,
			outErr:    "invalid syntax",
			isError:   true,
		},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			svc := service.NewService()

			r, err := service.MakeOperationEndpoint(svc, table.inState)(context.TODO(), table.in)
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

			assert.Equal(t, table.outResult, result.Result)
		})
	}
}
