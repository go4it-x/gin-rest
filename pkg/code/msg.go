package code

// message code Corresponding message
var message = map[uint32]string{
	Ok:           "SUCCESS",
	Fail:         "FAILED",
	Unauthorized: "未授权",
	NoRoute:      "404 NOT FOUND",

	// client
	ParamsIsInvalid:   "参数无效",
	LoginFailed:       "账号或密码错误",
	UserStatusDisable: "账号已禁用",

	// server
	ServerError:        "服务端错误",
	UserInfoFailed:     "获取用户失败",
	UserRegisterFailed: "用户注册失败",
}

// Value get msg by code
func Value(code uint32) string {
	if msg, ok := message[code]; ok {
		return msg
	} else {
		return ""
	}
}
