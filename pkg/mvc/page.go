package mvc

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func PageQuery(ctx *gin.Context) (int, int) {
	page := ParseIntDefault(ctx.Query("page"), 1)
	pageSize := ParseIntDefault(ctx.Query("page_size"), 10)
	return page, pageSize
}

func ParseIntDefault(v string, defaultValue int) int {
	i, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return i
}

func ParseInt(v string) int {
	return ParseIntDefault(v, 0)
}

func QueryInt(ctx *gin.Context, name string) int {
	return ParseInt(ctx.Query(name))
}
