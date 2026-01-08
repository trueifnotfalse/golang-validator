package validator

import (
	"bytes"
	"github.com/goccy/go-json"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/locale/en"
)

type Validator struct {
	loc locale.Interface
}

func New() *Validator {
	return &Validator{
		loc: getDefaultLocale(),
	}
}

func (r *Validator) SetLocale(v locale.Interface) *Validator {
	r.loc = v
	return r
}

func (r *Validator) Validate(body []byte, rules Rules) Errors {
	var (
		t   map[string]any
		err error
	)
	if len(body) == 0 {
		return r.Map(t, rules)
	}
	t, err = r.decodeBody(body)
	if err != nil {
		return Errors{"0": {err}}
	}

	return r.Map(t, rules)
}

func (r *Validator) Map(values map[string]any, rules Rules) Errors {
	var (
		err error
		i   int
	)
	l := r.getLocale()
	result := make(Errors)
	for k, rl := range rules {
		for i = 0; i < len(rl); i++ {
			err = rl[i].SetLocale(l).Valid(k, values)
			if err != nil {
				result[k] = append(result[k], err)
			}
		}
	}

	return result
}

func (r *Validator) decodeBody(body []byte) (map[string]any, error) {
	var t map[string]any
	reader := bytes.NewReader(body)
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	err := decoder.Decode(&t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *Validator) getLocale() locale.Interface {
	if r.loc != nil {
		return r.loc
	}
	return getDefaultLocale()
}

func getDefaultLocale() locale.Interface {
	return en.New()
}
