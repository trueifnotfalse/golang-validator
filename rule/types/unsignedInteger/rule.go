package unsignedInteger

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"math"
)

type Rule struct {
	max     uint64
	message string
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	i, ok := utils.ToUInt(utils.ToString(v))
	if !ok {
		return fmt.Errorf(r.message, key)
	}
	if i > r.max {
		return fmt.Errorf(r.message, key)
	}

	return nil
}

func UInt8() rule.Interface {
	return &Rule{
		max:     math.MaxUint8,
		message: "The %s must be an uint8.",
	}
}

func UInt16() rule.Interface {
	return &Rule{
		max:     math.MaxUint16,
		message: "The %s must be an uint16.",
	}
}

func UInt32() rule.Interface {
	return &Rule{
		max:     math.MaxUint32,
		message: "The %s must be an uint32.",
	}
}

func UInt64() rule.Interface {
	return &Rule{
		max:     math.MaxUint64,
		message: "The %s must be an uint64.",
	}
}
