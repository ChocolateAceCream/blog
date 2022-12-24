package apiV1

import (
	"fmt"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/dbTable"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/ChocolateAceCream/blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type AuthApi struct{}

// Captcha
// @Tags Auth
// @Summary Generate captcha image
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.CaptchaResponse, msg=string} "return base64 captcha image"
// @Router /api/public/auth/captcha [post]
func (a *AuthApi) GetCaptcha(c *gin.Context) {
	config := global.CONFIG.Captcha
	driver := base64Captcha.NewDriverDigit(config.Height, config.Width, config.Length, config.MaxSkew, config.DotCount)
	cp := base64Captcha.NewCaptcha(driver, store.AttachContext(c))
	if _, b64s, err := cp.Generate(); err != nil {
		global.LOGGER.Error("failed to fetch Captcha!", zap.Error(err))
		response.FailWithMessage("failed to fetch Captcha", c)
	} else {
		response.OkWithFullDetails(response.CaptchaResponse{
			Captcha: b64s,
		}, "Success", c)
	}
}

// Send verification code email
// @Tags Auth
// @Summary Send verification code email
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "return send email result"
// @Router /api/public/auth/sendEmailCode [post]
func (a *AuthApi) SendEmailCode(c *gin.Context) {
	currentUser, err := utils.GetValueFromSession[dbTable.User](c, "currentUser")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	code := utils.RandomString(global.CONFIG.Captcha.Length)
	global.REDIS.Set(c, global.CONFIG.Email.Prefix+currentUser.UUID.String(), code, time.Duration(global.CONFIG.Email.Expiration)*time.Second)
	body := fmt.Sprintf("verification code is <b>%v</b>", code)
	if err := utils.SendMail(currentUser.Email, "Verification Code", body); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Code has been sent", c)
}
