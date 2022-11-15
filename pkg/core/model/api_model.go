package model

import "time"

type UserRequestInfo struct {
	UserRequestID string `json:"userRequestId"`
	UserRequestIP string `json:"userRequestIp"`
}

type ApiInfoResponse struct {
	Name   string    `json:"name"`
	Status string    `json:"status"`
	Time   time.Time `json:"time"`
}

type ApiErrorResponse struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
	Cause      error  `json:"-"`
}

type DataResponse struct {
	Status bool `json:"status"`
}
