package errno

// 通用：基本错误
// Code must start with 1xxxxx
const (
	// ErrUnknown - 500: Internal server error.
	ErrUnknown int = iota + 100001

	// ErrGenerateToken - 500: Error occurred while generating token.
	ErrGenerateToken

	// ErrParseToken - 500: Error occurred while parsing from token.
	ErrParseToken

	// ErrBindAndValidation - 400: Error occurred while binding the request body to the struct or validation failed.
	ErrBindAndValidation

	// ErrDatabase - 500: Database error.
	ErrDatabase

	// ErrRedis - 500: Redis error.
	ErrRedis

	// ErrPasswordInvalid - 401: Password invalid.
	ErrPasswordInvalid

	// ErrPasswordIncorrect - 401: Password incorrect.
	ErrPasswordIncorrect

	// ErrPasswordHash - 500: Error occurred while hashing password.
	ErrPasswordHash

	// ErrClientRPCInit - 500: RPC client initialization error.
	ErrClientRPCInit

	// ErrRPCLink - 500: RPC service link error.
	ErrRPCLink

	// ErrEncodingFailed - 500: Encoding failed.
	ErrEncodingFailed

	// ErrMinio - 500: Minio error.
	ErrMinio

	// ErrOpenFormFile - 500: Open request's form file error.
	ErrOpenFormFile
)
