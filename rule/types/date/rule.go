package date

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"time"
)

type Rule struct {
	loc     locale.Interface
	format  string
	message string
}

func New(format string) rule.Interface {
	return &Rule{
		format:  format,
		message: "types.date",
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
	_, err := time.Parse(r.format, utils.ToString(v))
	if err != nil {
		return errors.New(r.getErrorMessage(key, r.format))
	}

	return nil
}

func (r *Rule) getErrorMessage(key, format string) string {
	if r.loc == nil {
		return r.message
	}

	return fmt.Sprintf(r.loc.Translate(r.message), key, format)
}
