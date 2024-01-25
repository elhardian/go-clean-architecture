package paginationHelper

import (
	"math"
	"strconv"
)

// Help To Create Pagination In Many Data
func AddPagination(totalData int, page string) (*Page, error) {
	if page == "" {
		return nil, ErrorPageEmpty
	}

	newPage, err := strconv.Atoi(page)
	if err != nil {
		return nil, ErrorPageInvalid
	}

	if newPage <= 0 {
		return nil, ErrorPage
	}

	limitData := 10
	totalPage := int(math.Ceil(float64(totalData) / float64(limitData)))

	last := newPage * limitData
	first := last - limitData

	zeroPage := &Page{TotalPage: 1, Data: "-", Page: newPage}
	if totalPage == 0 && newPage == 1 {
		return zeroPage, nil
	}

	if newPage > totalPage {
		return nil, ErrorMaxPage
	}

	dataCount := strconv.Itoa(first) + " - " + strconv.Itoa(last-1)
	if totalData < last {
		dataCount = strconv.Itoa(first) + " - " + strconv.Itoa(totalData-1)
	}

	hasNext := false
	if newPage < totalPage {
		hasNext = true
	}

	hasPrev := false
	if newPage > 1 {
		hasPrev = true
	}

	pages := &Page{
		TotalPage: totalPage,
		TotalRows: totalData,
		Page:      newPage,
		Data:      dataCount,
		HasNext:   hasNext,
		HasPrev:   hasPrev,
	}

	if totalPage == 0 {
		return zeroPage, nil
	}

	pages.First = first
	pages.Last = last
	if last > totalData {
		pages.First = first
		pages.Last = totalData
	}

	return pages, nil
}
