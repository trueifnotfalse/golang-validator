package required

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
)

type Rule struct {
	loc     locale.Interface
	message string
}

func New() rule.Interface {
	return &Rule{
		message: "required",
	}
}

func (r *Rule) SetLocale(v locale.Interface) rule.Interface {
	r.loc = v
	return r
}

func (r *Rule) Valid(key string, values map[string]any) error {
	_, ok := values[key]
	if !ok {
		return errors.New(r.getErrorMessage(key))
	}

	return nil
}

func (r *Rule) getErrorMessage(key string) string {
	if r.loc == nil {
		return r.message
	}

	return fmt.Sprintf(r.loc.Translate(r.message), key)
}
