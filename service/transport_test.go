package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/cfabrica46/gokit-calculator/service"
	"github.com/stretchr/testify/assert"
)

func TestDecodeAddAndSubtractRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSONAdd, err := json.Marshal(service.AddRequest{V1: v1Test, V2: v2Test})
	if err != nil {
		t.Error(err)
	}

	goodReqAdd, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSONAdd))
	if err != nil {
		t.Error(err)
	}

	badReqAdd, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	dataJSONSubtract, err := json.Marshal(service.SubtractRequest{V1: v1Test, V2: v2Test})
	if err != nil {
		t.Error(err)
	}

	goodReqSubtract, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSONSubtract))
	if err != nil {
		t.Error(err)
	}

	badReqSubtract, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		inAdd       *http.Request
		inSubtract  *http.Request
		outAdd      service.AddRequest
		outSubtract service.SubtractRequest
		outErr      string
	}{
		{
			inAdd:       goodReqAdd,
			inSubtract:  goodReqSubtract,
			outAdd:      service.AddRequest{V1: v1Test, V2: v2Test},
			outSubtract: service.SubtractRequest{V1: v1Test, V2: v2Test},
			outErr:      "",
		},
		{
			inAdd:       badReqAdd,
			inSubtract:  badReqSubtract,
			outAdd:      service.AddRequest{},
			outSubtract: service.SubtractRequest{},
			outErr:      "EOF",
		},
	} {
		t.Run("Add: "+strconv.Itoa(index), func(t *testing.T) {
			var result any
			var resultErr string

			r, err := service.DecodeAddRequest(context.TODO(), table.inAdd)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.AddRequest)
			if !ok {
				if (table.outAdd != service.AddRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Contains(t, resultErr, table.outErr)
			assert.Equal(t, table.outAdd, result)
		})
		t.Run("Subtract: "+strconv.Itoa(index), func(t *testing.T) {
			var result any
			var resultErr string

			r, err := service.DecodeSubtractRequest(context.TODO(), table.inSubtract)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.SubtractRequest)
			if !ok {
				if (table.outSubtract != service.SubtractRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Contains(t, resultErr, table.outErr)
			assert.Equal(t, table.outAdd, result)
		})
	}
}
