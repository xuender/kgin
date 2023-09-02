package valid

import (
	"log/slog"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kvalid"
)

// Bind 数据绑定并校验.
func Bind[T kvalid.RuleHolder](ctx *gin.Context, method string, old T) (T, error) {
	newT := NewPoint(old)

	if err := ctx.Bind(newT); err != nil {
		slog.Error("bind", err)

		return old, NewBadRequestError(err)
	}

	slog.Info("bind", "new", newT)

	return old, old.Validation(method).Bind(newT, old)
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
