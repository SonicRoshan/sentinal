package sentinal

import (
	"reflect"
	"strconv"
	"strings"
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

func from(value reflect.Value, validationData string) (bool, string, error) {
	fromList := strings.Split(validationData, ",")
	for _, item := range fromList {
		if item == value.String() {
			return true, "", nil
		}
	}
	return false, "value not in list", nil
}

func notFrom(value reflect.Value, validationData string) (bool, string, error) {
	valid, _, _ := from(value, validationData)
	if valid {
		return false, "value in list", nil
	}
	return true, "", nil
}
