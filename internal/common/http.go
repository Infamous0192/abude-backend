package common

import "abude-backend/pkg/pagination"

type BasicResponse struct {
	Message string `json:"message"`
}

type ResultResponse struct {
	Result interface{} `json:"result,omitempty"`
}

type GeneralResponse struct {
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

type PaginatedResponse struct {
	Metadata pagination.Metadata `json:"metadata"`
	Result   []interface{}       `json:"result"`
}

type ErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"`
}
