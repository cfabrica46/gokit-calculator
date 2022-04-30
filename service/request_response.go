package service

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
