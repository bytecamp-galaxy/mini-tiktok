namespace go publish

include "rpcmodel.thrift"

struct PublishRequest {
    1: required i64 UserId; // 登录用户 id，用于 UserQueryRequest
    2: required string Title; // 视频标题
    3: required binary Data; // 视频数据
}

struct PublishResponse {}

struct PublishListRequest {
    1: required i64 UserId; // 用户 id
    2: required i64 UserViewId;
}

struct PublishListResponse {
    1: required list<rpcmodel.Video> VideoList;
}

service PublishService {
    PublishResponse publishVideo(1: PublishRequest req); // 刷视频，返回一个视频列表
    PublishListResponse publishList(1: PublishListRequest req); // 获取用户曾上传的视频
}