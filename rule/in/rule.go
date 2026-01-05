package in

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"reflect"
)

type Rule struct {
	loc     locale.Interface
	message string
	f       func(value any, t reflect.Kind) bool
}

func New[V int64 | int32 | int16 | int8 | uint64 | uint32 | uint16 | uint8 | string](v []V) rule.Interface {
	return &Rule{
		message: "in",
		f: func(value any, t reflect.Kind) bool {
			if len(v) == 0 {
				return false
			}
			a := reflect.ValueOf(v[0])
			switch a.Kind() {
			case reflect.Slice, reflect.Array, reflect.Map:
				return false
			case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if t != reflect.Int64 && t != reflect.Uint64 {
					return false
				}
			case reflect.String:
				if t != reflect.String {
					return false
				}
			}
			var expected string
			asString := utils.ToString(value)
			for i := 0; i < len(v); i++ {
				expected = utils.ToString(v[i])
				if asString == expected {
					return true
				}
			}
			return false
		},
	}
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
	actual := reflect.ValueOf(v)
	if actual.Kind() == reflect.Slice || actual.Kind() == reflect.Array || actual.Kind() == reflect.Map {
		return errors.New(r.getErrorMessage(key))
	}
	t := reflect.Bool
	s := utils.ToString(v)
	if utils.IsString(v) {
		t = reflect.String
	} else if actual.Kind() == reflect.Bool {
		t = reflect.Bool
	} else if utils.IsUInt(s) {
		t = reflect.Uint64
	} else if utils.IsInt(s) {
		t = reflect.Int64
	}
	if r.f(v, t) {
		return nil
	}

	return errors.New(r.getErrorMessage(key))
}

func (r *Rule) getErrorMessage(key string) string {
	if r.loc == nil {
		return r.message
	}

	return fmt.Sprintf(r.loc.Translate(r.message), key)
}
