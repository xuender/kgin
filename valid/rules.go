package valid

import (
	"reflect"

	"github.com/xuender/kvalid"
)

func Validation(method string, models ...kvalid.RuleHolder) map[string]*kvalid.Rules {
	ret := map[string]*kvalid.Rules{}

	for _, model := range models {
		val := reflect.ValueOf(model)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		key := val.Type().Name()
		ret[key] = model.Validation(method)
	}

	return ret
}
