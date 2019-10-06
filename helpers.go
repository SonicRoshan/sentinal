package sentinal

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// ErrInvalidType occurs when data is of invalid type
var ErrInvalidType = errors.New("Invalid Data Type")

func getNumber(number reflect.Value) (float64, error) {

	if number.Kind() == reflect.Int {
		return float64(number.Int()), nil
	} else if number.Kind() == reflect.Float32 ||
		number.Kind() == reflect.Float64 {

		return number.Float(), nil
	}

	return 0.0, errors.Wrap(ErrInvalidType, "getNumber got non number type data")
}

func getNumberFromString(numberStr string) (float64, error) {
	if strings.Contains(numberStr, ".") {
		floatValue, err := strconv.ParseFloat(numberStr, 64)
		return floatValue, errors.Wrap(err, "ParseFloat returned error")
	}

	value, err := strconv.ParseInt(numberStr, 10, 64)

	return float64(value), errors.Wrap(err, "ParseInt returned error")
}
