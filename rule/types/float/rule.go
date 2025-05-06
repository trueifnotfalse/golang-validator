package float

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
)

type Rule struct {
	message string
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	if utils.IsString(v) {
		return fmt.Errorf(fmt.Sprintf(r.message, key))
	}
	if !utils.IsFloat(utils.ToString(v)) {
		return fmt.Errorf(fmt.Sprintf(r.message, key))
	}

	return nil
}

func New() rule.Interface {
	return &Rule{
		message: "The %s must be an float.",
	}
}
