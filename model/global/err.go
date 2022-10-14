package global

var (
	ErrorInternal         = NewError(1000, "内部错误", nil)
	ErrUserNotExist       = NewError(1001, "用户不存在", nil)
	ErrPasswordWrong      = NewError(1002, "密码错误", nil)
	ErrUserExist          = NewError(1003, "用户已存在", nil)
	ErrDataBind           = NewError(1004, "数据解析绑定失败", nil)
	ErrUsernameOrPassword = NewError(1005, "用户名或密码错误", nil)
	ErrUserRegister       = NewError(1006, "用户注册失败", nil)
	ErrSessionTimeout     = NewError(1007, "session超时", nil)
	ErrSessionNotExist    = NewError(1008, "session不存在", nil)
)

type CustomError struct {
	Code    int
	Message string
	Data    any
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewError(code int, message string, data any) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func NewCustomError(e *CustomError, data any) *CustomError {
	return &CustomError{
		Code:    e.Code,
		Message: e.Message,
		Data:    data,
	}
}
