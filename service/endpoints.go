package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeOperationEndpoint ...
func MakeOperationEndpoint(svc serviceInterface, st State) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(Request)

		result, err := svc.Operation(req.V1, req.V2, st)
		if err != nil {
			errMessage = err.Error()
		}

		return Response{Result: result, Err: errMessage}, nil
	}
}
