// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Code generated by "codegen -output ./errno_generated.go -type int base.go"; DO NOT EDIT.

package errno

// Init register error codes defines in this source code to `github.com/marmotedu/errors`
func Init() {
	register(ErrUnknown, 500, "Internal server error")
	register(ErrGenerateToken, 500, "Error occurred while generating token")
	register(ErrParseToken, 500, "Error occurred while parsing from token")
	register(ErrBindAndValidation, 400, "Error occurred while binding the request body to the struct or validation failed")
	register(ErrDatabase, 500, "Database error")
	register(ErrRedis, 500, "Redis error")
	register(ErrPasswordInvalid, 400, "Password invalid")
	register(ErrPasswordIncorrect, 401, "Password incorrect")
	register(ErrPasswordHash, 500, "Error occurred while hashing password")
	register(ErrClientRPCInit, 500, "RPC client initialization error")
	register(ErrRPCLink, 500, "RPC service link error")
	register(ErrEncodingFailed, 500, "Encoding failed")
	register(ErrMinio, 500, "Minio error")
	register(ErrOpenFormFile, 500, "Open request's form file error")
	register(ErrInvalidUser, 400, "User does not exist")
	register(ErrInvalidVideo, 400, "Video does not exist")
	register(ErrInvalidVideoType, 400, "Uploaded video type unsupported. (")
}
