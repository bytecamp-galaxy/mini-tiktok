package message

type MessageSendEvent struct {
	UserId     int64  `json:"user_id"`
	ToUserId   int64  `json:"to_user_id"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"from_user_id"`
	MsgContent string `json:"msg_content"`
}
