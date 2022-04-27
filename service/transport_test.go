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

func TestDecodeRequest(t *testing.T) {
	url := "localhost:8080"

	dataJSON, err := json.Marshal(service.Request{V1: v1Test, V2: v2Test})
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
		out    service.Request
		outErr string
	}{
		{goodReq, service.Request{V1: v1Test, V2: v2Test}, ""},
		{badReq, service.Request{}, "EOF"},
	} {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			var result interface{}
			var resultErr string

			r, err := service.DecodeRequest(context.TODO(), table.in)
			if err != nil {
				resultErr = err.Error()
			}

			result, ok := r.(service.Request)
			if !ok {
				if (table.out != service.Request{}) {
					t.Error("result is not of the type indicated")
				}
			}

			assert.Contains(t, resultErr, table.outErr)
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

			assert.Contains(t, resultErr, table.outErr)
		})
	}
}
