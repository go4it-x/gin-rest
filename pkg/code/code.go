package code

// Ok 成功
const Ok uint32 = 0

// Fail 失败
const Fail uint32 = 400

// Unauthorized 未授权
const Unauthorized uint32 = 401

// NoRoute 理由错误
const NoRoute uint32 = 404

//==================客户端 400xxx==================
// ParamsIsInvalid 参数无效
const ParamsIsInvalid uint32 = 400001

// LoginFailed 登录失败
const LoginFailed uint32 = 400002

// UserStatusDisable 用户状态不可用
const UserStatusDisable uint32 = 400003

//==================服务端 500xxx==================
// ServerError 服务端错误
const ServerError uint32 = 500001

// UserInfoFailed 用户信息错误
const UserInfoFailed uint32 = 500002

// UserRegisterFailed 用户注册失败
const UserRegisterFailed uint32 = 500003
