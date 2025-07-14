package nullable

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/rule/date"
	"github.com/trueifnotfalse/golang-validator/rule/types/str"
	"testing"
	"time"
)

func TestNullPositive(t *testing.T) {
	testData := map[string]any{
		"key": nil,
	}
	r := New(str.New(), date.New(time.RFC3339))
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestPositive(t *testing.T) {
	testData := map[string]any{
		"key": time.Now().Format(time.RFC3339),
	}
	r := New(str.New(), date.New(time.RFC3339))
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": []uint16{4, 6, 2},
	}
	r := New(str.New(), date.New(time.RFC3339))
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an string.", err.Error())
	}
}

func TestDateNegative(t *testing.T) {
	testData := map[string]any{
		"key": time.Now().Format(time.DateOnly),
	}
	r := New(str.New(), date.New(time.RFC3339))
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key does not match the format 2006-01-02T15:04:05Z07:00.", err.Error())
	}
}
