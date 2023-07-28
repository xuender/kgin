package kgin

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/logs"
)

func LogHandler(ctx *gin.Context) {
	start := time.Now()

	ctx.Next()

	query := ctx.Request.URL.RawQuery
	if query != "" {
		query = "?" + query
	}

	logs.I.Printf("[%d] %v %s (%s) %s%s",
		ctx.Writer.Status(),
		time.Since(start),
		ctx.ClientIP(),
		ctx.Request.Method,
		ctx.Request.URL.Path,
		query,
	)
}
