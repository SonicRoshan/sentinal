package sentinal

import (
	"reflect"
)

//Validate is used to validate an object
func Validate(object interface{}) (bool, error) {
	valueOf := reflect.ValueOf(object)
	typeOf := reflect.TypeOf(object)

	for i := 0; i < typeOf.NumField(); i++ {
		curField := typeOf.Field(i)

		for functionName, function := range functions {

			value, ok := curField.Tag.Lookup(functionName)

			if ok {
				valid, err := function(
					valueOf.Field(i),
					value,
				)

				if !valid || err != nil {
					return valid, err
				}
			}
		}
	}

	return true, nil
}
