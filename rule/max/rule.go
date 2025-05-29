package max

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"reflect"
)

type Rule struct {
	size    int64
	message map[string]string
}

func New(size int64) rule.Interface {
	return &Rule{
		size: size,
		message: map[string]string{
			"array":   "The %s must not have more than %d items.",
			"string":  "The %s must not be greater than %d characters.",
			"numeric": "The %s must not be greater than %d.",
		},
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
		if int64(rv.Len()) > r.size {
			return fmt.Errorf(r.message["array"], key, r.size)
		}
		return nil
	}
	s := utils.ToString(v)
	if utils.IsString(v) {
		if int64(len(s)) > r.size {
			return fmt.Errorf(r.message["string"], key, r.size)
		}
		return nil
	}
	f, ok := utils.ToFloat(s)
	if ok {
		if f > float64(r.size) {
			return fmt.Errorf(r.message["numeric"], key, r.size)
		}
		return nil
	}
	i, ok := utils.ToInt(s)
	if ok {
		if i > r.size {
			return fmt.Errorf(r.message["numeric"], key, r.size)
		}
		return nil
	}

	return nil
}
