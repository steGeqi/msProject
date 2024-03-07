package login

import (
	"github.com/gin-gonic/gin"
	"msProject.com/common"
)

type HandleLogin struct {
}

func (HandleLogin) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	ctx.JSON(200, result.Success("123456"))
}
