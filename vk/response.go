package vk

import (
	"github.com/Alnezis/GoVkBot/vk/object"
)

type Response struct {
	Response interface{}
	Error    *object.ResponseError
	Raw      map[string]interface{}
}
