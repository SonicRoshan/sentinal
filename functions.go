package sentinal

import (
	"reflect"
)

func maxInclusive(value reflect.Value, validationData string) (valid bool, err error) {
	number, err := getNumber(value)

	if err != nil {
		return false, err
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return false, err
	}

	valid = number <= maxValue
	return
}

func minInclusive(value reflect.Value, validationData string) (valid bool, err error) {
	number, err := getNumber(value)

	if err != nil {
		return false, err
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return false, err
	}

	valid = number >= maxValue
	return
}

func maxExclusive(value reflect.Value, validationData string) (valid bool, err error) {
	number, err := getNumber(value)

	if err != nil {
		return false, err
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return false, err
	}

	valid = number < maxValue
	return
}

func minExclusive(value reflect.Value, validationData string) (valid bool, err error) {
	number, err := getNumber(value)

	if err != nil {
		return false, err
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return false, err
	}

	valid = number > maxValue
	return
}
