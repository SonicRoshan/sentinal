package sentinal

import (
	"reflect"
)

var functions = map[string]func(reflect.Value, string) (bool, error){
	"max":          maxInclusive,
	"min":          minInclusive,
	"maxExclusive": maxExclusive,
	"minExclusive": minExclusive,
}
