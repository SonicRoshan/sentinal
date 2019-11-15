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

func notEmpty(value reflect.Value, validationData string) (bool, string, error) {
	if reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface()) &&
		validationData == "true" {
		return false, "value is nil", nil
	}
	return true, "", nil
}

func maxLength(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	defer handlePanic(&err, "Invalid Data")

	max, err := getNumberFromString(validationData)
	if err != nil {
		return
	}

	valid = value.Len() <= int(max)
	if !valid {
		msg = "Length is greater than max"
	}

	return
}

func minLength(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	defer handlePanic(&err, "Invalid Data")

	min, err := getNumberFromString(validationData)
	if err != nil {
		return
	}

	valid = value.Len() >= int(min)
	if !valid {
		msg = "Length is smaller than min"
	}

	return
}

func contains(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	defer handlePanic(&err, "Invalid Data")
	items := strings.Split(validationData, ",")
	for _, item := range items {
		contains := strings.Contains(value.String(), item)
		if !contains {
			return false, "Value does not contain " + item, nil
		}
	}
	return true, "", nil
}

func notContains(value reflect.Value, validationData string) (valid bool, msg string, err error) {
	defer handlePanic(&err, "Invalid Data")
	items := strings.Split(validationData, ",")
	for _, item := range items {
		contains := strings.Contains(value.String(), item)
		if contains {
			return false, "Value does contain " + item, nil
		}
	}
	return true, "", nil
}
