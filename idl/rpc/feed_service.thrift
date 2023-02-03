namespace go feed

include "rpcmodel.thrift"

struct FeedRequest {
    1: optional i64 LatestTime; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: required i64 uid; // 用户参数。未设置时，为负数；当此字段设置时，需要返回用户视频的点赞信息
}

// 例如当前请求的 latest_time 为 9:00，那么返回的视频列表时间戳为 [8:55, 7:40, 6:30, 6:00]
// 所有这些视频中，最早发布的是 6:00 的视频，那么 6:00 作为下一次请求时的 latest_time
// 那么下次请求返回的视频时间戳就会小于 6:00
struct FeedResponse {
    3: list<rpcmodel.Video> VideoList; // 视频列表
    4: optional i64 NextTime; // 本次返回的视频中，发布最早的时间，作为下次请求时的 latest_time
}

service FeedService {
    FeedResponse getFeed(1: FeedRequest req); // 刷视频，返回一个视频列表
}