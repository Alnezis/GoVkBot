package event

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type MessageReply struct {
	Message *object.Message `json:"" map:""`
}

func (m MessageReply) GetName() string {
	return MessageReplyEvent
}
