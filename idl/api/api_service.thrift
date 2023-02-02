namespace go api

/*==================================================================
                        User Service
====================================================================*/
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

/*==================================================================
                        Feed Service
====================================================================*/
struct FeedRequest {
    1: optional i64 LatestTime (api.query="latest_time");
    2: optional string Token (api.query="token"); // 可选参数，登录用户设置
}

// 例如当前请求的 latest_time 为 9:00，那么返回的视频列表时间戳为 [8:55, 7:40, 6:30, 6:00]
// 所有这些视频中，最早发布的是 6:00 的视频，那么 6:00 作为下一次请求时的 latest_time
// 那么下次请求返回的视频时间戳就会小于 6:00
struct FeedResponse {
    1: required i32 StatusCode (api.body="status_code"); // 状态码，0 - 成功，其他值 - 失败
    2: optional string StatusMsg (api.body="status_msg"); // 返回状态描述
    3: required list<Video> VideoList (api.body="video_list"); // 视频列表
    4: optional i64 NextTime (api.body="next_time"); // 本次返回的视频中，发布最早的时间，作为下次请求时的 latest_time
}

struct Video {
    1: required i64 Id (api.body="id"); // 视频唯一标识
    2: required User author (api.body="auther"); // 视频作者信息
    3: required string PlayUrl (api.body="play_url"); // 视频播放地址
    4: required string CoverUrl (api.body="cover_url"); // 视频封面地址
    5: required i64 FavoriteCount (api.body="favorite_count"); // 视频的点赞总数
    6: required i64 CommentCount (api.body="comment_count"); // 视频的评论总数
    7: required bool IsFavorite (api.body="is_favorite"); // true - 已点赞，false - 未点赞
    8: required string Title (api.body="title"); // 视频标题
}

service FeedApi {
    FeedResponse getFeed(1: FeedRequest req) (api.get="/douyin/feed/");
}

/*==================================================================
                        Publish Service
====================================================================*/
struct PublishActionRequest {
    1: required string Token (api.body="token"); // 登录用户设置
    2: required string Title (api.body="title"); // 视频标题
    3: required binary data (api.body="data"); // 视频数据
}

struct PublishActionResponse {
    1: required i32 StatusCode (api.body="status_code"); // 状态码，0 - 成功，其他值 - 失败
    2: optional string StatusMsg (api.body="status_msg"); // 返回状态描述
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

/*==================================================================
                        Favorite Service
====================================================================*/
struct FavoriteActionRequest {
    1: required i64 UserId (api.query="user_id");
    2: required string Token (api.query="token"); // 鉴权
    3: required i64 VideoId (api.query="video_id");
    4: required i32 ActionType (api.query="action_type", api.vd="$ == 1 || $ == 2");
}

struct FavoriteActionResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
}

struct FavoriteListRequest {
    1: required i64 UserId (api.query="id");
    2: required string Token (api.query="token");
}

struct FavoriteListResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: list<Video> VideoList (api.body="video_list");
}

service FavoriteApi {
    FavoriteActionResponse favoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/action/"); // 点赞 / 取消点赞
    FavoriteListResponse favoriteList(1: FavoriteListRequest req) (api.get="/douyin/favorite/list/"); // 点赞视频列表
}