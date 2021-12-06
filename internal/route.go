package internal

import (
	_ "sports/docs" // 这里需要引入本地已生成文档
	accountCtr "sports/internal/account/controller"
	coachCtr "sports/internal/coach/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	Prefix = "/api/v1/" // route prefix
)

func RouteApi(g *gin.Engine) {
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	baseRoute := g.Group(Prefix)
	accountRoute := baseRoute.Group("/account/users").Use(accountCtr.Auth)
	{
		accountRoute.GET("", accountCtr.UserList)                 // 获取用户列表
		accountRoute.POST("", accountCtr.AddUser)                 // 用户注册, 添加用户
		accountRoute.PUT("/:user_id", accountCtr.UpdateUser)      // 更新用户信息
		accountRoute.PATCH("/:user_id", accountCtr.UpdateUserPWD) // 更新密码
		accountRoute.GET("/:user_id", accountCtr.GetUser)         // 获取用户详情
		accountRoute.DELETE("/:user_id", accountCtr.DeleteUser)   // 删除用户
	}
	// authorization management route
	authRoute := baseRoute.Group("/auth")
	{
		// internal login
		authRoute.GET("/current-user", accountCtr.GetUserByToken)  // 根剧 token 获取用户
		authRoute.GET("/users/:username/tokens", accountCtr.Login) // 用户登录
	}

	coachRoute := baseRoute.Group("/coach").Use(accountCtr.Auth)
	{
		coachRoute.GET("/training", coachCtr.GetLiveTraining)
	}

}
