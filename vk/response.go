package vk

import (
	"GoVkBot/vk/object"
)

type Response struct {
	Response interface{}
	Error    *object.ResponseError
	Raw      map[string]interface{}
}
