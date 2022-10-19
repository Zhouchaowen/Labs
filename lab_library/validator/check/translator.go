package check

import (
	zhLocal "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
)

// 初始化翻译器
func TranslateInit(validate *validator.Validate) ut.Translator {
	zh_ch := zhLocal.New()
	uni := ut.New(zh_ch)                // 万能翻译器，保存所有的语言环境和翻译数据
	trans, _ := uni.GetTranslator("zh") // 翻译器
	_ = zh.RegisterDefaultTranslations(validate, trans)

	// 添加额外翻译
	_ = validate.RegisterTranslation("ulen", trans, func(ut ut.Translator) error {
		return ut.Add("ulen", "{0} 长度等于{1}!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("ulen", fe.Field(), fe.Param())
		return t
	})

	return trans
}
