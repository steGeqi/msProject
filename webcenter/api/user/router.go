package user

import "github.com/gin-gonic/gin"

type RouterUser struct {
}

func (u *RouterUser) Register(r *gin.Engine) {
	h := HandleUser{}
	r.GET("project/login/getCaptcha", h.getCaptcha)
}
