package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type FuncMakeEndpoint func(serviceInterface) endpoint.Endpoint

// MakeAddEndpoint ...
func MakeAddEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, _ := request.(Request)

		result, err := svc.Add(req.V1, req.V2)
		if err != nil {
			errMessage = err.Error()
		}

		return Response{Result: result, Err: errMessage}, nil
	}
}

// MakeSubtactEndpoint ...
func MakeSubtractEndpoint(svc serviceInterface) endpoint.Endpoint {
	return func(_ context.Context, request any) (any, error) {
		var errMessage string

		req, _ := request.(Request)

		result, err := svc.Subtract(req.V1, req.V2)
		if err != nil {
			errMessage = err.Error()
		}

		return Response{Result: result, Err: errMessage}, nil
	}
}
