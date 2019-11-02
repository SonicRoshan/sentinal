package sentinal

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dataCame bool

type user struct {
	username string
	age      int `max:"32"`
}

type test2 struct {
	name string `max:"32"`
}

type test3 struct {
	custom string `custom:"data"`
}

func customFunc(value reflect.Value, data string) (bool, string, error) {
	dataCame = true
	return true, "", nil
}

func TestValidation(t *testing.T) {
	assert := assert.New(t)

	userSchema := schemaType{
		"age": map[string]string{
			"max": "11",
		},
	}

	test2Schema := schemaType{
		"name": map[string]string{
			"max": "32",
		},
	}

	test3Schema := schemaType{
		"custom": map[string]string{
			"custom": "data",
		},
	}

	userData := user{username: "2", age: 10}
	valid, validationData, err := Validate(userData, userSchema)
	assert.True(valid)
	assert.Empty(validationData)
	assert.NoError(err)

	userData = user{username: "2", age: 120}
	valid, validationData, err = Validate(userData, userSchema)
	assert.False(valid)
	assert.NoError(err)
	assert.NotEmpty(validationData)

	data := test2{"A"}
	valid, validationData, err = Validate(data, test2Schema)
	assert.False(valid)
	assert.Error(err)
	assert.Empty(validationData)

	data3 := test3{"tt"}
	valid, validationData, err = Validate(data3, test3Schema, map[string]functionType{
		"custom": customFunc,
	})
	assert.True(valid)
	assert.Empty(validationData)
	assert.NoError(err)
	assert.True(dataCame)

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
