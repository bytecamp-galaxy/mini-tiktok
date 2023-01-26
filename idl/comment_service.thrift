namespace go user

include "user_service.thrift"

struct CommentActionRequest {
    1: required i64 VideoId;
    2: required i64 UserId;
    3: required i32 ActionType;
    4: optional string CommentText;
    5: optional i64 CommentId;
}

struct CommentActionResponse {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: optional Comment Comment;
}

struct CommentListRequest {
    1: required i64 VideoId;
}

struct CommentListResponse {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: list<Comment> CommentList;
}

struct Comment {
    1: required i64 Id;
    2: required user_service.User User;
    3: required string Conent;
    4: required string CreateDate;
}

service CommentService {
    CommentActionResponse commentAction(1: CommentActionRequest req);
    CommentListResponse commentList(1: CommentListRequest req);
}