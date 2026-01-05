package rule

import "github.com/trueifnotfalse/golang-validator/interface/locale"

type Interface interface {
	Valid(key string, values map[string]any) error
	SetLocale(v locale.Interface) Interface
}
