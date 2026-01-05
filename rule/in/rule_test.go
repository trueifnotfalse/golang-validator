package in

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestPositive(t *testing.T) {
	testData := map[string]any{
		"key": 6,
	}
	r := New([]uint8{3, 6, 91})
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestNegative(t *testing.T) {
	testData := map[string]any{
		"key": 10,
	}
	r := New([]uint8{3, 6, 91})
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "in", err.Error())
	}
}

func TestLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": 10,
	}
	r := New([]uint8{3, 6, 91}).SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The selected key is invalid.", err.Error())
	}
}
