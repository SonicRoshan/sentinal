package sentinal

import (
	"reflect"

	"github.com/pkg/errors"
)

func validateField(
	value reflect.Value,
	args map[string]string) (bool, []string, error) {

	msgs := []string{}
	for name, arg := range args {
		valid, msg, err := functions[name](value, arg)
		if !valid && err == nil {
			msgs = append(msgs, msg)
		} else if err != nil {
			return false, []string{}, err
		}
	}

	return len(msgs) <= 0, msgs, nil
}

// Validate is used to validate an object
func Validate(
	object interface{},
	schema map[string]map[string]string,
	customFunctionsArg ...map[string]functionType) (
	bool,
	map[string][]string,
	error) {

	valueOf := reflect.ValueOf(object)
	typeOf := reflect.TypeOf(object)
	output := map[string][]string{}

	// If custom functions are given, add it to functions map
	if customFunctionsArg != nil {
		for name, function := range customFunctionsArg[0] {
			functions[name] = function
		}
	}

	for i := 0; i < typeOf.NumField(); i++ {
		curField := typeOf.Field(i)

		if args, ok := schema[curField.Name]; ok {
			valid, msgs, err := validateField(valueOf.Field(i), args)
			if err != nil {
				err = errors.Wrap(err, "Error while validating field")
				return false, map[string][]string{}, err
			} else if !valid {
				output[curField.Name] = msgs
			}
		}

	}

	return len(output) <= 0, output, nil
}
