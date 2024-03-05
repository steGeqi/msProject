package main

import (
	"github.com/gin-gonic/gin"
	srv "msProject.com/common"
	"msProject/webcenter/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, "web center", ":1015")

}
