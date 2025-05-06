package array

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/utils"
)

type Rule struct {
	message string
}

func New() *Rule {
	return &Rule{
		message: "The %s must be an array.",
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	if !utils.IsArray(v) {
		return fmt.Errorf(fmt.Sprintf(r.message, key))
	}

	return nil
}
