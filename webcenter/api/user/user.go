package user

import "github.com/gin-gonic/gin"

type HandleUser struct {
}

func (*HandleUser) getCaptcha(ctx *gin.Context) {
	ctx.JSON(200, "DFJH")
}
