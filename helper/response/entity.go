package response

import paginationHelper "github.com/elhardian/go-clean-architecture/helper/pagination"

type Response struct {
	Status     bool                   `json:"status"`
	StatusCode int                    `json:"status_code"`
	Message    interface{}            `json:"message"`
	Code       string                 `json:"code"`
	Pagination *paginationHelper.Page `json:"pagination,omitempty"`
	Data       interface{}            `json:"data,omitempty"`
}
