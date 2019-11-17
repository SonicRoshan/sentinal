package sentinal

import (
	"reflect"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
)

type test struct {
	a int
}

func testFunction(
	assert *assert.Assertions,
	function func(reflect.Value, string) (bool, string, error),
	value interface{},
	invalidValue interface{},
	validationData string,
	invalidValidation ...string) {

	valid, msg, err := function(reflect.ValueOf(value), validationData)
	assert.True(valid)
	assert.NoError(err)
	assert.Zero(msg)

	valid, msg, err = function(reflect.ValueOf(invalidValue), validationData)
	assert.False(valid)
	assert.NoError(err)
	assert.NotZero(msg)

	if invalidValidation != nil {
		valid, msg, err = function(reflect.ValueOf(test{1}), validationData)
		assert.False(valid)
		assert.Error(err)
		assert.Zero(msg)

		for _, data := range invalidValidation {
			valid, msg, err = function(reflect.ValueOf(value), data)
			assert.False(valid)
			assert.Error(err)
			assert.Zero(msg)
		}
	}
}

func testFunctionWithBoolInputs(
	assert *assert.Assertions,
	function func(reflect.Value, string) (bool, string, error),
	value interface{},
	invalidValue interface{}) {

	valid, msg, err := function(reflect.ValueOf(value), "true")
	assert.True(valid)
	assert.NoError(err)
	assert.Zero(msg)

	valid, msg, err = function(reflect.ValueOf(invalidValue), "true")
	assert.False(valid)
	assert.NoError(err)
	assert.NotZero(msg)

	valid, msg, err = function(reflect.ValueOf("somethingElse"), "false")
	assert.True(valid)
	assert.NoError(err)
	assert.Zero(msg)
}

func Test(t *testing.T) {
	assert := assert.New(t)

	testFunction(assert, maxInclusive, 5, 6, "5", "abc")
	testFunction(assert, maxExclusive, 4.9, 5, "5", "abc")
	testFunction(assert, minInclusive, 5, 4, "5", "abc")
	testFunction(assert, minExclusive, 5.11119, 5.1, "5.1", "abc")
	testFunction(assert, from, "a", "e", "a,b,c,d")
	testFunction(assert, notFrom, "e", "a", "a,b,c,d")
	testFunction(assert, notEmpty, "a", "", "true")
	testFunction(assert, maxLength, "abc", "abcde", "3", "abc")
	testFunction(assert, minLength, "abcde", "abc", "5", "abc")
	testFunction(assert, notContains, "de", "test", "est")
	testFunction(assert, contains, "test", "de", "est")
	testFunctionWithBoolInputs(assert, isEmail, "sonicroshan122@gmail.com", "invalidEmail")

	function := simpleOverlay(govalidator.HasUpperCase, "test")
	testFunction(assert, function, "Valid", "invaid", "true")

	function = simpleOverlayReverse(govalidator.HasUpperCase, "test")
	testFunction(assert, function, "valid", "Invaid", "true")

}
