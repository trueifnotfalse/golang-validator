package date

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
)

func TestDatePositive(t *testing.T) {
	testData := map[string]any{
		"key": "2024-05-21",
	}
	r := New(time.DateOnly)
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestDateTimePositive(t *testing.T) {
	testData := map[string]any{
		"key": "2024-05-21T04:03:00Z",
	}
	r := New(time.RFC3339)
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestNegative(t *testing.T) {
	testData := map[string]any{
		"key": "2024-05-21T04:03:00Z",
	}
	r := New(time.DateOnly)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.date", err.Error())
	}
}

func TestLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": "2024-05-21T04:03:00Z",
	}
	r := New(time.DateOnly).SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key does not match the format 2006-01-02.", err.Error())
	}
}

func TestIntNegative(t *testing.T) {
	testData := map[string]any{
		"key": 290,
	}
	r := New(time.DateOnly)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.date", err.Error())
	}
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "asdfgh",
	}
	r := New(time.DateOnly)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.date", err.Error())
	}
}
