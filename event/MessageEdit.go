package event

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type MessageEdit struct {
	Message *object.Message `json:"" map:""`
}

func (m MessageEdit) GetName() string {
	return MessageEditEvent
}
