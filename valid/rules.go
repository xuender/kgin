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

func ValidationPut(models ...Put) map[string]Rules {
	ret := map[string]Rules{}

	for _, model := range models {
		val := reflect.ValueOf(model)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		key := val.Type().Name()
		ret[key] = putRules(model)
	}

	return ret
}

func ValidationPost(models ...Post) map[string]Rules {
	ret := map[string]Rules{}

	for _, model := range models {
		val := reflect.ValueOf(model)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}

		key := val.Type().Name()
		ret[key] = postRules(model)
	}

	return ret
}

func postRules(put Post) Rules {
	return newRules(put.ValidationPost().Validators())
}

func putRules(put Put) Rules {
	return newRules(put.ValidationPut().Validators())
}
