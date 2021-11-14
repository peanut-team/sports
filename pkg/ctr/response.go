package ctr

import (
	"net/http"
	"sports/pkg/errs"
	"sports/pkg/logger"
	"sports/pkg/page"

	"github.com/gin-gonic/gin"
)

// Success set code 200 and return object in JSON format
func Success(ctx *gin.Context, object interface{}) {
	SuccessWithCodeObject(ctx, http.StatusOK, object)
}

// Success set code 200 and return object of list
func SuccessList(ctx *gin.Context, object interface{}, page page.Paginate) {
	ctx.JSON(http.StatusOK, gin.H{
		"items":    object,
		"paginate": page,
	})
}

// return success and with block content
func SuccessBlock(ctx *gin.Context) {
	SuccessWithCodeObject(ctx, http.StatusNoContent, nil)
}

// common return
func SuccessWithCodeObject(ctx *gin.Context, code int, object interface{}) {
	ctx.JSON(code, object)
}

// 处理错误请求
func Err(ctx *gin.Context, err error) {
	logger.Errorf("url: %s error: %s", ctx.Request.URL.String(), err.Error())
	if ne, ok := err.(*errs.BasicError); ok {
		ne.Parse()
		if ne.HTTPStatusCode == 0 {
			ctx.JSON(http.StatusBadRequest, ne)
		} else {
			ctx.JSON(ne.HTTPStatusCode, ne)
		}
	} else {
		ctx.JSON(http.StatusInternalServerError, err)
	}
}
