package routes

import (
	"github.com/gin-gonic/gin"
	"lz-web-serviece/middleware"
	"lz-web-serviece/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Cors())
	r.Use(gin.Recovery())

	// 部署
	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static", "static/admin/static")
	r.Static("admin/favicon.ico", "static/admin/favicon.ico")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	_ = r.Run(utils.HttpPort)
}