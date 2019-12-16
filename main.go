package sentinal

import (
	"reflect"

	"github.com/ZeroTechh/hades"
	"github.com/pkg/errors"
)

// converts map[string]interface{} to map[string]map[string]string
func processYAMLData(data map[string]interface{}) map[string]map[string]string {
	output := map[string]map[string]string{}
	for key, value := range data {
		toAdd := value.(map[string]interface{})
		stringMap := map[string]string{}
		for key2, value2 := range toAdd {
			stringMap[key2] = value2.(string)
		}
		output[key] = stringMap
	}
	return output
}

// validateField is used to validate a single field with its args from schema
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

// Validate is used to validate an object with a golang defined schema
func Validate(
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

// ValidateWithYAML is used to validate an object with yaml schema
func ValidateWithYAML(
	object interface{},
	schemaFile string,
	schemaPaths []string,
	customFunctionsArg ...map[string]func(reflect.Value, string) (bool, string, error)) (
	bool,
	map[string][]string,
	error) {
	config := hades.GetConfig(schemaFile, schemaPaths)
	schema := processYAMLData(config.Data)

	return Validate(object, schema, customFunctionsArg...)
}
