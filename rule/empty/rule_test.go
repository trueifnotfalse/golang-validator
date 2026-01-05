package empty

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestStringPositive(t *testing.T) {
	testData := map[string]any{
		"key": "",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestIntPositive(t *testing.T) {
	testData := map[string]any{
		"key": 0,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestArrayPositive(t *testing.T) {
	testData := map[string]any{
		"key": []int8{},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "qwe",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "empty", err.Error())
	}
}

func TestStringLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": "qwe",
	}
	r := New().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be empty.", err.Error())
	}
}

func TestIntNegative(t *testing.T) {
	testData := map[string]any{
		"key": 10,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "empty", err.Error())
	}
}

func TestArrayNegative(t *testing.T) {
	testData := map[string]any{
		"key": []int8{2, 5, 23},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "empty", err.Error())
	}
}
