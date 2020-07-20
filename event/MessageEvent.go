package event

type ObjectMessageEvent struct {
	UserID  float64           `json:"user_id" map:"user_id"`
	PeerID  float64           `json:"peer_id" map:"peer_id"`
	EventID string            `json:"event_id" map:"event_id"`
	Payload map[string]string `json:"payload" map:"payload"`
	// Conversation Message ID
	MessageID float64 `json:"conversation_message_id" map:"conversation_message_id"`
}

func (m ObjectMessageEvent) GetName() string {
	return MessageEvent
}
