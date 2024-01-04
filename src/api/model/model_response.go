package model

type Response struct {
	RequestId string `json:"requestId,omitempty"`

	StatusMessage string `json:"statusMessage,omitempty"`

	StatusCode int `json:"statusCode,omitempty"`

	Pods *[]Pod `json:"pods,omitempty"`
}
