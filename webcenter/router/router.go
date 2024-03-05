package router

import (
	"github.com/gin-gonic/gin"
	"msProject/webcenter/api/user"
)

type Router interface {
	Register(r *gin.Engine)
}
type RegisterRouter struct {
}

func New() RegisterRouter {
	return RegisterRouter{}
}

func (RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Register(r)

}

func InitRouter(r *gin.Engine) {
	router := New()
	router.Route(&user.RouterUser{}, r)
}
