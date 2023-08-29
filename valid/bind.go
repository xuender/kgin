package valid

import (
	"log/slog"
	"reflect"

	"github.com/gin-gonic/gin"
)

// BindPut 数据绑定并校验.
func BindPut[T Put](ctx *gin.Context, old T) (T, error) {
	newT := NewPoint(old)

	if err := ctx.Bind(newT); err != nil {
		slog.Error("bind", err)

		return old, err
	}

	slog.Info("bind", "new", newT)

	return PutValue(newT, old)
}

func NewPoint[T any](src T) T {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	target := reflect.New(typ)

	ret, has := target.Interface().(T)
	if has {
		return ret
	}

	slog.Error("NewPoint error")

	return ret
}

// BindPost 数据绑定并校验.
func BindPost[T Post](ctx *gin.Context, old T) (T, error) {
	newT := NewPoint(old)

	if err := ctx.Bind(newT); err != nil {
		return old, err
	}

	return PostValue(newT, old)
}

// PostValue Post并校验.
func PostValue[T Post](src, target T) (T, error) {
	if err := src.ValidatePost(); err != nil {
		return target, err
	}

	newVal := reflect.ValueOf(src)
	val := reflect.ValueOf(target)

	if newVal.Kind() == reflect.Ptr {
		newVal = newVal.Elem()
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for _, vali := range src.ValidationPost().Validators() {
		i := _cache.Get(src, vali.Name())
		if i < 0 {
			continue
		}

		val.Field(i).Set(newVal.Field(i))
	}

	return target, nil
}

// PutValue Put并校验.
func PutValue[T Put](src, target T) (T, error) {
	if err := src.ValidatePut(); err != nil {
		return target, err
	}

	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)
	typ := reflect.TypeOf(target)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if targetVal.Kind() == reflect.Ptr {
		targetVal = targetVal.Elem()
		typ = typ.Elem()
	}

	src.ValidationPut().Validators()

	for _, vali := range src.ValidationPut().Validators() {
		slog.Info("put1", "field", vali.Name())

		if _, has := typ.FieldByName(vali.Name()); !has {
			continue
		}

		slog.Info("put2", "field", vali.Name())

		targetVal.FieldByName(vali.Name()).
			Set(srcVal.FieldByName(vali.Name()))
	}

	return target, nil
}
