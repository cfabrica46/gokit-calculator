package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// MakeAddEndpoint ...
func MakeAddEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(Request)

		result, err := svc.Operation(req.V1, req.V2)
		if err != nil {
			errMessage = err.Error()
		}

		return Response{Result: result, Err: errMessage}, nil
	}
}

// MakeSubtactEndpoint ...
func MakeSubtractEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(Request)

		result, err := svc.Operation(req.V1, req.V2)
		if err != nil {
			errMessage = err.Error()
		}

		return Response{Result: result, Err: errMessage}, nil
	}
}

/* // MakeOperationEndpoint ...
func MakeOperationEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		var errMessage string

		req, _ := request.(Request)

		result, err := svc.Operation(req.V1, req.V2)
		if err != nil {
			errMessage = err.Error()
		}

		return Response{Result: result, Err: errMessage}, nil
	}
} */
