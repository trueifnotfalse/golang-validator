package min

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"reflect"
)

type Rule struct {
	loc            locale.Interface
	size           int64
	messageArray   string
	messageString  string
	messageNumeric string
}

func New(size int64) rule.Interface {
	return &Rule{
		size:           size,
		messageArray:   "min.array",
		messageString:  "min.string",
		messageNumeric: "min.numeric",
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
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		if int64(rv.Len()) < r.size {
			return errors.New(r.getErrorMessage(r.messageArray, key, r.size))
		}
		return nil
	}
	s := utils.ToString(v)
	if utils.IsString(v) {
		if int64(len(s)) < r.size {
			return errors.New(r.getErrorMessage(r.messageString, key, r.size))
		}
		return nil
	}
	f, ok := utils.ToFloat(s)
	if ok {
		if f < float64(r.size) {
			return errors.New(r.getErrorMessage(r.messageNumeric, key, r.size))
		}
		return nil
	}
	i, ok := utils.ToInt(s)
	if ok {
		if i < r.size {
			return errors.New(r.getErrorMessage(r.messageNumeric, key, r.size))
		}
		return nil
	}

	return nil
}

func (r *Rule) getErrorMessage(message, key string, size int64) string {
	if r.loc == nil {
		return message
	}

	return fmt.Sprintf(r.loc.Translate(message), key, size)
}
