package validate

import (
	"github.com/douyu/jupiter/pkg/xlog"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translationsZh "github.com/go-playground/validator/v10/translations/zh"
)

var (
	v     *validator.Validate
	trans ut.Translator
)

func init() {
	v = validator.New()
	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
	err := translationsZh.RegisterDefaultTranslations(v, trans)
	if err != nil {
		xlog.Error(err.Error())
	}
}

// Do validate params
// params:Parameters to be verified
func Do(params interface{}) (msg string) {
	if err := v.Struct(params); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return e.Translate(trans)
		}
	}

	return
}
