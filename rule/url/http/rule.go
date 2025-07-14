package http

import (
	"fmt"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
	"net/url"
)

type Rule struct {
	message string
}

func New() rule.Interface {
	return &Rule{
		message: "The %s format is not HTTP URL.",
	}
}

func (r *Rule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	u, err := url.Parse(utils.ToString(v))
	if err != nil || u.Host == "" || (u.Scheme != "http" && u.Scheme != "https") {
		return fmt.Errorf(r.message, key)
	}

	return nil
}
