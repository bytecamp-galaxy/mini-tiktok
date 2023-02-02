namespace go publish

struct PublishRequest {
    1: required i64 UserId; // 登录用户 id，用于 UserQueryRequest
    2: required string Title; // 视频标题
    3: required binary Data; // 视频数据
}

struct PublishResponse {}

service PublishService {
    PublishResponse publishVideo(1: PublishRequest req); // 刷视频，返回一个视频列表
}