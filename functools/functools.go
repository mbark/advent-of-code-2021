package functools

import (
	"reflect"
)

func Every(s interface{}, fn func(i int) bool) bool {
	rv := reflect.ValueOf(s)
	every := true
	for i := 0; i < rv.Len(); i++ {
		if !fn(i) {
			every = false
			break
		}
	}

	return every
}

func Some(s interface{}, fn func(i int) bool) bool {
	rv := reflect.ValueOf(s)
	some := false
	for i := 0; i < rv.Len(); i++ {
		if fn(i) {
			some = true
			break
		}
	}

	return some
}
