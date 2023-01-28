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
                        Comment Service
====================================================================*/
struct CommentActionRequest {
    1: required i64 VideoId (api.query="video_id");
    3: required i32 ActionType (api.query="action_type", api.vd="$ == 1 || $ == 2");
    4: optional string CommentText (api.query="comment_text");
    5: optional i64 CommentId (api.query="comment_id");
    6: required string Token (api.query="token");
}

struct CommentActionResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: optional Comment Comment (api.body="comment");
}

struct CommentListRequest {
    1: required i64 VideoId (api.query="video_id");
    2: required string Token (api.query="token");
}

struct CommentListResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: list<Comment> CommentList (api.body="comment_list");
}

struct Comment {
    1: required i64 Id (api.body="id");
    2: required User User (api.body="user");
    3: required string Content (api.body="content");
    4: required string CreateDate (api.body="create_date");
}

service CommentApi {
    CommentActionResponse commentAction(1: CommentActionRequest req) (api.post="/douyin/comment/action/");
    CommentListResponse commentList(1: CommentListRequest req) (api.get="/douyin/comment/list/");
}