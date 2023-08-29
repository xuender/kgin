package valid

import (
	"reflect"
	"strings"
	"sync"
)

// nolint: gochecknoglobals
var _cache = &Cache{data: map[string]int{}}

type Cache struct {
	data  map[string]int
	mutex sync.Mutex
}

func (p *Cache) Get(src any, name string) int {
	typ := reflect.TypeOf(src)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	key := typ.Name() + "_" + name

	p.mutex.Lock()
	defer p.mutex.Unlock()

	if i, has := p.data[key]; has {
		return i
	}

	for num := 0; num < typ.NumField(); num++ {
		field := typ.Field(num)
		if field.Name == name ||
			strings.Contains(string(field.Tag), `"`+name) {
			p.data[key] = num

			return num
		}
	}

	return -1
}
