package validator

import (
	"bytes"
	"github.com/goccy/go-json"
)

func Validate(body []byte, rules Rules) Errors {
	if len(body) == 0 {
		return nil
	}
	t, err := decodeBody(body)
	if err != nil {
		return Errors{"0": {err}}
	}

	return Map(t, rules)
}

func Map(values map[string]any, rules Rules) Errors {
	var (
		err error
		i   int
	)
	result := make(Errors)
	for k, rl := range rules {
		for i = 0; i < len(rl); i++ {
			err = rl[i].Valid(k, values)
			if err != nil {
				result[k] = append(result[k], err)
			}
		}
	}

	return result
}

func decodeBody(body []byte) (map[string]any, error) {
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
