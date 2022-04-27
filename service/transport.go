package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodeRequest ...
func DecodeRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request Request
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, fmt.Errorf("error to decode: %w", err)
	}

	return request, nil
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return fmt.Errorf("error to encode: %w", err)
	}

	return nil
}
