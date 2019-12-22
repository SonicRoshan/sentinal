package sentinal

import (
	"reflect"

	"github.com/ZeroTechh/hades"
	"github.com/pkg/errors"
)

// ValidateFields is used to only validate non nill fields in an object
func ValidateFields(
	object interface{},
	schema map[string]map[string]string,
	customFunctionsArg ...map[string]func(reflect.Value, string) (bool, string, error)) (
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
		if !valueOf.Field(i).IsZero() {
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
	}

	return len(output) <= 0, output, nil
}

// ValidateFieldsWithYAML is used to validate non nill fields in an object with yaml schema
func ValidateFieldsWithYAML(
	object interface{},
	schemaFile string,
	schemaPaths []string,
	customFunctionsArg ...map[string]func(reflect.Value, string) (bool, string, error)) (
	bool,
	map[string][]string,
	error) {
	config := hades.GetConfig(schemaFile, schemaPaths)
	schema := processYAMLData(config.Data)

	return ValidateFields(object, schema, customFunctionsArg...)
}
