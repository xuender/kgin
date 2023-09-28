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

	switch data := err.(type) {
	case string:
		if len(data) > 4 && data[3] == ':' {
			if code, err := strconv.ParseInt(data[:3], base.Ten, base.SixtyFour); err == nil {
				ctx.String(int(code), data[4:])

				return
			}
		}

		ctx.String(http.StatusInternalServerError, data)
	case valid.BadRequestError:
		ctx.String(http.StatusBadRequest, data.String())
	case error:
		switch {
		case errors.Is(data, gorm.ErrRecordNotFound):
			ctx.String(http.StatusNotFound, data.Error())
		case errors.Is(data, valid.ErrOptimisticLock):
			ctx.String(http.StatusConflict, data.Error())
		default:
			switch {
			case strings.Contains(data.Error(), "UNIQUE"):
				ctx.String(http.StatusBadRequest, data.Error())
			default:
				ctx.String(http.StatusInternalServerError, data.Error())
			}
		}
	default:
		ctx.String(http.StatusInternalServerError, fmt.Sprintf("%v", data))
	}
}
