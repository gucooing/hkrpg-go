package sociality

var CHAT *Chat

// 每个聊天室只保留最新200条消息

type Chat struct {
	ChatRoom map[uint64]*Room // 房间号:=(低位uid <<32) | 高位uid
}

type Room struct {
	lastChatTime int64 // 最后聊天时间
}

func newChat() *Chat {
	return &Chat{}
}

func GetChat() *Chat {
	if CHAT == nil {
		CHAT = newChat()
	}
	return CHAT
}
