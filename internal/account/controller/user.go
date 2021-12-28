package controller

import (
	"fmt"
	"strconv"

	"sports/internal/account/service"
	"sports/pkg/api/account"
	"sports/pkg/ctr"
	"sports/pkg/errs"
	"sports/pkg/page"

	"github.com/gin-gonic/gin"
)

// @summary 获取用户信息
// @Description 传入 username，根据用户名查询用户
// @Param user_id	path	int	true	"用户ID"
// @Success 200 {object} account.User
// @Failure 500 {object} errs.BasicError
// @Router  /api/mc/v1/accounts/:user_id [get]
func GetUser(c *gin.Context) {
	idStr := c.Param("user_id")
	// check data
	if idStr == "" {
		ctr.Err(c, errs.ErrorParams.Params(fmt.Errorf("user_id is empty")))
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctr.Err(c, errs.ErrorParams.Params(fmt.Errorf("invalid id %s", idStr)))
		return
	}

	result, err := service.GetUser(int(id))
	if err != nil {
		ctr.Err(c, err)
		return
	}
	ctr.Success(c, result)
}

type UserListResp struct {
	Items    []*account.UserItem `json:"items"`    // 对象列表
	Paginate page.Paginate       `json:"paginate"` // Page
}

// @summary 获取用户列表
// @Description 传入 page 参数，查询用户列表
// @Param page.Paginate	query	page.Paginate true  "分页参数"
// @Success 200 {object} UserListResp
// @Failure 500 {object} errs.BasicError
// @Router   /api/v1/account/users [get]
func UserList(ctx *gin.Context) {
	var query page.Paginate
	if err := ctx.BindQuery(&query); err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(err))
		return
	}

	result, p, err := service.UserList(query)
	if err != nil {
		ctr.Err(ctx, err)
		return
	}

	ctr.SuccessList(ctx, result, p)
}

// @summary 添加用户
// @Description 添加用户
// @Param account.User	body	account.User true  "用户信息"
// @Success 200 {object} account.User
// @Failure 500 {object} errs.BasicError
// @Router  /api/v1/account/users [post]
func AddUser(ctx *gin.Context) {
	var query account.User
	if err := ctx.ShouldBind(&query); err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(err))
		return
	}
	// check data
	if query.Username == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("username cannot be empty")))
		return
	}
	if query.Password == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("password cannot be empty")))
		return
	}

	result, err := service.AddUser(&query)
	if err != nil {
		ctr.Err(ctx, err)
		return
	}
	ctr.Success(ctx, result)
}

type UpdateUserReq struct {
	ID       int32  `json:"id,omitempty"`               // 用户 ID
	Username string `json:"username" example:"mick"`    // 用户名
	Email    string `json:"email" example:"123@ee.com"` // 邮箱
}

// @summary 更新用户信息
// @Description 更新用户信息
// @Param user_id	path	int	true	"用户ID"
// @Param UpdateUserReq	body	UpdateUserReq true  "用户信息"
// @Success 200 {object} account.User
// @Failure 500 {object} errs.BasicError
// @Router  /api/v1/account/users/:user_id [put]
func UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
	// check data
	if idStr == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("user_id is empty")))
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("invalid user_id %s", idStr)))
		return
	}

	var query account.User
	if err := ctx.ShouldBind(&query); err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(err))
		return
	}

	query.ID = int32(id)
	// check data
	if query.Username == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("username cannot be empty")))
		return
	}

	result, err := service.UpdateUser(&query, service.Details_UpdateUserOption)
	if err != nil {
		ctr.Err(ctx, err)
		return
	}
	ctr.Success(ctx, result)
}

type UpdateUserPWDReq struct {
	ID       int32  `json:"id,omitempty"`                              // 用户 ID
	Password string `json:"password,omitempty" example:"123@password"` // 密码
}

// @summary 更新用户密码
// @Description 更新用户密码
// @Param user_id	path	int	true	"用户ID"
// @Param UpdateUserPWDReq	body	UpdateUserPWDReq true  "用户信息"
// @Success 200 {object} account.User
// @Failure 500 {object} errs.BasicError
// @Router  /api/v1/account/users/:user_id [PATCH]
func UpdateUserPWD(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
	// check data
	if idStr == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("user_id is empty")))
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("invalid user_id %s", idStr)))
		return
	}

	var query account.User
	if err := ctx.ShouldBind(&query); err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(err))
		return
	}

	query.ID = int32(id)
	// check data
	if query.Username == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("username cannot be empty")))
		return
	}
	if query.Password == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("password cannot be empty")))
		return
	}

	_, err = service.UpdateUser(&query, service.PWD_UpdateUserOption)
	if err != nil {
		ctr.Err(ctx, err)
		return
	}
	ctr.SuccessBlock(ctx)
}

// @summary 删除用户
// @Description 删除用户
// @Param user_id	path	int	true	"用户ID"
// @Success 200
// @Failure 500 {object} errs.BasicError
// @Router  /api/v1/account/users/:user_id [delete]
func DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("user_id")
	// check data
	if idStr == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("user_id is empty")))
		return
	}
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("invalid user_id %s", idStr)))
		return
	}

	err = service.DeleteUser(int32(id))
	if err != nil {
		ctr.Err(ctx, err)
		return
	}
	ctr.SuccessBlock(ctx)
}

// @summary 根据 HTTP Header Token 获取当前用户信息
// @Description 根据 HTTP Header Token 获取当前用户信息，需要在头部添加 Authorization
// @Param Authorization	header	string	true	"授权 token"
// @Success 200 {object} account.User
// @Failure 500 {object} errs.BasicError
// @Router  /api/v1/auth/current-user [get]
func GetUserByToken(ctx *gin.Context) {
	var user int32
	if u, err := authRequest(ctx.Request); err != nil {
		ctr.Err(ctx, errs.Unauthorized)
		ctx.Abort()
		return
	} else {
		user = u
	}

	result, err := service.GetUser(int(user))
	if err != nil {
		ctr.Err(ctx, err)
		return
	}
	ctr.Success(ctx, result)
}
