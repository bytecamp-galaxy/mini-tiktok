package errno

// 下列错误码，除了提供枚举值之外，无其他目的

const ErrSuccess int32 = 0

const (
	ErrStatusNotFound int = iota + 100101
	ErrStatusMethodNotAllowed
	ErrStatusInternalServerError
	ErrUnauthorized
)
