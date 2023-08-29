package db

import (
	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/types"
	"gorm.io/gorm"
)

// nolint: gochecknoglobals
var (
	_limits = [...]string{"limit", "l", "length"}
	_offset = [...]string{"offset", "o", "start", "s"}
)

func Query(ctx *gin.Context, gdb *gorm.DB) *gorm.DB {
	var (
		limit  = 100
		offset = 0
	)

	for _, key := range _limits {
		if value := ctx.Query(key); value != "" {
			if num, err := types.ParseInteger[int](value); err == nil {
				limit = num

				break
			}
		}
	}

	for _, key := range _offset {
		if value := ctx.Query(key); value != "" {
			if num, err := types.ParseInteger[int](value); err == nil {
				offset = num

				break
			}
		}
	}

	return gdb.Limit(limit).Offset(offset)
}
