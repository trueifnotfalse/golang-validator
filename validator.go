package validator

import (
	"bytes"
	"encoding/json"
)

func Validate(body []byte, rules Rules) map[string][]error {
	t, err := decodeBody(body)
	if err != nil {
		return map[string][]error{"0": {err}}
	}

	return Map(t, rules)
}

func Map(values map[string]any, rules Rules) map[string][]error {
	var (
		err error
		i   int
	)
	result := make(map[string][]error)
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
