package errs

import "net/http"

const SystemCodePrefix Prefix = "E_SYSTEM"

var (
	ErrorSystem = SystemCodePrefix.Code("SYSTEM_ERROR").
		Message("system error: %v").
		HTTPCode(http.StatusInternalServerError) // 系统异常
	ErrorParams = SystemCodePrefix.Code("PARAM_PARSE_ERROR").
		Message("params parse error : %v").
		HTTPCode(http.StatusBadRequest) // 参数解析错误
	Unauthorized = SystemCodePrefix.Code("UNAUTHORIZED").
		Message("unauthorized request").
		HTTPCode(http.StatusUnauthorized) // 未认证
	PermissionDenied = SystemCodePrefix.Code("PERMISSION_DENIED").
		Message("can not access resources").
		HTTPCode(http.StatusForbidden) // 无权限访问
)
