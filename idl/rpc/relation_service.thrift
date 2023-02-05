namespace go relation

include "rpcmodel.thrift"

struct RelationActionRequest {
    1: required i64 UserId;
    2: required i64 ToUserId;
    3: required i32 ActionType;
}

struct RelationActionResponse {}

struct RelationFollowListRequest {
    1: required i64 UserId;
    2: required i64 UserViewId;
}

struct RelationFollowListResponse {
    1: required list<rpcmodel.User> UserList;
}

struct RelationFollowerListRequest {
    1: required i64 UserId;
    2: required i64 UserViewId;
}

struct RelationFollowerListResponse {
    1: required list<rpcmodel.User> UserList;
}

service RelationService {
    RelationActionResponse relationAction(1: RelationActionRequest req);
    RelationFollowListResponse relationFollowList(1: RelationFollowListRequest req);
    RelationFollowerListResponse relationFollowerList(1: RelationFollowerListRequest req);
}