package service

import (
	"context"
	"encoding/json"
	"net/http"
)

// DecodeAddRequest ...
func DecodeAddRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request AddRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	req = request

	return
}

// DecodeSubtractRequest ...
func DecodeSubtractRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request SubtractRequest

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	req = request

	return
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if err = json.NewEncoder(w).Encode(response); err != nil {
		return
	}

	return
}
