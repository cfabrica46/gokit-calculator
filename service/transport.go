package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DecodeRequest ...
func DecodeRequest(_ context.Context, r *http.Request) (any, error) {
	var request Request

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
