package max

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlicePositive(t *testing.T) {
	testData := map[string]any{
		"key": []uint8{10, 5, 6, 4, 6},
	}
	r := New(10)
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestSliceNegative(t *testing.T) {
	testData := map[string]any{
		"key": []uint8{10, 5, 6, 4, 6},
	}
	r := New(4)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must not have more than 4 items.", err.Error())
	}
}

func TestMapPositive(t *testing.T) {
	testData := map[string]any{
		"key": map[uint8]uint8{10: 3, 5: 2, 6: 4, 4: 1, 16: 5},
	}
	r := New(10)
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestMapNegative(t *testing.T) {
	testData := map[string]any{
		"key": map[uint8]uint8{10: 3, 5: 2, 6: 4, 4: 1, 16: 5},
	}
	r := New(4)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must not have more than 4 items.", err.Error())
	}
}

func TestStringPositive(t *testing.T) {
	testData := map[string]any{
		"key": "asdqerrtda",
	}
	r := New(10)
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestStringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "asdqerrtda",
	}
	r := New(4)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must not be greater than 4 characters.", err.Error())
	}
}

func TestIntPositive(t *testing.T) {
	testData := map[string]any{
		"key": 97,
	}
	r := New(100)
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestIntNegative(t *testing.T) {
	testData := map[string]any{
		"key": 97,
	}
	r := New(94)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must not be greater than 94.", err.Error())
	}
}

func TestFloatPositive(t *testing.T) {
	testData := map[string]any{
		"key": 97.45,
	}
	r := New(100)
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestFloatNegative(t *testing.T) {
	testData := map[string]any{
		"key": 94.1,
	}
	r := New(94)
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must not be greater than 94.", err.Error())
	}
}
