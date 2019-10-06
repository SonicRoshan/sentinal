package sentinal

import (
	"reflect"
)

//Validate is used to validate an object
func Validate(object interface{}) (valid bool, output map[string]string, err error) {
	valueOf := reflect.ValueOf(object)
	typeOf := reflect.TypeOf(object)
	output = map[string]string{}
	var msg string

	for i := 0; i < typeOf.NumField(); i++ {
		curField := typeOf.Field(i)

		for functionName, function := range functions {

			value, ok := curField.Tag.Lookup(functionName)

			if ok {
				valid, msg, err = function(
					valueOf.Field(i),
					value,
				)

				if msg != "" {
					output[curField.Name] = msg
					return
				} else if err != nil || !valid {
					return
				}
			}
		}
	}

	return
}
