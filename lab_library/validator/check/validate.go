package check

import (
	"errors"
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Check struct {
	Validate  *validator.Validate
	Translate ut.Translator
}

func NewCheck() *Check {
	va := validator.New()
	// 自定义tag注册
	RegisterAll(va)
	// 注册万能翻译
	trans := TranslateInit(va)
	return &Check{
		Validate:  va,
		Translate: trans,
	}
}

func (c *Check) Var(field interface{}, tag string) error {
	err := c.Validate.Var(field, tag)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return errors.New(fmt.Sprintf("%v", validationErrors.Translate(c.Translate)))
		}
	}
	return nil
}

func (c *Check) Struct(s interface{}) error {
	err := c.Validate.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return errors.New(fmt.Sprintf("%v", validationErrors.Translate(c.Translate)))
		} else {
			return err
		}
	}
	return nil
}

func (c *Check) ValidateMap(data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	err := c.Validate.ValidateMap(data, rules)
	return err
}

func (c *Check) RegisterAlias(alias, tags string) {
	c.Validate.RegisterAlias(alias, tags)
	return
}
