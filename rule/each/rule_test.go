package each

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"github.com/trueifnotfalse/golang-validator/rule/types/date"
	"github.com/trueifnotfalse/golang-validator/rule/types/integer"
	"github.com/trueifnotfalse/golang-validator/rule/types/str"
	"testing"
	"time"
)

func TestIntPositive(t *testing.T) {
	testData := map[string]any{
		"key": []int16{45, 45, 22},
	}
	r := New(integer.Int16())
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestDatePositive(t *testing.T) {
	testData := map[string]any{
		"key": []string{"2025-07-01"},
	}
	r := New(date.New(time.DateOnly))
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestNegative(t *testing.T) {
	testData := map[string]any{
		"key": 33,
	}
	r := New(str.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "each", err.Error())
	}
}

func TestLocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": 33,
	}
	r := New(str.New()).SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an array.", err.Error())
	}
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": []uint16{4, 6, 2},
	}
	r := New(str.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.str", err.Error())
	}
}

func TestDateNegative(t *testing.T) {
	testData := map[string]any{
		"key": []string{"qwe", "2025-07-01"},
	}
	r := New(date.New(time.DateOnly))
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.date", err.Error())
	}
}
