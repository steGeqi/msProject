package main

import (
	"github.com/gin-gonic/gin"
	srv "msProject.com/common"
)

func main() {
	r := gin.Default()
	srv.Run(r, "web center", ":1015")
}
