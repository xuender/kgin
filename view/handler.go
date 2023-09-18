package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler 页面访问记录.
func Handler(key string, viewer Viewer) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Next()

		if ctx.Request.Method != http.MethodGet {
			return
		}

		page := ctx.Param(key)
		if page == "" {
			page = ctx.Request.URL.String()
		}

		viewer.View(page, ctx.RemoteIP())
	}
}
