package kgin

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/logs"
)

const _calldepth = 4

func LogHandler(ctx *gin.Context) {
	start := time.Now()

	ctx.Next()

	query := ctx.Request.URL.RawQuery
	if query != "" {
		query = "?" + query
	}

	_ = logs.I.Output(_calldepth, fmt.Sprintf("[%d] %v %s (%s) %s%s",
		ctx.Writer.Status(),
		time.Since(start),
		ctx.ClientIP(),
		ctx.Request.Method,
		ctx.Request.URL.Path,
		query,
	))
}
