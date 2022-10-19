package check

import (
	"github.com/go-playground/validator/v10"
	"strconv"
	"unicode/utf8"
)

// 注册自定义tag校验
func RegisterAll(valid *validator.Validate) {
	_ = valid.RegisterValidation("ulen", ulen)
	_ = valid.RegisterValidation("ugt", ugt)
	_ = valid.RegisterValidation("ult", ult)
}

// 汉字长度校验
func ulen(fl validator.FieldLevel) bool {
	length := utf8.RuneCountInString(fl.Field().String())
	param, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}
	if length == param {
		return true
	}
	return false
}

// 汉字长度大于
func ugt(fl validator.FieldLevel) bool {
	length := utf8.RuneCountInString(fl.Field().String())
	param, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}
	if length > param {
		return true
	}
	return false
}

// 汉字长度小于
func ult(fl validator.FieldLevel) bool {
	length := utf8.RuneCountInString(fl.Field().String())
	param, err := strconv.Atoi(fl.Param())
	if err != nil {
		return false
	}
	if length < param {
		return true
	}
	return false
}
