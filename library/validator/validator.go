package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

//InitTrans 初始化校验器
func InitTrans(locale, label string) (err error) {
	//修改gin框架的validator引擎，实现自定义
	if ev, ok := binding.Validator.Engine().(*v.Validate); ok {
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		//arg1:备用的语言环境 args...需要支持的语言环境
		uni := ut.New(enT, zhT, enT)

		var ok bool
		//返回给定语言环境的默认指定翻译器
		Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		//获取struct{}中label的字段作开头的提示信息 没有label标签则默认结构体字段名
		ev.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get(label)
			return name
		})

		//注册翻译器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(ev, Trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(ev, Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(ev, Trans)
		}
		return
	}
	return
}

//GetErrMsg 获取验证器中的错误信息
func GetErrMsg(errMap v.ValidationErrorsTranslations) string {
	var errMsg []string
	for _, msg := range errMap {
		errMsg = append(errMsg, msg)
	}
	return strings.Join(errMsg, ";")
}
