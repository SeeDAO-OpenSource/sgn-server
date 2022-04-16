package mvc

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PageQuery(ctx *gin.Context) (int, int) {
	page := parseInt(ctx.Query("page"), 1)
	pageSize := parseInt(ctx.Query("page_size"), 10)
	return page, pageSize
}

func parseInt(v string, defaultValue int) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return i
}
