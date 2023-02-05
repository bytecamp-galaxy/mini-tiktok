# 错误码

mini-tiktok 系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。

## 功能说明

如果返回结果中存在 `status_code` 字段，则表示调用 API 接口失败。例如：

```json
{
  "status_code": 100001,
  "status_msg": "Internal server error"
}
```

上述返回中 `status_code` 表示错误码，`status_msg` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500 (Internal Server Error)。

## 错误码列表

mini-tiktok 系统支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| ErrUnknown | 100001 | 500 | Internal server error |
| ErrGenerateToken | 100002 | 500 | Error occurred while generating token |
| ErrParseToken | 100003 | 500 | Error occurred while parsing from token |
| ErrBindAndValidation | 100004 | 400 | Error occurred while binding the request body to the struct or validation failed |
| ErrDatabase | 100005 | 500 | Database error |
| ErrPasswordInvalid | 100006 | 401 | Password invalid |
| ErrPasswordIncorrect | 100007 | 401 | Password incorrect |
| ErrPasswordHash | 100008 | 500 | Error occurred while hashing password |
| ErrClientRPCInit | 100009 | 500 | RPC client initialization error |
| ErrRPCLink | 100010 | 500 | RPC service link error |
| ErrEncodingFailed | 100011 | 500 | Encoding failed |
| ErrMinio | 100012 | 500 | Minio error |
| ErrOpenFormFile | 100013 | 500 | Open request's form file error |

