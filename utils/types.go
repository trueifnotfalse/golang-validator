package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

const (
	integer         string = "^-?[0-9]+$"
	unsignedInteger string = "^[0-9]+$"
	float           string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	boolean         string = "^(true)|(false)$"
)

func IsString(v any) bool {
	_, ok := v.(string)
	return ok
}

func ToString(v any) string {
	str, ok := v.(string)
	if !ok {
		str = fmt.Sprintf("%v", v)
	}
	return str
}

func ToMap(v any) (map[string]any, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return nil, false
	}
	keys := rv.MapKeys()
	if len(keys) == 0 {
		return nil, false
	}
	result := make(map[string]any)
	for i := 0; i < len(keys); i++ {
		result[keys[i].String()] = rv.MapIndex(keys[i]).Interface()
	}

	return result, true
}

func ToFloat64Slice(v any) ([]float64, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Array && rv.Kind() != reflect.Slice {
		return nil, false
	}
	l := rv.Len()
	result := make([]float64, l)
	var (
		f  float64
		ok bool
	)
	for i := 0; i < l; i++ {
		v = rv.Index(i).Interface()
		f, ok = ToFloat(ToString(v))
		if !ok {
			return nil, false
		}
		result[i] = f
	}

	return result, true
}

func ToSlice(v any) ([]any, bool) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Array && rv.Kind() != reflect.Slice {
		return nil, false
	}
	l := rv.Len()
	result := make([]any, l)
	for i := 0; i < l; i++ {
		result[i] = rv.Index(i).Interface()
	}

	return result, true
}

func ToInt(v string) (int64, bool) {
	if !isType(v, integer) {
		return 0, false
	}
	i, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		return 0, false
	}
	return i, true
}

func ToUInt(v string) (uint64, bool) {
	if !isType(v, unsignedInteger) {
		return 0, false
	}
	i, err := strconv.ParseUint(v, 10, 0)
	if err != nil {
		return 0, false
	}
	return i, true
}

func ToFloat(v string) (float64, bool) {
	if !isType(v, float) {
		return 0, false
	}
	i, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, false
	}
	return i, true
}

func IsInt(v string) bool {
	return isType(v, integer)
}

func IsUInt(v string) bool {
	return isType(v, unsignedInteger)
}

func IsFloat(v string) bool {
	return isType(v, float)
}

func isType(v string, re string) bool {
	return regexp.MustCompile(re).MatchString(v)
}

func IsArray(v any) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Array || rv.Kind() == reflect.Slice {
		return true
	}

	return false
}

func IsMap(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Map
}

func IsBool(v any) bool {
	rv := reflect.ValueOf(v)
	return rv.Kind() == reflect.Bool
}

func Type(v any) reflect.Kind {
	s := ToString(v)
	if IsString(v) {
		return reflect.String
	} else if IsUInt(s) {
		return reflect.Uint
	} else if IsInt(s) {
		return reflect.Int
	} else if IsFloat(s) {
		return reflect.Float64
	}
	rv := reflect.ValueOf(v)

	return rv.Kind()
}
