package sentinal

import (
	"reflect"
	"testing"

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
	invalidValidation string) {

	valid, msg, err := function(reflect.ValueOf(value), validationData)
	assert.True(valid)
	assert.NoError(err)
	assert.Zero(msg)

	valid, msg, err = function(reflect.ValueOf(invalidValue), validationData)
	assert.False(valid)
	assert.NoError(err)
	assert.NotZero(msg)

	valid, msg, err = function(reflect.ValueOf(test{1}), validationData)
	assert.False(valid)
	assert.Error(err)
	assert.Zero(msg)

	valid, msg, err = function(reflect.ValueOf(value), invalidValidation)
	assert.False(valid)
	assert.Error(err)
	assert.Zero(msg)
}

func Test(t *testing.T) {
	assert := assert.New(t)

	testFunction(assert, maxInclusive, 5, 6, "5", "abc")
	testFunction(assert, maxExclusive, 4.9, 5, "5", "abc")
	testFunction(assert, minInclusive, 5, 4, "5", "abc")
	testFunction(assert, minExclusive, 5.11119, 5.1, "5.1", "abc")
}
