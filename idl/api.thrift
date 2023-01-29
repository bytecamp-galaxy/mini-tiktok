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

struct FavoriteActionRequest{
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
    3: list<Video> VidoeList (api.body="video_list");
}

service FavoriteService{
    FavoriteActionResponse favoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/action/"); // 点赞/取消点赞
    FavoriteListResponse favoriteList(1: FavoriteListRequest req) (api.get="/douyin/favorite/list/"); //return 点赞视频列表
}

struct Video {
    1: i64 Id (api.body="id"); // 视频唯一标识
    2: User author (api.body="auther"); // 视频作者信息
    3: string PlayUrl (api.body="play_url"); // 视频播放地址
    4: string CoverUrl (api.body="cover_url"); // 视频封面地址
    5: i64 FavoriteCount (api.body="favorite_count"); // 视频的点赞总数
    6: i64 CommentCount (api.body="comment_count"); // 视频的评论总数
    7: bool IsFavorite (api.body="is_favorite"); // true-已点赞，false-未点赞
    8: string Title (api.body="title"); // 视频标题
}