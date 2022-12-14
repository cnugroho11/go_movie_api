package utils

import (
	"github.com/cnugroho11/movie_api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(ctx *gin.Context) models.Pagination {
	limit := 5
	page := 1
	sort := "id"
	query := ctx.Request.URL.Query()

	for key, val := range query {
		queryValue := val[len(val)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
