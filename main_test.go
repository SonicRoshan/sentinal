package sentinal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataCame bool

type test2 struct {
	a string
	b int
}

func customFunc(value reflect.Value, data string) (bool, string, error) {
	dataCame = true
	return true, "", nil
}

func TestValidation(t *testing.T) {
	assert := assert.New(t)

	schema := []string{"schema", "../schema"}

	// Testing that sentinal outputs true for a valid struct
	data := test2{"a", 100}
	valid, msg, err := ValidateWithYAML(data, "schema.yaml", schema,
		map[string]func(reflect.Value, string) (bool, string, error){
			"custom": customFunc,
		})
	assert.True(valid)
	assert.Empty(msg)
	assert.NoError(err)

	// Testing that sentinal outputs false for a invalid struct
	data = test2{"abc", -10}
	valid, msg, err = ValidateWithYAML(data, "schema.yaml", schema,
		map[string]func(reflect.Value, string) (bool, string, error){
			"custom": customFunc,
		})
	assert.False(valid)
	assert.NotEmpty(msg)
	assert.NoError(err)

	// Testing that sentinal outputs error for a invalid schema data
	data = test2{"abc", -10}
	valid, msg, err = ValidateWithYAML(data, "schema3.yaml", schema)
	assert.False(valid)
	assert.Empty(msg)
	assert.Error(err)

	// Testing field validation
	data = test2{a: "a"}
	valid, msg, err = ValidateFieldsWithYAML(data, "schema.yaml", schema,
		map[string]func(reflect.Value, string) (bool, string, error){
			"custom": customFunc,
		})
	assert.True(valid)
	assert.Empty(msg)
	assert.NoError(err)

	data = test2{a: "abcdf"}
	valid, msg, err = ValidateFieldsWithYAML(data, "schema.yaml", schema)
	assert.False(valid)
	assert.NotEmpty(msg)
	assert.NoError(err)

	data = test2{a: "a"}
	valid, msg, err = ValidateFieldsWithYAML(data, "schema3.yaml", schema)
	assert.False(valid)
	assert.Empty(msg)
	assert.Error(err)
}

func TestPanicHandler(t *testing.T) {
	assert := assert.New(t)

	testFunc := func() (err error) {
		defer handlePanic(&err, "Error")
		panic("123")
	}

	err := testFunc()
	assert.Error(err)
}
