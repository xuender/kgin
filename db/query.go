package db

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/xuender/kit/los"
	"github.com/xuender/kit/types"
	"gorm.io/gorm"
)

func Query[T any](ctx *gin.Context, gdb *gorm.DB) *Result[T] {
	return BaseStrQuery[T](gdb.Model(new(T)), ctx.DefaultQuery("limit", "10"), ctx.DefaultQuery("offset", "0"))
}

func QueryModel[T any](ctx *gin.Context, gdb *gorm.DB) *Result[T] {
	return BaseStrQuery[T](gdb, ctx.DefaultQuery("limit", "10"), ctx.DefaultQuery("offset", "0"))
}

func BaseStrQuery[T any](gdb *gorm.DB, limit, offset string) *Result[T] {
	var (
		numLimit  = 0
		numOffset = 0
	)

	if num, err := types.ParseInteger[int](limit); err == nil {
		numLimit = num
	}

	if num, err := types.ParseInteger[int](offset); err == nil {
		numOffset = num
	}

	return BaseQuery[T](gdb, numLimit, numOffset)
}

func BaseQuery[T any](gdb *gorm.DB, limit, offset int) *Result[T] {
	if limit <= 0 {
		limit = 10
	}

	var count int64

	lo.Must0(gdb.Count(&count).Error)

	list := []T{}
	los.Must0(gdb.Limit(limit).Offset(offset).Find(&list).Error)

	return &Result[T]{
		Count:  count,
		Limit:  limit,
		Offset: offset,
		Data:   list,
	}
}
