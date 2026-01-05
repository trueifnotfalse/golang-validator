package float

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestPositive(t *testing.T) {
	testData := map[string]any{
		"key": 20.4,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestIntPositive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.float", err.Error())
	}
}

func TestStringLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := New().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an float.", err.Error())
	}
}
