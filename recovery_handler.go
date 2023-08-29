package kgin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kgin/valid"
	"github.com/xuender/kit/types"
)

func RecoveryHandler(ctx *gin.Context) {
	defer deferRecover(ctx)

	ctx.Next()
}

func deferRecover(ctx *gin.Context) {
	err := recover()
	if err == nil {
		return
	}

	if data, ok := err.(valid.BadRequestError); ok {
		ctx.String(http.StatusBadRequest, data.Err())

		return
	}

	switch data := err.(type) {
	case string:
		if len(data) > 4 && data[3] == ':' {
			if code, err := types.ParseInteger[int](data[:3]); err == nil {
				ctx.String(code, data[4:])

				return
			}
		}

		ctx.String(http.StatusInternalServerError, data)
	case error:
		ctx.String(http.StatusInternalServerError, data.Error())
	default:
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", data))
	}
}
