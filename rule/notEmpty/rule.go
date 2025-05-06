package notEmpty

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"reflect"
)

type Rule struct {
	message string
}

func New() rule.Interface {
	return &Rule{
		message: "The %s must not be empty.",
	}
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
			return fmt.Errorf(fmt.Sprintf(r.message, key))
		}
		return nil
	}
	s := utils.ToString(v)
	if utils.IsString(v) {
		if len(s) == 0 {
			return fmt.Errorf(fmt.Sprintf(r.message, key))
		}
		return nil
	}
	ui, ok := utils.ToUInt(s)
	if ok {
		if ui == 0 {
			return fmt.Errorf(fmt.Sprintf(r.message, key))
		}
		return nil
	}
	i, ok := utils.ToInt(s)
	if ok {
		if i == 0 {
			return fmt.Errorf(fmt.Sprintf(r.message, key))
		}
		return nil
	}

	return nil
}
