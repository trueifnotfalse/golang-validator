package v4

import (
	"errors"
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"net"
)

type Rule struct {
	loc     locale.Interface
	message string
}

func New() rule.Interface {
	return &Rule{
		message: "types.ip.v4",
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
	ip := net.ParseIP(utils.ToString(v))
	if ip == nil || ip.To4() == nil {
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
