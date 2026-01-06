# golang-validator

Validate golang request data with simple rules. Inspired by Laravel's request validation.

### validating a simple value

```go
package main

import (
	"fmt"

	validator "github.com/trueifnotfalse/golang-validator"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
)

func main() {
	data := `{
		"name": "John Doe",
		"age" : 104
	}`

	rules := map[string][]rule.Interface{
		"name": {
			validator.Required(), // not nil or empty
			validator.String(),   // string
			validator.Max(255),   // 255 characters length
		},
		"age": {validator.UInt8(), validator.Min(0), validator.Max(100)},
	}

	v := validator.New()
	errs := v.Validate([]byte(data), rules)
	for field, err := range errs {
		fmt.Printf("%s - %v\n", field, err)
	}
}
```

### creating own rule

```go
package main

import (
	"errors"
	"fmt"
	"strings"

	validator "github.com/trueifnotfalse/golang-validator"
	"github.com/trueifnotfalse/golang-validator/interface/locale"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/utils"
)

type MyRule struct{}

func (r *MyRule) SetLocale(v locale.Interface) rule.Interface {
	return r
}

func (r *MyRule) Valid(key string, values map[string]any) error {
	v, ok := values[key]
	if !ok {
		return nil
	}
	if !utils.IsString(v) {
		return nil
	}
	s := utils.ToString(v)
	if s == strings.ToTitle(s) {
		return nil
	}
	return errors.New("All words must be capitalized.")
}

func main() {
	data := `{
		"name": "john Doe",
		"age" : 104
	}`
	r := &MyRule{}
	rules := map[string][]rule.Interface{
		"name": {
			validator.Required(), // not nil or empty
			validator.String(),   // string
			validator.Max(255),   // 255 characters length
			r,
		},
		"age": {validator.UInt8(), validator.Min(0), validator.Max(100)},
	}

	v := validator.New()
	errs := v.Validate([]byte(data), rules)
	for field, err := range errs {
		fmt.Printf("%s - %v\n", field, err)
	}
}
```

### creating own locale

```go
package main

import (
	"fmt"

	validator "github.com/trueifnotfalse/golang-validator"
	"github.com/trueifnotfalse/golang-validator/interface/rule"
	"github.com/trueifnotfalse/golang-validator/locale"
)

func main() {
	data := `{
		"name": "John Doe",
		"age" : 104
	}`
	
	myLocale := locale.Locale{
		"max.numeric": "The age in field `%s` is to high must be not greater than %d.",
	}

	rules := map[string][]rule.Interface{
		"name": {
			validator.Required(), // not nil or empty
			validator.String(),   // string
			validator.Max(255),   // 255 characters length
		},
		"age": {validator.UInt8(), validator.Min(0), validator.Max(100)},
	}

	v := validator.New().SetLocale(myLocale)
	errs := v.Validate([]byte(data), rules)
	for field, err := range errs {
		fmt.Printf("%s - %v\n", field, err)
	}
}
```