package error

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "页面未发现")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	UserNotExist              = NewError(1001, "用户名不存在")
	UserExist                 = NewError(1002, "用户名已存在")
	ErrorPassword             = NewError(1003, "登陆密码错误")
	DirError                  = NewError(1004, "failed to create save directory")
	DirPermissionError        = NewError(1005, "insufficient file permissions")
)
