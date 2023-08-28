package kgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecoveryHandler(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			switch data := err.(type) {
			case string:
				ctx.String(http.StatusInternalServerError, data)
			case error:
				ctx.String(http.StatusInternalServerError, data.Error())
			default:
				ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", data))
			}
		}
	}()

	ctx.Next()
}
