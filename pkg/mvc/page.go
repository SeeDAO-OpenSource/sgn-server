package mvc

import "github.com/gin-gonic/gin"

type Page struct {
	Limmit int
	Skip   int
}

func PageQuery(ctx *gin.Context) (int, int) {
	var page Page
	ctx.BindQuery(&page)
	return page.Limmit, page.Skip
}
