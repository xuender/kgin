package kgin

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/hash"
)

const (
	_Etag         = "Etag"
	_IfNoneMatch  = "If-None-Match"
	_CacheControl = "Cache-Control"
)

func EtagHandler(maxAge uint) gin.HandlerFunc {
	cacheControl := fmt.Sprintf("max-age=%d", maxAge)

	return func(ctx *gin.Context) {
		if ctx.Request.Method != http.MethodGet {
			ctx.Next()

			return
		}

		writer := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = writer

		ctx.Next()

		var (
			bys = writer.body.Bytes()
			tag = hash.SipHashHex(bys)
		)

		writer.ResponseWriter.Header().Set(_Etag, tag)

		if tag == ctx.Request.Header.Get(_IfNoneMatch) {
			ctx.Writer.Header().Set(_CacheControl, cacheControl)
			ctx.Status(http.StatusNotModified)

			return
		}

		if _, err := writer.ResponseWriter.Write(bys); err != nil {
			slog.Error("etag", err)
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}
