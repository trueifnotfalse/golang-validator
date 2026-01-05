package integer

import (
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/locale/en"
	"testing"
)

func TestInt8Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := Int8()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt16Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := Int16()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt32Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := Int32()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt64Positive(t *testing.T) {
	testData := map[string]any{
		"key": 20,
	}
	r := Int64()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt8BelowZeroPositive(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := Int8()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt8Negative(t *testing.T) {
	testData := map[string]any{
		"key": 290,
	}
	r := Int8()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int8", err.Error())
	}
}

func TestInt8StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := Int8()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int8", err.Error())
	}
}

func TestInt8LocaleNegative(t *testing.T) {
	testData := map[string]any{
		"key": 290,
	}
	r := Int8().SetLocale(en.New())
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "The key must be an int8.", err.Error())
	}
}

func TestInt16BelowZeroPositive(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := Int16()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt16Negative(t *testing.T) {
	testData := map[string]any{
		"key": 65657,
	}
	r := Int16()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int16", err.Error())
	}
}

func TestInt16StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := Int16()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int16", err.Error())
	}
}

func TestInt32BelowZeroPositive(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := Int32()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt32Negative(t *testing.T) {
	testData := map[string]any{
		"key": 4294967395,
	}
	r := Int32()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int32", err.Error())
	}
}

func TestInt32StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := Int32()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int32", err.Error())
	}
}

func TestInt64BelowZeroPositive(t *testing.T) {
	testData := map[string]any{
		"key": -20,
	}
	r := Int64()
	err := r.Valid("key", testData)
	assert.Nil(t, err)
}

func TestInt64Negative(t *testing.T) {
	testData := map[string]any{
		"key": uint64(9223372036854975807),
	}
	r := Int64()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int64", err.Error())
	}
}

func TestInt64StringNegative(t *testing.T) {
	testData := map[string]any{
		"key": "20",
	}
	r := Int64()
	err := r.Valid("key", testData)
	assert.NotNil(t, err)
	if err != nil {
		assert.Equal(t, "types.int64", err.Error())
	}
}
