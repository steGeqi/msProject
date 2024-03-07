package main

import (
	"github.com/gin-gonic/gin"
	"log"
	srv "msProject.com/common"
	"msProject.com/common/logs"
	"msProject/webcenter/router"
)

func main() {
	r := gin.Default()
	//从配置中读取日志配置，初始化日志
	lc := &logs.LogConfig{
		DebugFileName: "D:\\msproject\\project_ws\\logs\\debug\\project-debug.log",
		InfoFileName:  "D:\\msproject\\project_ws\\logs\\info\\project-info.log",
		WarnFileName:  "D:\\msproject\\project_ws\\logs\\error\\project-error.log",
		MaxSize:       500,
		MaxAge:        28,
		MaxBackups:    3,
	}
	err := logs.InitLogger(lc)
	if err != nil {
		log.Fatalln(err)
	}
	router.InitRouter(r)
	srv.Run(r, "web center", ":80")

}
