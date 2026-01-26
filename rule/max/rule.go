package max

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"reflect"
	"unicode/utf8"
)

type  Rule [T int|float64] struct {
	loc            locale.Interface
	size           T
	messageArray   string
	messageString  string
	messageNumeric string
}

func New[T int|float64](size T) rule.Interface {
	return &Rule[T]{
		size:           size,
		messageArray:   "max.array",
		messageString:  "max.string",
		messageNumeric: "max.numeric",
	}
}

func (r *Rule[T]) SetLocale(v locale.Interface) rule.Interface {
	r.loc = v
	return r
}

func (r *Rule[T]) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		if int64(rv.Len()) > int64(r.size) {
			return errors.New(r.getErrorMessage(r.messageArray, key, r.size))
		}
		return nil
	}
	s := utils.ToString(v)
	if utils.IsString(v) {
		if int64(utf8.RuneCountInString(s)) > int64(r.size) {
			return errors.New(r.getErrorMessage(r.messageString, key, r.size))
		}
		return nil
	}
	f, ok := utils.ToFloat(s)
	if ok {
		if f > float64(r.size) {
			return errors.New(r.getErrorMessage(r.messageNumeric, key, r.size))
		}
		return nil
	}
	i, ok := utils.ToInt(s)
	if ok {
		if i > int64(r.size) {
			return errors.New(r.getErrorMessage(r.messageNumeric, key, r.size))
		}
		return nil
	}

	return nil
}

func (r *Rule[T]) getErrorMessage(message, key string, size T) string {
	if r.loc == nil {
		return message
	}

	return fmt.Sprintf(r.loc.Translate(message), key, size)
}
