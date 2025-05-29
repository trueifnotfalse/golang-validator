package integer

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"math"
)

type Rule struct {
	min     int64
	max     int64
	message string
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	if utils.IsString(v) {
		return fmt.Errorf(r.message, key)
	}
	i, ok := utils.ToInt(utils.ToString(v))
	if !ok {
		return fmt.Errorf(r.message, key)
	}
	if i < r.min {
		return fmt.Errorf(r.message, key)
	}

	if i > r.max {
		return fmt.Errorf(r.message, key)
	}

	return nil
}

func Int8() rule.Interface {
	return &Rule{
		min:     math.MinInt8,
		max:     math.MaxInt8,
		message: "The %s must be an int8.",
	}
}

func Int16() rule.Interface {
	return &Rule{
		min:     math.MinInt16,
		max:     math.MaxInt16,
		message: "The %s must be an int16.",
	}
}

func Int32() rule.Interface {
	return &Rule{
		min:     math.MinInt32,
		max:     math.MaxInt32,
		message: "The %s must be an int32.",
	}
}

func Int64() rule.Interface {
	return &Rule{
		min:     math.MinInt64,
		max:     math.MaxInt64,
		message: "The %s must be an int64.",
	}
}
