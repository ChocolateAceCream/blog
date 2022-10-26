package utils

import (
	"regexp"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var validatorMapper map[string]func(fl validator.FieldLevel) bool

var (
	//RegIDcheck 检查身份证
	RegIDcheck = regexp.MustCompile(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X)$)`)
	//RegHTTPCheck 检查HTTP格式
	RegHTTPCheck = regexp.MustCompile(`^((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`)
	//RegPhoneCheck 检查电话格式
	RegPhoneCheck = regexp.MustCompile(`1[345678]\d{9}`)
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for validatorName, validatorFunction := range validatorMapper {
			err := v.RegisterValidation(validatorName, validatorFunction)
			if err != nil {
				global.LOGGER.Info("validator register success")
			}
		}

	}
}

func idcheck(fl validator.FieldLevel) bool {
	return RegIDcheck.MatchString(fl.Field().String())
}

func httpcheck(fl validator.FieldLevel) bool {
	return RegHTTPCheck.MatchString(fl.Field().String())
}

func phonecheck(fl validator.FieldLevel) bool {
	return RegPhoneCheck.MatchString(fl.Field().String())
}
