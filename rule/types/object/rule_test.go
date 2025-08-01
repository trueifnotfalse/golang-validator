package object

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[string]uint8{"key": 8},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestNegative(t *testing.T) {
	testData := map[string]any{
		"key": []uint8{10, 5, 6, 4, 6},
	}
	r := New()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an object.", err.Error())
	}
}
