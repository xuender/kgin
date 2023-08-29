package db

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/xuender/kgin/valid"
	"github.com/xuender/kit/types"
	"gorm.io/gorm"
)

// nolint: gochecknoglobals
var (
	_limits = [...]string{"limit", "l", "length"}
	_offset = [...]string{"offset", "o", "start", "s"}
)

func Query[T valid.Valid](ctx *gin.Context, gdb *gorm.DB) *Result[T] {
	var (
		limit  = 100
		offset = 0
		count  int64
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

	gdb = gdb.Model(new(T))

	lo.Must0(gdb.Count(&count).Error)

	list := []T{}
	lo.Must0(gdb.Limit(limit).Offset(offset).Find(&list).Error)

	return &Result[T]{
		Count:  count,
		Limit:  limit,
		Offset: offset,
		Data:   list,
	}
}
