package controller

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"sports/config"
	"sports/internal/account/service"
	"sports/pkg/api/account"
	"sports/pkg/ctr"
	"sports/pkg/errs"
)

type JwtUserInfo struct {
	ID       int32  `json:"id" gorm:"size:32;not null;primary_key;unique_index"`
	UserName string `json:"username"` // 用户名称
}

type JwtClaims struct {
	User JwtUserInfo `json:",inline"`
	jwt.StandardClaims
}

// 用户登录成功后调取此方法获取 Token
func SignUserToken(user int32, secret string) string {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtClaims{User: JwtUserInfo{ID: user}})
	k, _ := t.SignedString([]byte(secret))
	return k
}

func authRequest(req *http.Request) (int32, error) {
	token := ""
	if token = req.Header.Get("Authorization"); token == "" {
		return 0, errs.Unauthorized
	}
	user := JwtClaims{}
	_, err := jwt.ParseWithClaims(token, &user, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().Secret), nil
	})
	if err != nil {
		return 0, err
	}
	return user.User.ID, nil
}

func Auth(ctx *gin.Context) {
	var user int32
	if u, err := authRequest(ctx.Request); err != nil {
		ctr.Err(ctx, errs.Unauthorized)
		ctx.Abort()
		return
	} else {
		user = u
	}
	ctx.Set("user", user)
	ctx.Next()
}

// @summary 登录
// @Description 登录，输入用户 ID，用户密码获取授权 Token
// @Param account.LoginInfo	query	account.LoginInfo true  "用户登录信息"
// @Success 200 {object} account.AuthInfo
// @Failure 500 {object} errs.BasicError
// @Router  /api/v1/account/users [post]
func Login(ctx *gin.Context) {
	var loginInfo account.LoginInfo
	if err := ctx.ShouldBind(&loginInfo); err != nil {
		ctr.Err(ctx, errs.ErrorParams.Params(err))
		return
	}
	// check data
	if loginInfo.ID == 0 {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("user_id is empty")))
		return
	}
	if loginInfo.Password == "" {
		ctr.Err(ctx, errs.ErrorParams.Params(fmt.Errorf("password is empty")))
		return
	}

	result, err := service.Login(&loginInfo)
	if err != nil {
		ctr.Err(ctx, err)
		return
	}

	// get token
	token := SignUserToken(loginInfo.ID, config.GetConfig().Secret)
	result.Token = token

	ctr.Success(ctx, result)
}
