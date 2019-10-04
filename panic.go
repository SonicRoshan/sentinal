package sentinal

import "errors"

//handlePanic is used to handle panic and return error
func handlePanic(err *error, errorStr string) {
	if r := recover(); r != nil {
		*err = errors.New(errorStr)
	}
}
