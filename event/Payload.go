package event

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type Payload struct {
	Payload        map[string]string
	PrivateMessage *object.PrivateMessage
}

func (c *Payload) GetName() string {
	return PayloadEvent
}
