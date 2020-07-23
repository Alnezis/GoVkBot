package event

type MessageDeny struct {
	Event
	UserID int `json:"user_id" map:"user_id"`
}

func (m MessageDeny) GetName() string {
	return MessageDenyEvent
}
