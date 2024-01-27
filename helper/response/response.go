package response

import (
	"encoding/json"
	"net/http"

	paginationHelper "github.com/elhardian/go-clean-architecture/helper/pagination"
)

// Return JSON
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// Return JSON When Response Is Success
func OnSuccess(w http.ResponseWriter, status bool, code string, data interface{}, pagination *paginationHelper.Page) {
	if code == "" {
		code = "OK"
	}

	res := &Response{
		Status:     status,
		StatusCode: http.StatusOK,
		Message:    "Success",
		Code:       code,
		Data:       data,
		Pagination: pagination,
	}

	WriteJSON(w, http.StatusOK, res)
}

// Return JSON When Response Is Failed
func OnError(w http.ResponseWriter, status bool, statusCode int, message interface{}, code string) {
	res := &Response{
		Status:     status,
		StatusCode: statusCode,
		Message:    message,
		Code:       code,
	}

	WriteJSON(w, statusCode, res)
}
