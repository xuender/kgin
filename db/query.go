package db

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/xuender/kgin/valid"
	"github.com/xuender/kit/types"
	"gorm.io/gorm"
)

func Query[T valid.Valid](ctx *gin.Context, gdb *gorm.DB) *Result[T] {
	var (
		limit  = 100
		offset = 0
		count  int64
	)

	if num, err := types.ParseInteger[int](ctx.DefaultQuery("limit", "100")); err == nil {
		limit = num
	}

	if num, err := types.ParseInteger[int](ctx.DefaultQuery("offset", "0")); err == nil {
		offset = num
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
