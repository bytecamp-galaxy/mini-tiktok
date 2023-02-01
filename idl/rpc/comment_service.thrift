namespace go comment

include "rpcmodel.thrift"

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
    3: optional rpcmodel.Comment Comment;
}

struct CommentListRequest {
    1: required i64 VideoId;
}

struct CommentListResponse {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: list<rpcmodel.Comment> CommentList;
}

service CommentService {
    CommentActionResponse commentAction(1: CommentActionRequest req);
    CommentListResponse commentList(1: CommentListRequest req);
}