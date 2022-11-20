package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/response"
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
// @Router /api/v1/auth/captcha [post]
func (a AuthApi) GetCaptcha(c *gin.Context) {
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
