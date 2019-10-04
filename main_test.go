package sentinal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	username string
	age      int `max:"32"`
}

func TestValidation(t *testing.T) {
	assert := assert.New(t)

	userData := user{username: "2", age: 10}
	valid, err := Validate(userData)
	assert.True(valid)
	assert.NoError(err)

	userData = user{username: "2", age: 120}
	valid, err = Validate(userData)
	assert.False(valid)
	assert.NoError(err)

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
