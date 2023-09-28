package kgin

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kgin/valid"
	"github.com/xuender/kit/base"
	"gorm.io/gorm"
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

	switch val := err.(type) {
	case string:
		if len(val) > 4 && val[3] == ':' {
			if code, err := strconv.ParseInt(val[:3], base.Ten, base.SixtyFour); err == nil {
				ctx.String(int(code), val[4:])

				return
			}
		}

		ctx.String(http.StatusInternalServerError, val)
	case valid.BadRequestError:
		ctx.String(http.StatusBadRequest, val.String())
	case error:
		outputError(val, ctx)
	default:
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", val))
	}
}

func outputError(err error, ctx *gin.Context) {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		ctx.String(http.StatusNotFound, err.Error())
	case errors.Is(err, valid.ErrOptimisticLock):
		ctx.String(http.StatusConflict, err.Error())
	default:
		switch {
		case strings.Contains(err.Error(), "UNIQUE"):
			ctx.String(http.StatusBadRequest, err.Error())
		default:
			ctx.String(http.StatusInternalServerError, err.Error())
		}
	}
}
