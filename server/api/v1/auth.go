package apiV1

import (
	"fmt"
	"image/color"
	"time"

	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/request"
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
	var driver base64Captcha.Driver
	if config.DigitsOnly == true {
		driver = base64Captcha.NewDriverDigit(config.Height, config.Width, config.Length, config.MaxSkew, config.DotCount)
	} else {
		driverString := base64Captcha.DriverString{
			Height:          config.Height,
			Width:           config.Width,
			NoiseCount:      0,
			ShowLineOptions: 2,
			Length:          config.Length,
			Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
			BgColor: &color.RGBA{
				R: 3,
				G: 102,
				B: 214,
				A: 125,
			},
			Fonts: []string{"wqy-microhei.ttc"},
		}
		driver = driverString.ConvertFonts()
	}
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
// @Param data body request.SendEmail true "email"
// @Success 200 {object} response.Response{msg=string} "return send email result"
// @Router /api/public/auth/sendEmailCode [post]
func (a *AuthApi) SendEmailCode(c *gin.Context) {
	var req request.SendEmail
	if err := c.ShouldBindJSON(&req); err != nil {
		global.LOGGER.Error("SendEmailCode error: ", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	currentSession := utils.GetSession(c)
	code := utils.RandomString(global.CONFIG.Captcha.Length)
	fmt.Println("------email code-------", code)
	global.REDIS.Set(c, global.CONFIG.Email.Prefix+currentSession.UUID, code, time.Duration(global.CONFIG.Email.Expiration)*time.Second)
	body := fmt.Sprintf("verification code is <b>%v</b>, expired in 2 minutes", code)
	if err := utils.SendMail(req.Email, "Verification Code", body); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Code has been sent", c)
}
