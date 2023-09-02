package valid

import (
	"reflect"

	"github.com/xuender/kvalid"
)

func GetRules(method string, models ...kvalid.RuleHolder) map[string]*kvalid.Rules {
	ret := map[string]*kvalid.Rules{}

	for _, model := range models {
		val := reflect.ValueOf(model)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		ret[val.Type().Name()] = model.Validation(method)
	}

	return ret
}
