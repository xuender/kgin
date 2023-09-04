package valid

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kvalid"
)

// Bind 数据绑定并校验.
func Bind[T kvalid.RuleHolder[T]](ctx *gin.Context, method string, old T) error {
	newT := NewPoint(old)

	if err := ctx.Bind(newT); err != nil {
		return NewBadRequestError(err)
	}

	return NewBadRequestError(old.Validation(method).Bind(newT, old))
}

func NewPoint[T any](src T) T {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	target := reflect.New(typ)
	ret, _ := target.Interface().(T)

	return ret
}
