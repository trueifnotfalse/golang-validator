package validator

import (
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/rule/date"
	"github.com/trueifnotfalse/golang-validator/rule/each"
	"github.com/trueifnotfalse/golang-validator/rule/empty"
	"github.com/trueifnotfalse/golang-validator/rule/in"
	"github.com/trueifnotfalse/golang-validator/rule/ip/v4"
	"github.com/trueifnotfalse/golang-validator/rule/max"
	"github.com/trueifnotfalse/golang-validator/rule/min"
	"github.com/trueifnotfalse/golang-validator/rule/notEmpty"
	"github.com/trueifnotfalse/golang-validator/rule/required"
	"github.com/trueifnotfalse/golang-validator/rule/types/array"
	"github.com/trueifnotfalse/golang-validator/rule/types/boolean"
	"github.com/trueifnotfalse/golang-validator/rule/types/float"
	"github.com/trueifnotfalse/golang-validator/rule/types/integer"
	"github.com/trueifnotfalse/golang-validator/rule/types/object"
	"github.com/trueifnotfalse/golang-validator/rule/types/str"
	"github.com/trueifnotfalse/golang-validator/rule/types/unsignedInteger"
	"github.com/trueifnotfalse/golang-validator/rule/url/http"
)

type Rules map[string][]rule.Interface

func String() rule.Interface {
	return str.New()
}

func Int8() rule.Interface {
	return integer.Int8()
}

func Int16() rule.Interface {
	return integer.Int16()
}

func Int32() rule.Interface {
	return integer.Int32()
}

func Int64() rule.Interface {
	return integer.Int64()
}

func UInt8() rule.Interface {
	return unsignedInteger.UInt8()
}

func UInt16() rule.Interface {
	return unsignedInteger.UInt16()
}

func UInt32() rule.Interface {
	return unsignedInteger.UInt32()
}

func UInt64() rule.Interface {
	return unsignedInteger.UInt64()
}

func Array() rule.Interface {
	return array.New()
}

func Object() rule.Interface {
	return object.New()
}

func Boolean() rule.Interface {
	return boolean.New()
}

func Float() rule.Interface {
	return float.New()
}

func Date(format string) rule.Interface {
	return date.New(format)
}

func Required() rule.Interface {
	return required.New()
}

func Each(rules ...rule.Interface) rule.Interface {
	return each.New(rules...)
}

func In[V int64 | int32 | int16 | int8 | uint64 | uint32 | uint16 | uint8 | string](v []V) rule.Interface  {
	return in.New(v)
}

func Empty() rule.Interface {
	return empty.New()
}

func NotEmpty() rule.Interface {
	return notEmpty.New()
}

func Max(v int64) rule.Interface {
	return max.New(v)
}

func Min(v int64) rule.Interface {
	return min.New(v)
}

func HttpUrl() rule.Interface {
	return http.New()
}

func IpV4() rule.Interface {
	return v4.New()
}
