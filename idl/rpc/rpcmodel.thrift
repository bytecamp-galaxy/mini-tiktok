namespace go rpcmodel

struct User {
    1: required i64 Id;
    2: required string Name;
    3: required i64 FollowCount;
    4: required i64 FollowerCount;
    5: required bool IsFollow;
    6: required string Avatar;
    7: required string BackgroundImage;
    8: required string Signature;
    9: required i64 TotalFavorited;
    10: required i64 WorkCount;
    11: required i64 FavoriteCount;
}

struct Comment {
    1: required i64 Id;
    2: required User User;
    3: required string Content;
    4: required string CreateDate;
}

struct Video {
    1: required i64 Id; // 视频唯一标识
    2: required User author; // 视频作者信息
    3: required string PlayUrl; // 视频播放地址
    4: required string CoverUrl; // 视频封面地址
    5: required i64 FavoriteCount; // 视频的点赞总数
    6: required i64 CommentCount; // 视频的评论总数
    7: required bool IsFavorite; // true - 已点赞，false - 未点赞
    8: required string Title; // 视频标题
}
