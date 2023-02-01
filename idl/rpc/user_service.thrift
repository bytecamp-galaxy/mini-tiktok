namespace go user

include "rpcmodel.thrift"

struct UserRegisterRequest {
    1: string Username;
    2: string Password;
}

struct UserRegisterResponse {
    1: i64 UserId;
}

struct UserLoginRequest {
    1: string Username;
    2: string Password;
}

struct UserLoginResponse {
    1: i64 UserId;
}

struct UserQueryRequest {
    1: i64 UserId;
}

struct UserQueryResponse {
    1: rpcmodel.User User;
}

service UserService {
    UserRegisterResponse userRegister(1: UserRegisterRequest req);
    UserLoginResponse userLogin(1: UserLoginRequest req);
    UserQueryResponse userQuery(1: UserQueryRequest req);
}