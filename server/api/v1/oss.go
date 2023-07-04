package apiV1

import (
	"github.com/ChocolateAceCream/blog/global"
	"github.com/ChocolateAceCream/blog/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OssApi struct{}

// UploadFile
// @Tags      Oss
// @Summary   file uploader
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file                                                           true  "上传文件示例"
// @Success   200   {object}  response.Response{data=response.UploadFileResult,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /api/v1/oss/upload [post]
func (oa *OssApi) Uploader(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.LOGGER.Error("Fail to receive file", zap.Error(err))
		response.FailWithMessage("Fail to receive file", c)
		return
	}
	filePath, err := ossService.UploadFile(header) // 文件上传后拿到文件路径
	if err != nil {
		global.LOGGER.Error("fail to add file path to DB!", zap.Error(err))
		response.FailWithMessage("fail to add file path to DB", c)
		return
	}
	response.OkWithFullDetails(response.UploadFileResult{Url: filePath}, "success", c)
}
