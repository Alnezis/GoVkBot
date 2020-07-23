package event

type MessageAllow struct {
	Event
	UserID int    `json:"user_id" map:"user_id"`
	Key    string `json:"key" map:"key"`
}

func (m MessageAllow) GetName() string {
	return MessageAllowEvent
}
