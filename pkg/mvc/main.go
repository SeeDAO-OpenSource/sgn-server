package mvc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataResult struct {
	Status  int
	Message string
	Data    interface{}
	Success bool
}

func Error(ctx *gin.Context, err error) {
	status := http.StatusInternalServerError
	ctx.JSON(status, &DataResult{
		Status:  status,
		Message: err.Error(),
		Data:    false,
		Success: false,
	})
}

func Fail(ctx *gin.Context, code int, message string) {
	status := http.StatusInternalServerError
	ctx.JSON(status, &DataResult{
		Status:  status,
		Message: message,
		Data:    false,
		Success: false,
	})
}

func Ok(ctx *gin.Context, data interface{}) {
	status := http.StatusOK
	ctx.JSON(status, &DataResult{
		Status:  status,
		Message: "",
		Data:    data,
		Success: true,
	})
}
