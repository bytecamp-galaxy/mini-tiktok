namespace go api

struct UserRegisterRequest {
    1: required string Username (api.query="username");
    2: required string Password (api.query="password");
}

struct UserRegisterResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: required string StatusMsg (api.body="status_msg");
    3: required i64 UserId (api.body="user_id");
    4: required string Token (api.body="token");
}

service UserApi {
    UserRegisterResponse userRegister(1: UserRegisterRequest req) (api.post="/douyin/user/register/");
}

struct PublishActionRequest {

}

struct PublishActionResponse {

}

service PublishApi {
    PublishActionResponse publishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/");
}