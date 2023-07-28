package kgin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HTMLHandler(html string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html;charset=utf-8")
		ctx.String(http.StatusOK, html)
	}
}
