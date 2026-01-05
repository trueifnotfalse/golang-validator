package each

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
)

type Rule struct {
	loc     locale.Interface
	message string
	rules   []rule.Interface
}

func New(rules ...rule.Interface) rule.Interface {
	return &Rule{
		message: "each",
		rules:   rules,
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
	s := reflect.ValueOf(v)
	if s.Kind() != reflect.Slice {
		return errors.New(r.getErrorMessage(key))
	}
	if s.IsNil() {
		return nil
	}
	var (
		j   int
		err error
		k   string
	)
	for i := 0; i < s.Len(); i++ {
		for j = 0; j < len(r.rules); j++ {
			k = fmt.Sprintf("%s.%d", key, i)
			err = r.rules[j].SetLocale(r.loc).Valid(k, map[string]any{k: s.Index(i).Interface()})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Rule) getErrorMessage(key string) string {
	if r.loc == nil {
		return r.message
	}

	return fmt.Sprintf(r.loc.Translate(r.message), key)
}
