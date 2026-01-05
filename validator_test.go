package validator

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/locale/ru"
	"testing"
)

func TestValidatePositive(t *testing.T) {
	testData := map[string]any{
		"name":   "Vasiliy",
		"age":    54,
		"gender": 1,
	}
	d, err := json.Marshal(testData)
	assert.Nil(t, err)
	genderList := []uint8{1, 2}
	rl := map[string][]rule.Interface{
		"name":   {Required(), String()},
		"age":    {Required(), UInt8()},
		"gender": {In(genderList)},
	}
	v := New()
	errs := v.Validate(d, rl)
	assert.Equal(t, 0, len(errs))
}

func TestValidateNegative(t *testing.T) {
	testData := map[string]any{
		"age":    541,
		"gender": 3,
	}
	d, err := json.Marshal(testData)
	assert.Nil(t, err)
	genderList := []uint8{1, 2}
	rl := map[string][]rule.Interface{
		"name":   {Required(), String()},
		"age":    {Required(), UInt8()},
		"gender": {In(genderList)},
	}
	v := New()
	errs := v.Validate(d, rl)
	assert.Equal(t, 3, len(errs))
	assert.Equal(t, "The name field is required.", errs["name"][0].Error())
	assert.Equal(t, "The age must be an uint8.", errs["age"][0].Error())
	assert.Equal(t, "The selected gender is invalid.", errs["gender"][0].Error())
}

func TestValidateEmptyDataPositive(t *testing.T) {
	testData := map[string]any{}
	d, err := json.Marshal(testData)
	assert.Nil(t, err)
	rl := map[string][]rule.Interface{
		"name": {String()},
		"age":  {UInt8()},
	}
	v := New()
	errs := v.Validate(d, rl)
	assert.Equal(t, 0, len(errs))
}

func TestValidateEmptyDataNegative(t *testing.T) {
	testData := map[string]any{}
	d, err := json.Marshal(testData)
	assert.Nil(t, err)
	rl := map[string][]rule.Interface{
		"name": {Required(), String()},
		"age":  {UInt8()},
	}
	v := New()
	errs := v.Validate(d, rl)
	assert.NotEqual(t, 0, len(errs))
}

func TestValidateHttpLocalizationNegative(t *testing.T) {
	testData := map[string]any{
		"key": "tcp://exmple.com/",
	}
	d, err := json.Marshal(testData)
	assert.Nil(t, err)
	rl := map[string][]rule.Interface{
		"key": {HttpUrl()},
	}
	v := New().SetLocale(ru.New())
	errs := v.Validate(d, rl)
	if !assert.NotEqual(t, 0, len(errs)) {
		return
	}
	if !assert.NotEqual(t, 0, len(errs["key"])) {
		return
	}
	assert.Equal(t, "Формат key не соответствует HTTP-URL.", errs["key"][0].Error())
}
