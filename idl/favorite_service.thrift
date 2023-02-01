namespace go favorite

include "feed_service.thrift"

struct FavoriteActionRequest {
    1: required i64 UserId;
    2: required string Token; //鉴权
    3: required i64 VideoId;
    4: required i32 ActionType; //1 点赞  2 取消
}

struct FavoriteActionResponse {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
}

struct FavoriteListRequest {
    1: required i64 UserId;
    2: required string Token;
}

struct FavoriteListResponse {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: list<feed_service.Video> VidoeList;
}

service FavoriteService {
    FavoriteActionResponse favoriteAction(1: FavoriteActionRequest req); // 点赞/取消点赞
    FavoriteListResponse favoriteList(1: FavoriteListRequest req); //return 点赞视频列表
}