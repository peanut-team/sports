package account

type User struct {
	Username string `json:"username" example:"mick"`                   // 用户名
	Email    string `json:"email" example:"123@ee.com"`                // 邮箱
	Password string `json:"password,omitempty" example:"123@password"` // 密码
}
