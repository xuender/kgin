package valid

import (
	"reflect"
	"strings"

	"github.com/AgentCosmic/xvalid"
)

type Rules map[string][]xvalid.Validator

func newRules(validators []xvalid.Validator) Rules {
	rmap := make(map[string][]xvalid.Validator)

	for _, val := range validators {
		if !val.HTMLCompatible() {
			continue
		}

		name := val.Name()
		if index := strings.Index(name, ","); index > 0 {
			name = name[:index]
		}

		rules, has := rmap[name]
		if !has {
			rules = []xvalid.Validator{}
		}

		rules = append(rules, val)
		rmap[name] = rules
	}

	return rmap
}

func Validation(method string, models ...Valid) map[string]Rules {
	ret := map[string]Rules{}

	for _, model := range models {
		val := reflect.ValueOf(model)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		key := val.Type().Name()
		ret[key] = getRules(model, method)
	}

	return ret
}

func getRules(put Valid, method string) Rules {
	return newRules(put.Validation(method).Validators())
}
