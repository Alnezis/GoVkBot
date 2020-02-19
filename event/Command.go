package event

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type Command struct {
	Command        string
	Args           []string
	PrivateMessage *object.PrivateMessage
}

func (c *Command) GetName() string {
	return CommandEvent
}
