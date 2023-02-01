namespace go publish

struct PublishRequest {
    1: required i64 uid; // 登录用户 id，用于 UserQueryRequest
    2: required string Title; // 视频标题
    3: required binary data; // 视频数据
}

struct PublishResponse {
    1: required i32 StatusCode; // 状态码，0 - 成功，其他值 - 失败
    2: optional string StatusMsg; // 返回状态描述
}

service PublishService{
    PublishResponse publishVideo(1: PublishRequest req); // 刷视频，返回一个视频列表
}