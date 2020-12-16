package param

// SayTest parameter
type SayTest struct {
	Username string `form:"username" validate:"required" desc:"用户名"` // 用户名
}
