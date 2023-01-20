namespace go user

struct UserRegisterRequest {
    1: string Username;
    2: string Password;
}

struct UserRegisterResponse {
    1: i32 StatusCode;
    2: string StatusMsg;
    3: i64 UserId;
}

struct UserLoginRequest {
    1: string Username;
    2: string Password;
}

struct UserLoginResponse {
    1: i32 StatusCode;
    2: string StatusMsg;
    3: i64 UserId;
}

struct UserRequest {
    1: i64 UserId;
    2: string Token;
}

struct UserResponse {
    1: i32 StatusCode;
    2: string StatusMsg;
    3: User User;
}

struct User {
    1: i64 Id;
    2: string Name;
    3: i64 FollowCount;
    4: i64 FollowerCount;
    5: bool IsFollow;
}

service UserService {
    UserRegisterResponse userRegister(1: UserRegisterRequest req);
    UserLoginResponse userLogin(1: UserLoginRequest req);
    UserResponse userQuery(1: UserRequest req);
}