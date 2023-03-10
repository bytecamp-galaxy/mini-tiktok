namespace go api

/*==================================================================
                        User Service
====================================================================*/
struct UserRegisterRequest {
    1: required string Username (api.query="username", api.vd="(len($) > 6 && len($) < 32)");
    2: required string Password (api.query="password", api.vd="(len($) > 6 && len($) < 32)");
}

struct UserRegisterResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: required i64 UserId (api.body="user_id");
    4: required string Token (api.body="token");
}

struct UserLoginRequest {
    1: required string Username (api.query="username", api.vd="(len($) > 6 && len($) < 32)");
    2: required string Password (api.query="password", api.vd="(len($) > 6 && len($) < 32)");
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
    6: optional string Avatar (api.body="avatar");
    7: optional string BackgroundImage (api.body="background_image");
    8: optional string Signature (api.body="signature");
    9: optional i64 TotalFavorited (api.body="total_favorited");
    10: optional i64 WorkCount (api.body="work_count");
    11: optional i64 FavoriteCount (api.body="favorite_count");
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
    2: optional string Token (api.query="token"); // ?????????????????????????????????
}

// ????????????????????? latest_time ??? 9:00?????????????????????????????????????????? [8:55, 7:40, 6:30, 6:00]
// ?????????????????????????????????????????? 6:00 ?????????????????? 6:00 ??????????????????????????? latest_time
// ?????????????????????????????????????????????????????? 6:00
struct FeedResponse {
    1: required i32 StatusCode (api.body="status_code"); // ????????????0 - ?????????????????? - ??????
    2: optional string StatusMsg (api.body="status_msg"); // ??????????????????
    3: required list<Video> VideoList (api.body="video_list"); // ????????????
    4: optional i64 NextTime (api.body="next_time"); // ??????????????????????????????????????????????????????????????????????????? latest_time
}

struct Video {
    1: required i64 Id (api.body="id"); // ??????????????????
    2: required User author (api.body="author"); // ??????????????????
    3: required string PlayUrl (api.body="play_url"); // ??????????????????
    4: required string CoverUrl (api.body="cover_url"); // ??????????????????
    5: required i64 FavoriteCount (api.body="favorite_count"); // ?????????????????????
    6: required i64 CommentCount (api.body="comment_count"); // ?????????????????????
    7: required bool IsFavorite (api.body="is_favorite"); // true - ????????????false - ?????????
    8: required string Title (api.body="title"); // ????????????
}

service FeedApi {
    FeedResponse getFeed(1: FeedRequest req) (api.get="/douyin/feed/");
}

/*==================================================================
                        Publish Service
====================================================================*/
struct PublishActionRequest {
    1: required string Token (api.form="token"); // ??????????????????
    2: required string Title (api.form="title"); // ????????????
    3: required binary Data (api.form="data"); // ????????????
}

struct PublishActionResponse {
    1: required i32 StatusCode (api.body="status_code"); // ????????????0 - ?????????????????? - ??????
    2: optional string StatusMsg (api.body="status_msg"); // ??????????????????
}

struct PublishListRequest {
    1: required string Token (api.query="token"); // ??????????????????
    2: required i64 UserId (api.query="user_id"); // ?????? id
}

struct PublishListResponse {
    1: required i32 StatusCode (api.body="status_code"); // ????????????0 - ?????????????????? - ??????
    2: optional string StatusMsg (api.body="status_msg"); // ??????????????????
    3: required list<Video> VideoList (api.body="video_list")
}


service PublishApi {
    PublishActionResponse publishAction(1: PublishActionRequest req) (api.post="/douyin/publish/action/");
    PublishListResponse publishList(1: PublishListRequest req) (api.get="/douyin/publish/list/");
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
    3: required list<Comment> CommentList (api.body="comment_list"); // json:"comment_list,required"`
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
    1: required string Token (api.query="token"); // ??????
    2: required i64 VideoId (api.query="video_id");
    3: required i32 ActionType (api.query="action_type", api.vd="$ == 1 || $ == 2");
}

struct FavoriteActionResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
}

struct FavoriteListRequest {
    1: required i64 UserId (api.query="user_id");
    2: required string Token (api.query="token");
}

struct FavoriteListResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: required list<Video> VideoList (api.body="video_list");
}

service FavoriteApi {
    FavoriteActionResponse favoriteAction(1: FavoriteActionRequest req) (api.post="/douyin/favorite/action/"); // ?????? / ????????????
    FavoriteListResponse favoriteList(1: FavoriteListRequest req) (api.get="/douyin/favorite/list/"); // ??????????????????
}

/*==================================================================
                        Relation Service
====================================================================*/
struct RelationActionRequest {
    1: required string Token (api.query="token"); // ??????
    2: required i64 ToUserId (api.query="to_user_id");
    3: required i32 ActionType (api.query="action_type", api.vd="$ == 1 || $ == 2");
}

struct RelationActionResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
}

struct RelationFollowListRequest {
    1: required string Token (api.query="token");
    2: required i64 UserId (api.query="user_id");
}

struct RelationFollowListResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: required list<User> UserList (api.body="user_list");
}

struct RelationFollowerListRequest {
    1: required string Token (api.query="token");
    2: required i64 UserId (api.query="user_id");
}

struct RelationFollowerListResponse {
    1: required i32 StatusCode (api.body="status_code");
    2: optional string StatusMsg (api.body="status_msg");
    3: required list<User> UserList (api.body="user_list");
}

service RelationApi {
    RelationActionResponse relationAction(1: RelationActionRequest req) (api.post="/douyin/relation/action/");
    RelationFollowListResponse relationFollowList(1: RelationFollowListRequest req) (api.get="/douyin/relation/follow/list/");
    RelationFollowerListResponse relationFollowerList(1: RelationFollowerListRequest req) (api.get="/douyin/relation/follower/list/");
}