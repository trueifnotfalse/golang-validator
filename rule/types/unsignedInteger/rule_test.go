package unsignedInteger

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestUInt8Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := UInt8()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestUInt16Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := UInt16()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestUInt32Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := UInt32()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestUInt64Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := UInt64()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestUInt8BelowZeroNegative(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := UInt8()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint8", err.Error())
	}
}

func TestUInt8Negative(t *testing.T) {
	testData := map[string]any{
		"key": 290,
	}
	r := UInt8()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint8", err.Error())
	}
}

func TestUInt8StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := UInt8()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint8", err.Error())
	}
}

func TestUInt8LocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": 290,
	}
	r := UInt8().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an uint8.", err.Error())
	}
}

func TestUInt16BelowZeroNegative(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := UInt16()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint16", err.Error())
	}
}

func TestUInt16Negative(t *testing.T) {
	testData := map[string]any{
		"key": 65657,
	}
	r := UInt16()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint16", err.Error())
	}
}

func TestUInt16StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := UInt16()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint16", err.Error())
	}
}

func TestUInt32BelowZeroNegative(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := UInt32()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint32", err.Error())
	}
}

func TestUInt32Negative(t *testing.T) {
	testData := map[string]any{
		"key": 4294967395,
	}
	r := UInt32()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint32", err.Error())
	}
}

func TestUInt32StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := UInt32()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint32", err.Error())
	}
}

func TestUInt64BelowZeroNegative(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := UInt64()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint64", err.Error())
	}
}

func TestUInt64StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := UInt64()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.uint64", err.Error())
	}
}
