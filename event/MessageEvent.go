package event

type ObjectMessageEvent struct {
	UserID  int               `json:"user_id" map:"user_id"`
	PeerID  int               `json:"peer_id" map:"peer_id"`
	EventID string            `json:"event_id" map:"event_id"`
	Payload map[string]string `json:"payload" map:"payload"`
	// Conversation Message ID
	MessageID int `json:"conversation_message_id" map:"conversation_message_id"`
}

func (m ObjectMessageEvent) GetName() string {
	return MessageEvent
}
