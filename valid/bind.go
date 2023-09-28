package valid

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/xuender/kvalid"
)

// Bind 数据绑定并校验.
func Bind[T kvalid.RuleHolder[T]](ctx *gin.Context, method string, old T) error {
	newT, _ := NewPoint(old)

	if err := ctx.Bind(newT); err != nil {
		return NewBadRequestError(err)
	}

	return NewBadRequestError(old.Validation(method).Bind(newT, old))
}

func LockBind[T kvalid.RuleHolder[T]](ctx *gin.Context, method string, old T) error {
	newT, lock := NewPoint(old)

	if err := ctx.Bind(newT); err != nil {
		return NewBadRequestError(err)
	}

	if lock {
		if err := checkUpdateAt(newT, old); err != nil {
			return err
		}
	}

	return NewBadRequestError(old.Validation(method).Bind(newT, old))
}

func checkUpdateAt[T kvalid.RuleHolder[T]](newT, old T) error {
	newV := reflect.ValueOf(newT)
	if newV.Kind() == reflect.Pointer {
		newV = newV.Elem()
	}

	newV = newV.FieldByName("UpdatedAt")

	oldV := reflect.ValueOf(old)
	if oldV.Kind() == reflect.Pointer {
		oldV = oldV.Elem()
	}

	oldV = oldV.FieldByName("UpdatedAt")

	if !newV.Equal(oldV) {
		nv1 := newV.MethodByName("Sec").Call(nil)[0]
		ov1 := oldV.MethodByName("Sec").Call(nil)[0]

		if !nv1.Equal(ov1) {
			return ErrOptimisticLock
		}
	}

	return nil
}

func NewPoint[T any](src T) (T, bool) {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	target := reflect.New(typ)
	ret, _ := target.Interface().(T)
	_, has := typ.FieldByName("UpdatedAt")

	return ret, has
}
