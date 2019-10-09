package sentinal

import (
	"reflect"

	"github.com/pkg/errors"
)

//validateField is used to validate a single field against all tags
func validateField(tags reflect.StructTag, value reflect.Value) (bool, []string, error) {
	msgs := []string{}

	for functionName, function := range functions {

		tagValue, ok := tags.Lookup(functionName)

		if ok {
			_, msg, err := function(
				value,
				tagValue,
			)

			if msg != "" {
				msgs = append(msgs, msg)
			} else if err != nil {
				return false, []string{}, err
			}
		}
	}

	return len(msgs) <= 0, msgs, nil
}

//Validate is used to validate an object
func Validate(
	object interface{},
	customFunctionsArg ...map[string]functionType) (bool, map[string][]string, error) {

	valueOf := reflect.ValueOf(object)
	typeOf := reflect.TypeOf(object)
	output := map[string][]string{}

	if customFunctionsArg != nil {
		for name, function := range customFunctionsArg[0] {
			functions[name] = function
		}
	}

	for i := 0; i < typeOf.NumField(); i++ {
		curField := typeOf.Field(i)
		curValue := valueOf.Field(i)
		valid, msgs, err := validateField(curField.Tag, curValue)

		if err != nil {
			err = errors.Wrap(err, "Error while validating field")
			return false, map[string][]string{}, err
		} else if !valid {
			output[curField.Name] = msgs
		}

	}

	return len(output) <= 0, output, nil
}
