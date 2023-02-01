namespace go feed
include "user_service.thrift"

struct FeedRequest {
    1: optional i64 LatestTime; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
    2: optional string Token; // 可选参数，登录用户设置
}

// 例如当前请求的 latest_time 为 9:00，那么返回的视频列表时间戳为 [8:55,7:40, 6:30, 6:00]
// 所有这些视频中，最早发布的是 6:00 的视频，那么 6:00 作为下一次请求时的 latest_time
// 那么下次请求返回的视频时间戳就会小于 6:00
struct FeedResponse {
    1: i32 StatusCode; // 状态码，0-成功，其他值-失败
    2: optional string StatusMsg; // 返回状态描述
    3: list<Video> VideoList; // 视频列表
    4: optional i64 NextTime; // 本次返回的视频中，发布最早的时间，作为下次请求时的 latest_time
}

struct Video {
    1: i64 Id; // 视频唯一标识
    2: user_service.User author; // 视频作者信息
    3: string PlayUrl; // 视频播放地址
    4: string CoverUrl; // 视频封面地址
    5: i64 FavoriteCount; // 视频的点赞总数
    6: i64 CommentCount; // 视频的评论总数
    7: bool IsFavorite; // true-已点赞，false-未点赞
    8: string Title; // 视频标题
}

service FeedService{
    FeedResponse getFeed(1: FeedRequest req); // 刷视频，返回一个视频列表
}