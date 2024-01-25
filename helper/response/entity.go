package jsonHelper

import paginationHelper "github.com/elhardian/go-clean-architecture/helper/pagination"

type response struct {
	Status     bool                   `json:"status"`
	StatusCode int                    `json:"status_code"`
	Message    interface{}            `json:"message"`
	Code       string                 `json:"code"`
	Pagination *paginationHelper.Page `json:"pagination,omitempty"`
	Data       interface{}            `json:"data,omitempty"`
}
