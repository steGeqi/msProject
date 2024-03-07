package login

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"msProject.com/common"
	"msProject/webcenter/pkg/model"
	"time"
)

type HandleLogin struct {
}

func (HandleLogin) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	// 获取参数
	mobile := ctx.PostForm("mobile")
	if !common.VerifyMobile(mobile) {
		ctx.JSON(200, result.Fail(model.LoginMobileNotLegal, "不合格"))
		return
	}
	// 生成验证码
	code := 12345
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("调用短信平台发送短信")
		fmt.Println(mobile, code)
	}()
	ctx.JSON(200, result.Success("123456"))
}
