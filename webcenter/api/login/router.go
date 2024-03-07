package login

import "github.com/gin-gonic/gin"

type RouterLogin struct {
}

func (*RouterLogin) Register(r *gin.Engine) {
	g := r.Group("/project/login")
	h := New()
	g.POST("/getCaptcha", h.GetCaptcha)
}
