package login

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"msProject.com/common"
	"msProject/webcenter/pkg/dao"
	"msProject/webcenter/pkg/model"
	"msProject/webcenter/pkg/repo"
	"time"
)

type HandleLogin struct {
	cache repo.Cache
}

func New() *HandleLogin {
	return &HandleLogin{
		cache: dao.Rc,
	}
}

// GetCaptcha获取手机验证码

func (hl *HandleLogin) GetCaptcha(ctx *gin.Context) {
	result := &common.Result{}
	// 获取参数
	mobile := ctx.PostForm("mobile")
	if !common.VerifyMobile(mobile) {
		ctx.JSON(200, result.Fail(model.LoginMobileNotLegal, "不合格"))
		return
	}
	// 生成验证码
	code := "12345"
	go func() {
		time.Sleep(2 * time.Second)
		log.Println("调用短信平台发送短信")
		zap.L().Info("短信平台调用成功 info")
		zap.L().Error("短信平台调用成功 error")
		zap.L().Debug("短信平台调用,debug")
		c, concel := context.WithTimeout(context.Background(), 2*time.Second)
		defer concel()
		err := hl.cache.Put(c, "REGISTER_"+mobile, code, 15*time.Minute)
		if err != nil {
			log.Println("验证码存入redis错误,cause by:", err)
		}
		fmt.Println(mobile, code)
	}()
	ctx.JSON(200, result.Success("123456"))
}
