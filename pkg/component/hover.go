package component

import (
	"encoding/json"
)

type HoverAction struct {
	Action string      `json:"action"`
	Value  interface{} `json:"value"`
}

func HoverText(components []IChatComponent) *HoverAction {
	messages := make([]interface{}, 0)
	for _, c := range components {
		j, err := c.JSON()
		if err == nil {
			messages = append(messages, json.RawMessage(j))
		}
	}
	return &HoverAction{
		"show_text",
		messages,
	}
}

func HoverItem(itemjson string) *HoverAction {
	return &HoverAction{
		"show_item",
		itemjson,
	}
}

func HoverEntity(enttiyJson string) *HoverAction {
	return &HoverAction{
		"show_entity",
		enttiyJson,
	}
}
