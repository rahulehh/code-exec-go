package models

type ExecuteRequest struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type ExecuteResponse struct {
	Stdout string `json:"stdout"`
	Error  string `json:"error,omitempty"`
}
