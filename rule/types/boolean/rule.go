package boolean

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/utils"
)

type Rule struct {
	message string
}

func New() *Rule {
	return &Rule{
		message: "The %s must be an boolean.",
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	if utils.IsBool(v) {
		return nil
	}
	return fmt.Errorf(fmt.Sprintf(r.message, key))
}
