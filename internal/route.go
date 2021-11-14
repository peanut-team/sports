package internal

import (
	_ "sports/docs" // 这里需要引入本地已生成文档
	accountCtr "sports/internal/account/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	Prefix = "/api/mc/v1/" // route prefix
)

func RouteApi(g *gin.Engine) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRoute := g.Group(Prefix)
	accountRoute := baseRoute.Group("/account")
	{
		accountRoute.GET("/:username", accountCtr.GetUser) // 用户列表
	}
}
