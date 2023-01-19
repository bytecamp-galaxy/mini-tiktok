namespace go api

struct UserRegisterRequest {
    1: required string Username (api.query="username", api.vd="(len($) > 0 && len($) < 32); msg:'Illegal format'");
    2: required string Password (api.query="password", api.vd="(len($) > 0 && len($) < 32); msg:'Illegal format'");
}

struct UserRegisterResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: required string StatusMsg (api.body="status_msg");
    3: required i64 UserId (api.body="user_id");
    4: required string Token (api.body="token");
}

struct UserLoginRequest {
    1: required string Username (api.query="username", api.vd="(len($) > 0 && len($) < 32); msg:'Illegal format'");
    2: required string Password (api.query="password", api.vd="(len($) > 0 && len($) < 32); msg:'Illegal format'");
}

struct UserLoginResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: required string StatusMsg (api.body="status_msg");
    3: required i64 UserId (api.body="user_id");
    4: required string Token (api.body="token");
}

struct UserRequest {
    1: required i64 UserId (api.query="user_id");
    2: required string Token (api.query="token");
}

struct UserResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: required string StatusMsg (api.body="status_msg");
}

service UserApi {
    UserRegisterResponse userRegister(1: UserRegisterRequest req) (api.post="/douyin/user/register/");
    UserLoginResponse userLogin(1: UserLoginRequest req) (api.post="/douyin/user/login/");
    UserResponse userQuery(1: UserRequest req) (api.get="/douyin/user/");
}

struct PublishActionRequest {

}

struct PublishActionResponse {

}

service PublishApi {
    PublishActionResponse publishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/");
}