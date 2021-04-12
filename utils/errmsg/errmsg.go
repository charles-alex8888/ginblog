package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
	// code 以1000.开头==》 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	// code 以2000.开头==》 文章模块的错误
	// code 以3000.开头==》 分类模块的错误
)

var codeMsg = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户以存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式不正确",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
