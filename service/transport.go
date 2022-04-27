package service

import (
	"context"
	"encoding/json"
	"net/http"
)

/* // DecodeAddRequest ...
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
} */

func DecodeRequest(_ context.Context, r *http.Request) (req interface{}, err error) {
	var request Request
	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	return request, nil
}

// EncodeResponse ...
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) (err error) {
	return json.NewEncoder(w).Encode(response)
}
