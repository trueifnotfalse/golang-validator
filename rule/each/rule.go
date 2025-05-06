package each

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"reflect"
)

type Rule struct {
	message string
	rules   []rule.Interface
}

func New(rules []rule.Interface) rule.Interface {
	return &Rule{
		message: "The %s must be an array.",
		rules:   rules,
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	s := reflect.ValueOf(v)
	if s.Kind() != reflect.Slice {
		return fmt.Errorf(fmt.Sprintf(r.message, key))
	}
	if s.IsNil() {
		return nil
	}
	var (
		j   int
		err error
	)
	for i := 0; i < s.Len(); i++ {
		for j = 0; j < len(r.rules); j++ {
			key = fmt.Sprintf("%s.%d", key, i)
			err = r.rules[j].Valid(key, map[string]any{key: s.Index(i).Interface()})
			if err != nil {
				return err
			}
		}
	}

	return nil
}
