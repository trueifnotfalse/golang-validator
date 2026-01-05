package required

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestPositive(t *testing.T) {
	testData := map[string]any{
		"key": 10,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestNegative(t *testing.T) {
	testData := map[string]any{
		"keya": 10,
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "required", err.Error())
	}
}

func TestLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"keya": 10,
	}
	r := New().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key field is required.", err.Error())
	}
}
