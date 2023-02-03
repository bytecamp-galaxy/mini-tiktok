namespace go favorite

include "rpcmodel.thrift"

struct FavoriteActionRequest {
    1: required i64 UserId;
    2: required i64 VideoId;
    3: required i32 ActionType; // 1 点赞, 2 取消
}

struct FavoriteActionResponse {}

struct FavoriteListRequest {
    1: required i64 UserId;
}

struct FavoriteListResponse {
    1: required list<rpcmodel.Video> VideoList;
}

service FavoriteService {
    FavoriteActionResponse favoriteAction(1: FavoriteActionRequest req); // 点赞 / 取消点赞
    FavoriteListResponse favoriteList(1: FavoriteListRequest req); // return 点赞视频列表
}