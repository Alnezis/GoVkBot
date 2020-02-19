package event

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type MessageEdit struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageEdit) GetName() string {
	return MessageEditEvent
}
