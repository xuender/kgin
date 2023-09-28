package kgin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/hash"
)

const (
	Etag          = "Etag"
	IfNoneMatch   = "If-None-Match"
	_cacheControl = "Cache-Control"
	_cache        = "max-age=10"
)

func String(ctx *gin.Context, code int, format string, values ...any) {
	if statusPass(ctx, code) {
		ctx.String(code, format, values...)

		return
	}

	var bys []byte

	if len(values) == 0 {
		bys = []byte(format)
	} else {
		bys = []byte(fmt.Sprintf(format, values...))
	}

	if isCache(ctx, bys) {
		return
	}

	ctx.String(code, format, values...)
}

func JSON(ctx *gin.Context, code int, obj any) {
	if statusPass(ctx, code) {
		ctx.JSON(code, obj)

		return
	}

	bys, err := json.Marshal(obj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)

		return
	}

	if isCache(ctx, bys) {
		return
	}

	ctx.JSON(code, obj)
}

func statusPass(ctx *gin.Context, code int) bool {
	return ctx.Request.Method != http.MethodGet ||
		code < http.StatusOK ||
		code >= http.StatusMultipleChoices
}

func isCache(ctx *gin.Context, data []byte) bool {
	tag := hash.SipHashHex(data)
	ctx.Writer.Header().Set(Etag, tag)

	if tag == ctx.Request.Header.Get(IfNoneMatch) {
		ctx.Writer.Header().Set(_cacheControl, _cache)
		ctx.Status(http.StatusNotModified)

		return true
	}

	return false
}
