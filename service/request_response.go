package service

/* // AddRequest (string, string) (int, error).
type AddRequest struct {
	V1 string `json:"v1"`
	V2 string `json:"v2"`
}

// AddResponse (string, string) (int, error).
type AddResponse struct {
	Err    string `json:"err,omitempty"`
	Result int    `json:"result"`
} */

/* // SubtractRequest (string, string) (int, error).
type SubtractRequest struct {
	V1 string `json:"v1"`
	V2 string `json:"v2"`
}

// SubtractResponse (string, string) (int, error).
type SubtractResponse struct {
	Err    string `json:"err,omitempty"`
	Result int    `json:"result"`
} */

// Request ...
type Request struct {
	V1 string `json:"v1"`
	V2 string `json:"v2"`
}

// Response ...
type Response struct {
	Err    string `json:"err,omitempty"`
	Result int    `json:"result"`
}
