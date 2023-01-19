namespace go user

struct UserRegisterRequest {
    1: string Username;
    2: string Password;
}

struct UserRegisterResponse {
    1: i32 StatusCode;
    2: string StatusMsg;
    3: i64 UserId;
    4: string Token;  // usused
}

struct UserLoginRequest {
    1: string Username;
    2: string Password;
}

struct UserLoginResponse {
    1: i32 StatusCode;
    2: string StatusMsg;
    3: i64 UserId;
    4: string Token;  // usused
}

service UserService {
    UserRegisterResponse userRegister(1: UserRegisterRequest req);
    UserLoginResponse userLogin(1: UserLoginRequest req);
}