package unsignedInteger

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
	max     uint64
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
	i, ok := utils.ToUInt(utils.ToString(v))
	if !ok {
		return errors.New(r.getErrorMessage(key))
	}
	if i > r.max {
		return errors.New(r.getErrorMessage(key))
	}

	return nil
}

func UInt8() rule.Interface {
	return &Rule{
		max:     math.MaxUint8,
		message: "types.uint8",
	}
}

func UInt16() rule.Interface {
	return &Rule{
		max:     math.MaxUint16,
		message: "types.uint16",
	}
}

func UInt32() rule.Interface {
	return &Rule{
		max:     math.MaxUint32,
		message: "types.uint32",
	}
}

func UInt64() rule.Interface {
	return &Rule{
		max:     math.MaxUint64,
		message: "types.uint64",
	}
}

func (r *Rule) getErrorMessage(key string) string {
	if r.loc == nil {
		return r.message
	}

	return fmt.Sprintf(r.loc.Translate(r.message), key)
}
