package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty" query:"limit"`
	Page       int         `json:"page,omitempty" query:"page"`
	Sort       string      `json:"sort,omitempty" query:"sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

var allowedSortFields = map[string]bool{
	"id":    true,
	"title": true,
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	fields := strings.Split(p.Sort, " ")
	if len(fields) > 1 && allowedSortFields[fields[0]] {
		return p.Sort
	}
	return "created_at desc"
}

func GetPagination(c *gin.Context) (*Pagination, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return nil, err
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		return nil, err
	}

	sort := c.DefaultQuery("sort", "")

	// Handle sorting direction
	direction := c.Query("sortDesc")
	if direction != "" {
		switch direction {
		case "true":
			sort += " desc"
		case "false":
			sort += " asc"
		default:
			return nil, fmt.Errorf("invalid sort direction")
		}
	}

	// Create and return pagination object
	p := &Pagination{
		Page:  page,
		Limit: limit,
		Sort:  sort,
	}

	return p, nil
}
