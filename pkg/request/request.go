package request

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetSort(c *gin.Context, fallback map[string]string) map[string]string {
	sort := c.QueryArray("sort")
	paramSort := make(map[string]string)
	if len(sort) > 0 {
		//split with : as separator
		for _, v := range sort {
			if strings.Contains(v, ":") {
				arrSort := strings.Split(v, ":")
				paramSort[arrSort[0]] = arrSort[1]
			}
		}
	} else {
		paramSort = fallback
	}
	return paramSort
}
func GetPage(c *gin.Context) int {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}
	return page
}
func GetPerPage(c *gin.Context) int {
	perPageStr := c.DefaultQuery("per_page", "10")
	perPage, _ := strconv.Atoi(perPageStr)
	switch {
	case perPage > 100:
		perPage = 100
	case perPage <= 0:
		perPage = 10
	}
	return perPage
}
func GetMultipleFilter(c *gin.Context, key string, filter map[string]interface{}) map[string]interface{} {
	filterStr := c.DefaultQuery(key, "")
	if filterStr == "" {
		return filter
	}
	filterArr := strings.Split(filterStr, ",")
	var filters []interface{}
	for _, v := range filterArr {
		filters = append(filters, v)
	}
	if len(filters) > 0 {
		filter[key] = filters
	}
	return filter
}
func GetMultipleFilters(c *gin.Context, key []string, filter map[string]interface{}) map[string]interface{} {
	for _, v := range key {
		filter = GetMultipleFilter(c, v, filter)
	}
	return filter
}

func GetFilter(c *gin.Context, key string, filter map[string]interface{}) map[string]interface{} {
	value := c.DefaultQuery(key, "")
	if value != "" {
		filter[key] = value
	}
	return filter
}
func GetFilters(c *gin.Context, key []string, filter map[string]interface{}) map[string]interface{} {
	for _, v := range key {
		filter = GetFilter(c, v, filter)
	}
	return filter
}

func GetID(c *gin.Context) uint64 {
	idStr := c.Param("id")
	if idStr == "" {
		return 0
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0
	}
	return id
}

func GetIDByKey(c *gin.Context, key string) uint64 {
	idStr := c.Param(key)
	if idStr == "" {
		return 0
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return 0
	}
	return id
}
