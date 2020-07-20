package object

const (
	MESSAGEACTION_CHAT_PHOTO_UPDATE        = "chat_photo_update"
	MESSAGEACTION_CHAT_PHOTO_REMOVE        = "chat_photo_REMOVE"
	MESSAGEACTION_CHAT_TITLE_UPDATE        = "chat_title_update"
	MESSAGEACTION_CHAT_INVITE_USER         = "chat_invite_user"
	MESSAGEACTION_CHAT_INVITE_USER_BY_LINK = "chat_invite_user_by_link"
	MESSAGEACTION_CHAT_KICK_USER           = "chat_kick_user"
	MESSAGEACTION_CHAT_PIN_MESSAGE         = "chat_pin_message"
	MESSAGEACTION_CHAT_UNPIN_MESSAGE       = "chat_unpin_message"
)

type MessageActionPhoto struct {
	Photo50  string `json:"photo_50"`
	Photo100 string `json:"photo_100"`
	Photo200 string `json:"photo_200"`
}

type MessageAction struct {
	Type     string             `json:"type" map:"type"`
	MemberID float64            `json:"member_id" map:"member_id"`
	Text     string             `json:"text" map:"text"`
	Email    string             `json:"email" map:"email"`
	Photo    MessageActionPhoto `json:"photo" map:"photo"`
}

type PrivateMessage struct {
	Message    *Message    `json:"message" map:"message"`
	ClientInfo *ClientInfo `json:"client_info" map:"client_info"`
}

type ClientInfo struct {
	ButtonActions  []string `json:"button_actions" map:"button_actions"`
	Keyboard       bool     `json:"keyboard" map:"keyboard"`
	InlineKeyboard bool     `json:"inline_keyboard" map:"inline_keyboard"`
	Carousel       bool     `json:"carousel" map:"carousel"`
	LangID         float64  `json:"lang_id" map:"lang_id"`
}

type Message struct {
	ID float64 `json:"id" map:"id"`
	// Conversation Message ID
	MessageID         float64           `json:"conversation_message_id" map:"conversation_message_id"`
	Date              float64           `json:"date" map:"date"`
	PeerID            float64           `json:"peer_id" map:"peer_id"`
	ExpireTTL         float64           `json:"expire_ttl,omitempty" map:"expire_ttl,omitempty"`
	FromID            float64           `json:"from_id" map:"from_id"`
	UserID            float64           `json:"from_id" map:"from_id"`
	Text              string            `json:"text" map:"text"`
	RandomID          float64           `json:"random_id" map:"random_id"`
	Ref               string            `json:"ref" map:"ref"`
	RefSource         string            `json:"ref_source" map:"ref_source"`
	Attachments       []*Attachment     `json:"attachments" map:"attachments"`
	Important         bool              `json:"important" map:"important"`
	Geo               *Geo              `json:"geo" map:"geo"`
	Payload           string            `json:"payload" map:"payload"`
	ForwardedMessages []*PrivateMessage `json:"fwd_messages" map:"fwd_messages"`
	ReplyMessage      *PrivateMessage   `json:"reply_message" map:"reply_message"`
	Action            *MessageAction    `json:"action" map:"action"`
}
