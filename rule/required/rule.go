package required

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
)

type Rule struct {
	message string
}

func New() rule.Interface {
	return &Rule{
		message: "The %s field is required.",
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	_, ok := values[key]
	if !ok {
		return fmt.Errorf(fmt.Sprintf(r.message, key))
	}

	return nil
}
