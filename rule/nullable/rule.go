package nullable

import (
	"github.com/trueifnotfalse/golang-validator/interface/rule"
)

type Rule struct {
	rules []rule.Interface
}

func New(rules ...rule.Interface) rule.Interface {
	return &Rule{
		rules: rules,
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	if v == nil {
		return nil
	}
	var err error
	for i := 0; i < len(r.rules); i++ {
		err = r.rules[i].Valid(key, values)
		if err != nil {
			return err
		}
	}

	return nil
}
