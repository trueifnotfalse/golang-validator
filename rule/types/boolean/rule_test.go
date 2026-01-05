package boolean

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestPositive(t *testing.T) {
	testData := map[string]any{
		"key": true,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestIntNegative(t *testing.T) {
	testData := map[string]any{
		"key": 1,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.boolean", err.Error())
	}
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "true",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.boolean", err.Error())
	}
}

func TestStringLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": "true",
	}
	r := New().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an boolean.", err.Error())
	}
}
