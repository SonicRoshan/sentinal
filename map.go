package sentinal

import (
	"reflect"
)

var functions = map[string]func(reflect.Value, string) (bool, string, error){
	"max":          maxInclusive,
	"min":          minInclusive,
	"maxExclusive": maxExclusive,
	"minExclusive": minExclusive,
	"from":         from,
	"notFrom":      notFrom,
	"notEmpty":     notEmpty,
}
