package notEmpty

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
}

func New() rule.Interface {
	return &Rule{
		message: "not.empty",
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
		if rv.Len() == 0 {
			return errors.New(r.getErrorMessage(key))
		}
		return nil
	}
	s := utils.ToString(v)
	if utils.IsString(v) {
		if len(s) == 0 {
			return errors.New(r.getErrorMessage(key))
		}
		return nil
	}
	ui, ok := utils.ToUInt(s)
	if ok {
		if ui == 0 {
			return errors.New(r.getErrorMessage(key))
		}
		return nil
	}
	i, ok := utils.ToInt(s)
	if ok {
		if i == 0 {
			return errors.New(r.getErrorMessage(key))
		}
		return nil
	}

	return nil
}

func (r *Rule) getErrorMessage(key string) string {
	if r.loc == nil {
		return r.message
	}

	return fmt.Sprintf(r.loc.Translate(r.message), key)
}
