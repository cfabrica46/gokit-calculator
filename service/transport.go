package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodeAddRequest ...
func DecodeAddRequest(_ context.Context, r *http.Request) (any, error) {
	var request AddRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("error to decode: %w", err)
	}

	return request, nil
}

// DecodeSubtractRequest ...
func DecodeSubtractRequest(_ context.Context, r *http.Request) (any, error) {
	var request SubtractRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("error to decode: %w", err)
	}

	return request, nil
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response any) (err error) {
	if err = json.NewEncoder(w).Encode(response); err != nil {
		return
	}

	return
}
