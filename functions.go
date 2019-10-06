package sentinal

import (
	"reflect"
	"strconv"
)

func maxInclusive(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	number, err := getNumber(value)

	if err != nil {
		return false, "", err
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return false, "", err
	}

	valid = number <= maxValue
	if !valid {
		msg = "inclusive max value is " + strconv.FormatInt(int64(maxValue), 10)
	}

	return
}

func minInclusive(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	number, err := getNumber(value)

	if err != nil {
		return
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return
	}

	valid = number >= maxValue
	if !valid {
		msg = "inclusive min value is " + strconv.FormatInt(int64(maxValue), 10)
	}
	return
}

func maxExclusive(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	number, err := getNumber(value)

	if err != nil {
		return
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return
	}

	valid = number < maxValue
	if !valid {
		msg = "exclusive max value is " + strconv.FormatInt(int64(maxValue), 10)
	}
	return
}

func minExclusive(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	number, err := getNumber(value)

	if err != nil {
		return
	}

	maxValue, err := getNumberFromString(validationData)

	if err != nil {
		return
	}

	valid = number > maxValue
	if !valid {
		msg = "exclusive min value is " + strconv.FormatInt(int64(maxValue), 10)
	}
	return
}
