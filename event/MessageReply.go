package event

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type MessageReply struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageReply) GetName() string {
	return MessageReplyEvent
}
