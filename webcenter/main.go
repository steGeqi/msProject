package main

import (
	"github.com/gin-gonic/gin"
	srv "msProject.com/common"
	"msProject/webcenter/config"
	"msProject/webcenter/router"
)

func main() {
	r := gin.Default()
	config.AppConf.InitZaplog()
	router.InitRouter(r)
	srv.Run(r, "web center", ":80")

}
