package keyboard

import (
	"encoding/json"
	"github.com/Alnezis/GoVkBot/vk"
	"log"
	"net/url"
)

type Action struct {
	Type    string `json:"type,omitempty"`
	Payload string `json:"payload,omitempty"`
	Link    string `json:"link,omitempty"`
	Label   string `json:"label,omitempty"`
	AppID   int    `json:"app_id,omitempty"`
	OwnerID int    `json:"owner_id,omitempty"`
	Hash    string `json:"hash,omitempty"`
}

type Button struct {
	Action Action `json:"action,omitempty"`
	Color  string `json:"color,omitempty"`
}

func New(inline bool, args ...[]Button) vk.H {
	h := vk.H{}
	params := vk.H{}

	h["inline"] = inline
	h["one_time"] = false
	h["buttons"] = args
	jsonKeyboard, err := json.Marshal(h)
	if err != nil {
		log.Println(err)

	}
	params["keyboard"] = string(jsonKeyboard)
	return params
}

func NewOne(oneTime bool, args ...[]Button) vk.H {
	h := vk.H{}
	params := vk.H{}

	h["inline"] = false
	h["one_time"] = oneTime
	h["buttons"] = args
	jsonKeyboard, err := json.Marshal(h)
	if err != nil {
		log.Println(err)
	}
	params["keyboard"] = string(jsonKeyboard)
	return params
}

func ButtonText(label, command, color string, payload map[string]interface{}) Button {
	payload["command"] = command
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}

	b := Button{
		Action: Action{
			Type:    "text",
			Payload: string(jsonPayload),
			Label:   url.QueryEscape(label),
		},
		Color: color,
	}
	return b
}

func ButtonLink(label, link string) Button {
	b := Button{
		Action: Action{
			Type:  "open_link",
			Link:  link,
			Label: url.QueryEscape(label),
		},
	}

	return b
}

func ButtonApp(label string, ownerId, appId int) Button {
	b := Button{
		Action: Action{
			Type:    "open_app",
			OwnerID: ownerId,
			AppID:   appId,
			Label:   url.QueryEscape(label),
		},
	}

	return b
}

func ButtonCallBack(label, command, color string, payload map[string]interface{}) Button {
	payload["command"] = command
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	return Button{
		Action: Action{
			Type:    "callback",
			Payload: string(jsonPayload),
			Label:   url.QueryEscape(label),
		},
		Color: color,
	}
}
