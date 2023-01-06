package utils

import (
	"regexp"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/dlclark/regexp2"
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
	validatorMapper := map[string]func(fl validator.FieldLevel) bool{
		"idCheck":       idCheck,
		"httpCheck":     httpCheck,
		"phoneCheck":    phoneCheck,
		"passwordCheck": passwordCheck,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for validatorName, validatorFunction := range validatorMapper {
			err := v.RegisterValidation(validatorName, validatorFunction)
			if err != nil {
				global.LOGGER.Info("validator register success")
			}
		}

	}
}

func idCheck(fl validator.FieldLevel) bool {
	return RegIDcheck.MatchString(fl.Field().String())
}

func httpCheck(fl validator.FieldLevel) bool {
	return RegHTTPCheck.MatchString(fl.Field().String())
}

func phoneCheck(fl validator.FieldLevel) bool {
	return RegPhoneCheck.MatchString(fl.Field().String())
}

func passwordCheck(fl validator.FieldLevel) bool {
	expr := `^(?![0-9a-zA-Z]+$)(?![a-zA-Z!@#$%^&*]+$)(?![0-9!@#$%^&*]+$)[0-9A-Za-z!@#$%^&*]{8,16}$`
	reg, _ := regexp2.Compile(expr, 0)
	m, _ := reg.FindStringMatch(fl.Field().String())
	return m != nil
}
