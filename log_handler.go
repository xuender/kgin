package kgin

import (
	"context"
	"log/slog"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

const _calldepth = 4

func LogHandler(ctx *gin.Context) {
	start := time.Now()

	ctx.Next()

	query := ctx.Request.URL.RawQuery
	if query != "" {
		query = "?" + query
	}

	var pcs [1]uintptr

	runtime.Callers(_calldepth, pcs[:])

	record := slog.NewRecord(time.Now(), slog.LevelInfo, "gin", pcs[0])
	args := []any{
		"status",
		ctx.Writer.Status(),
		"elapsed",
		time.Since(start),
		"ip",
		ctx.ClientIP(),
		"method",
		ctx.Request.Method,
		"path",
		ctx.Request.URL.Path,
	}

	if query != "" {
		args = append(args, "query", query)
	}

	record.Add(args...)
	_ = slog.Default().Handler().Handle(context.Background(), record)
}
