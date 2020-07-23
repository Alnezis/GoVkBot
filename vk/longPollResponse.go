package vk

type LongPollUpdate struct {
	EventType string      `json:"type"`
	Object    interface{} `json:"object"`
	GroupId   int         `json:"group_id"`
}

type LongPollResponse struct {
	TS      string           `json:"ts"`
	Updates []LongPollUpdate `json:"updates"`
	Failed  int              `json:"failed"`
}
