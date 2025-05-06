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
	i, err := strconv.ParseFloat(v, 10)
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
	switch rv.Kind() {
	case reflect.Array, reflect.Slice:
		return true
	}
	return false
}

func IsMap(v any) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		return true
	}
	return false
}

func IsBool(v any) bool {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Bool:
		return true
	}
	return false
}

func Type(v any) reflect.Kind {
	s := ToString(v)
	if IsString(v) {
		return reflect.String
	}  else if IsUInt(s) {
		return reflect.Uint
	} else if IsInt(s) {
		return reflect.Int
	} else if IsFloat(s) {
		return reflect.Float64
	}
	rv := reflect.ValueOf(v)

	return rv.Kind()
}
