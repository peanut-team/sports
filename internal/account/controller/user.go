package controller

import (
	"fmt"
	"sports/internal/account/service"
	"sports/pkg/ctr"
	"sports/pkg/errs"

	"github.com/gin-gonic/gin"
)

// @summary 获取用户信息
// @Description 传入 username，根据用户名查询用户
// @Param username	path	string	true	"nick"
// @Success 200 {object} account.User
// @Failure 500 {object} errs.BasicError
// @Router  /api/mc/v1/account/:username [get]
func GetUser(c *gin.Context) {
	username := c.Param("username")
	// check data
	if username == "" {
		ctr.Err(c, errs.ErrorParams.Params(fmt.Errorf("username is empty")))
		return
	}

	result, err := service.GetUser(username)
	if err != nil {
		ctr.Err(c, err)
		return
	}
	ctr.Success(c, result)
}
