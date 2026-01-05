package integer

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"math"
)

type Rule struct {
	loc     locale.Interface
	min     int64
	max     int64
	message string
}

func (r *Rule) SetLocale(v locale.Interface) rule.Interface {
	r.loc = v
	return r
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	if utils.IsString(v) {
		return errors.New(r.getErrorMessage(key))
	}
	i, ok := utils.ToInt(utils.ToString(v))
	if !ok {
		return errors.New(r.getErrorMessage(key))
	}
	if i < r.min {
		return errors.New(r.getErrorMessage(key))
	}

	if i > r.max {
		return errors.New(r.getErrorMessage(key))
	}

	return nil
}

func Int8() rule.Interface {
	return &Rule{
		min:     math.MinInt8,
		max:     math.MaxInt8,
		message: "types.int8",
	}
}

func Int16() rule.Interface {
	return &Rule{
		min:     math.MinInt16,
		max:     math.MaxInt16,
		message: "types.int16",
	}
}

func Int32() rule.Interface {
	return &Rule{
		min:     math.MinInt32,
		max:     math.MaxInt32,
		message: "types.int32",
	}
}

func Int64() rule.Interface {
	return &Rule{
		min:     math.MinInt64,
		max:     math.MaxInt64,
		message: "types.int64",
	}
}

func (r *Rule) getErrorMessage(key string) string {
	if r.loc == nil {
		return r.message
	}

	return fmt.Sprintf(r.loc.Translate(r.message), key)
}
