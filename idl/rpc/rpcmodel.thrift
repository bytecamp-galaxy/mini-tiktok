namespace go rpcmodel

struct User {
    1: required i64 Id;
    2: required string Name;
    3: required i64 FollowCount;
    4: required i64 FollowerCount;
    5: required bool IsFollow;
}

struct Comment {
    1: required i64 Id;
    2: required User User;
    3: required string Content;
    4: required string CreateDate;
}

struct Video {
    1: i64 Id; // 视频唯一标识
    2: User author; // 视频作者信息
    3: string PlayUrl; // 视频播放地址
    4: string CoverUrl; // 视频封面地址
    5: i64 FavoriteCount; // 视频的点赞总数
    6: i64 CommentCount; // 视频的评论总数
    7: bool IsFavorite; // true - 已点赞，false - 未点赞
    8: string Title; // 视频标题
}
