package kgin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kgin/valid"
	"github.com/xuender/kit/base"
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
		ctx.String(http.StatusBadRequest, data.String())

		return
	}

	switch data := err.(type) {
	case string:
		if len(data) > 4 && data[3] == ':' {
			if code, err := strconv.ParseInt(data[:3], base.Ten, base.SixtyFour); err == nil {
				ctx.String(int(code), data[4:])

				return
			}
		}

		ctx.String(http.StatusInternalServerError, data)
	case NotFoundError, NotFoundIDError, NotFoundKeyError:
		err, _ := data.(error)
		ctx.String(http.StatusNotFound, err.Error())
	case error:
		ctx.String(http.StatusInternalServerError, data.Error())
	default:
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", data))
	}
}
