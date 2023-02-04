package errno

// 通用：基本错误
// Code must start with 1xxxxx
const (
	// ErrUnknown - 500: Internal server error.
	ErrUnknown int = iota + 100001

	// ErrTokenInvalid - 401: Token invalid.
	ErrTokenInvalid

	// ErrTokenGeneration - 500: Error occurred while generating token.
	ErrTokenGeneration

	// ErrParseToken - 500: Error occurred while parsing from token.
	ErrParseToken

	// ErrBindAndValidation - 400: Error occurred while binding the request body to the struct or validation failed.
	ErrBindAndValidation

	// ErrDatabase - 500: Database error.
	ErrDatabase

	// ErrPasswordInvalid - 401: Password invalid.
	ErrPasswordInvalid

	// ErrPasswordIncorrect - 401: Password incorrect.
	ErrPasswordIncorrect

	// ErrPasswordHash - 500: Error occurred while hashing password.
	ErrPasswordHash

	// ErrClientRPCInit - 500: RPC client initialization error.
	ErrClientRPCInit

	// ErrRPCProcess - 500: RPC service process error.
	ErrRPCProcess

	// ErrRPCLink - 500: RPC service link error.
	ErrRPCLink

	// ErrRPCMutualCall - 500: RPC mutual call error.
	ErrRPCMutualCall

	// ErrEncodingFailed - 500: Encoding failed.
	ErrEncodingFailed

	// ErrMinio - 500: Minio error.
	ErrMinio

	// ErrOpenFormFile - 500: Open request's form file error.
	ErrOpenFormFile
)
