namespace go api

struct UserRegisterRequest {
    1: required string Username (api.query="username", api.vd="(len($) > 0 && len($) < 32)");
    2: required string Password (api.query="password", api.vd="(len($) > 0 && len($) < 32)");
}

struct UserRegisterResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: required i64 UserId (api.body="user_id");
    4: required string Token (api.body="token");
}

struct UserLoginRequest {
    1: required string Username (api.query="username", api.vd="(len($) > 0 && len($) < 32)");
    2: required string Password (api.query="password", api.vd="(len($) > 0 && len($) < 32)");
}

struct UserLoginResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: required i64 UserId (api.body="user_id");
    4: required string Token (api.body="token");
}

struct UserQueryRequest {
    1: required i64 UserId (api.query="user_id");
    2: required string Token (api.query="token");
}

struct UserQueryResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: required User User (api.body="user");
}

struct User {
    1: required i64 Id (api.body="id");
    2: required string Name (api.body="name");
    3: optional i64 FollowCount (api.body="follow_count");
    4: optional i64 FollowerCount (api.body="follower_count");
    5: required bool IsFollow (api.body="is_follow");
}

service UserApi {
    UserRegisterResponse userRegister(1: UserRegisterRequest req) (api.post="/douyin/user/register/");
    UserLoginResponse userLogin(1: UserLoginRequest req) (api.post="/douyin/user/login/");
    UserQueryResponse userQuery(1: UserQueryRequest req) (api.get="/douyin/user/");
}

struct PublishActionRequest {

}

struct PublishActionResponse {

}

service PublishApi {
    PublishActionResponse publishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/");
}

/*==================================================================
                        Favorite Service
====================================================================*/

struct FavoriteActionsRequest{
    1: required i64 UserId (api.query="user_id");
    2: required string Token (api.query="token"); //鉴权
    3: required i64 VideoId (api.query="video_id");
    4: required i32 ActionType (api.query="action_type", api.vd="$ == 1 || $ == 2");
}

struct FavoriteActionResponse{
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
}

struct FavoriteListRequest{
    1: required i64 UserId (api.query="id");
    2: required string Token (api.query="token");
}

struct FavoriteListResponse{
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
//    3: list<feed_service.Video> VidoeList;
}

service FavoriteService{
    FavoriteActionResponse favoriteAction(1: FavoriteActionsRequest req) (api.post="/douyin/favorite/action/"); // 点赞/取消点赞
    FavoriteListResponse favoriteList(1: FavoriteActionsRequest req) (api.get="/douyin/favorite/list/"); //return 点赞视频列表
}