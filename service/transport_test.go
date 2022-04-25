package service_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/cfabrica46/gokit-calculator/service"
	"github.com/stretchr/testify/assert"
)

func TestDecodeAddRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSON, err := json.Marshal(service.AddRequest{V1: v1Test, V2: v2Test})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		in     *http.Request
		out    service.AddRequest
		outErr string
	}{
		{goodReq, service.AddRequest{V1: v1Test, V2: v2Test}, ""},
		{badReq, service.AddRequest{}, "EOF"},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeAddRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.AddRequest)
			if !ok {
				if (table.out != service.AddRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestDecodeSubtractRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSON, err := json.Marshal(service.SubtractRequest{V1: v1Test, V2: v2Test})
	if err != nil {
		t.Error(err)
	}

	goodReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(dataJSON))
	if err != nil {
		t.Error(err)
	}

	badReq, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte{}))
	if err != nil {
		t.Error(err)
	}

	for index, table := range []struct {
		in     *http.Request
		out    service.SubtractRequest
		outErr string
	}{
		{goodReq, service.SubtractRequest{V1: v1Test, V2: v2Test}, ""},
		{badReq, service.SubtractRequest{}, "EOF"},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeSubtractRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.SubtractRequest)
			if !ok {
				if (table.out != service.SubtractRequest{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Equal(t, table.outErr, resultErr)
			assert.Equal(t, table.out, result)
		})
	}
}

func TestEncodeResponse(t *testing.T) {
	for index, table := range []struct {
		in     interface{}
		outErr string
	}{
		{"test", ""},
		{func() {}, "json: unsupported type: func()"},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			var resultErr string

			err := service.EncodeResponse(context.TODO(), httptest.NewRecorder(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			assert.Equal(t, table.outErr, resultErr)
		})
	}
}
