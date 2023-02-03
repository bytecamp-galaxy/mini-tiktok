namespace go user

include "rpcmodel.thrift"

struct UserRegisterRequest {
    1: required string Username;
    2: required string Password;
}

struct UserRegisterResponse {
    1: required i64 UserId;
}

struct UserLoginRequest {
    1: required string Username;
    2: required string Password;
}

struct UserLoginResponse {
    1: required i64 UserId;
}

struct UserQueryRequest {
    1: required i64 UserId;
}

struct UserQueryResponse {
    1: required rpcmodel.User User;
}

service UserService {
    UserRegisterResponse userRegister(1: UserRegisterRequest req);
    UserLoginResponse userLogin(1: UserLoginRequest req);
    UserQueryResponse userQuery(1: UserQueryRequest req);
}