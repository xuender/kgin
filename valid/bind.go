package valid

import (
	"log/slog"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kit/set"
)

// Bind 数据绑定并校验.
func Bind[T Valid](ctx *gin.Context, method string, old T) (T, error) {
	newT := NewPoint(old)

	if err := ctx.Bind(newT); err != nil {
		slog.Error("bind", err)

		return old, NewBadRequestError(err)
	}

	slog.Info("bind", "new", newT)

	return Value(newT, old, method)
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

// Value Post并校验.
func Value[T Valid](src, target T, method string) (T, error) {
	if err := src.Validate(method); err != nil {
		return target, NewBadRequestError(err)
	}

	srcVal := reflect.ValueOf(src)
	targetVal := reflect.ValueOf(target)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if targetVal.Kind() == reflect.Ptr {
		targetVal = targetVal.Elem()
	}

	set := set.NewSet[int]()

	for _, vali := range src.Validation(method).Validators() {
		index := _cache.Get(src, vali.Name())
		if index < 0 || set.Has(index) {
			continue
		}

		field := targetVal.Field(index)

		field.Set(srcVal.Field(index))
		set.Add(index)
	}

	return target, nil
}
