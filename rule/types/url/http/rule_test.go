package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpPositive(t *testing.T) {
	testData := map[string]any{
		"key": "http://exmple.com/",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestHttpsPositive(t *testing.T) {
	testData := map[string]any{
		"key": "https://exmple.com/",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestSchemaNegative(t *testing.T) {
	testData := map[string]any{
		"key": "tcp://exmple.com/",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key format is not HTTP URL.", err.Error())
	}
}

func TestHostEmptyNegative(t *testing.T) {
	testData := map[string]any{
		"key": "http://",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key format is not HTTP URL.", err.Error())
	}
}
