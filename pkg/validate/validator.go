package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

var (
	v *validator.Validate
)

func init() {
	v = validator.New()
}

// Do validate params
// params:Parameters to be verified
// excludes:Parameter field name without validation
func Do(params interface{}, excludes ...string) (msg string) {
	if err := v.Struct(params); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fieldName := strings.ToLower(e.Field())
			for _, exclude := range excludes {
				if fieldName == strings.ToLower(exclude) {
					continue
				}
			}

			switch e.Tag() {
			case "required":
				msg = fmt.Sprintf("%v is required", e.Field())
			case "email":
				msg = fmt.Sprintf("%v must be a mail", e.Field())
			default:
				msg = err.Error()
			}

			return msg
		}
	}

	return
}
