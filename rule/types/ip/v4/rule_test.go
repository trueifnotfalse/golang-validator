package v4

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositive(t *testing.T) {
	testData := map[string]any{
		"key": "192.168.10.56",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestWhitePositive(t *testing.T) {
	testData := map[string]any{
		"key": "8.8.8.8",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestLocalHostPositive(t *testing.T) {
	testData := map[string]any{
		"key": "127.0.0.1",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestRandomStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "1dzxjc6suj3dnadjkfbn",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be a valid IPv4 address.", err.Error())
	}
}

func TestRangeNegative(t *testing.T) {
	testData := map[string]any{
		"key": "8.318.8.8",
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be a valid IPv4 address.", err.Error())
	}
}
