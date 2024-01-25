package paginationHelper

type Page struct {
	TotalRows int    `json:"total_rows"`
	TotalPage int    `json:"total_page"`
	Page      int    `json:"page"`
	Data      string `json:"data"`
	HasNext   bool   `json:"has_next"`
	HasPrev   bool   `json:"has_prev"`
	First     int    `json:"-"`
	Last      int    `json:"-"`
}
