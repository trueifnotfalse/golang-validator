package array

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestIntPositive(t *testing.T) {
	testData := map[string]any{
		"key": []int{1, 2},
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
		assert.Equal(t, "types.array", err.Error())
	}
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "1,2,3",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.array", err.Error())
	}
}

func TestStringLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": "1,2,3",
	}
	r := New().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an array.", err.Error())
	}
}
