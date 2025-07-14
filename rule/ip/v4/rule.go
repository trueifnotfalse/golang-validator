package v4

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"net"
)

type Rule struct {
	message string
}

func New() rule.Interface {
	return &Rule{
		message: "The %s must be a valid IPv4 address.",
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	ip := net.ParseIP(utils.ToString(v))
	if ip == nil || ip.To4() == nil {
		return fmt.Errorf(r.message, key)
	}

	return nil
}
