package event

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type MessageNew struct {
	PrivateMessage *object.PrivateMessage `json:"" map:""`
}

func (m MessageNew) GetName() string {
	return MessageNewEvent
}
