package kgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSONRecoveryHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			switch data := err.(type) {
			case string:
				ctx.JSON(http.StatusInternalServerError, data)
			case error:
				ctx.JSON(http.StatusInternalServerError, data.Error())
			default:
				ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("%v", data))
			}
		}
	}()

	ctx.Next()
}
