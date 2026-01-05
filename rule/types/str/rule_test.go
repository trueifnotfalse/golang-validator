package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
)

func TestStringPositive(t *testing.T) {
	testData := map[string]any{
		"key": "qwerty",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestIntPositive(t *testing.T) {
	testData := map[string]any{
		"key": "432",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestIntNegative(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.str", err.Error())
	}
}

func TestIntLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := New().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an string.", err.Error())
	}
}

func TestObjectNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[string]any{
			"20":56,
		},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.str", err.Error())
	}
}
