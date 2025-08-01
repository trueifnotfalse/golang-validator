package date

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"time"
)

type Rule struct {
	format  string
	message string
}

func New(format string) rule.Interface {
	return &Rule{
		format:  format,
		message: "The %s does not match the format %s.",
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	_, err := time.Parse(r.format, utils.ToString(v))
	if err != nil {
		return fmt.Errorf(r.message, key, r.format)
	}

	return nil
}
