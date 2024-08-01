package component

import (
	"encoding/json"

	"github.com/git-fal7/sealantern/pkg/chat"
)

type IChatComponent interface {
	JSON() (string, error)
}

type ChatComponent struct {
	Bold          bool             `json:"bold,omitempty"`
	Italic        bool             `json:"italic,omitempty"`
	Underlined    bool             `json:"underlined,omitempty"`
	StrikeThrough bool             `json:"strikethrough,omitempty"`
	Obfuscated    bool             `json:"obfuscated,omitempty"`
	Color         *string          `json:"color,omitempty"`
	Insertion     *string          `json:"insertion,omitempty"`
	ClickEvent    *ClickAction     `json:"clickEvent,omitempty"`
	HoverEvent    *HoverAction     `json:"hoverEvent,omitempty"`
	Extra         []IChatComponent `json:"extra,omitempty"`
}

func (component *ChatComponent) SetBold(bold bool) {
	component.Bold = bold
}
func (component *ChatComponent) SetItalic(italic bool) {
	component.Italic = italic
}
func (component *ChatComponent) SetUnderlined(underlined bool) {
	component.Underlined = underlined
}
func (component *ChatComponent) SetStrikeThrough(strikeThrough bool) {
	component.StrikeThrough = strikeThrough
}
func (component *ChatComponent) SetObfuscated(obfuscated bool) {
	component.Obfuscated = obfuscated
}
func (component *ChatComponent) SetColor(color chat.Color) {
	component.Color = color.GetName()
}
func (component *ChatComponent) SetInsertion(insertion string) {
	component.Insertion = &insertion
}
func (component *ChatComponent) SetClickEvent(action *ClickAction) {
	component.ClickEvent = action
}
func (component *ChatComponent) SetHoverEvent(action *HoverAction) {
	component.HoverEvent = action
}
func (component *ChatComponent) AddExtra(extra IChatComponent) {
	component.Extra = append(component.Extra, extra)
}
func (component *ChatComponent) SetExtra(extra []IChatComponent) {
	component.Extra = extra
}
func (component *ChatComponent) JSON() (string, error) {
	b, err := json.Marshal(component)
	return string(b), err
}

type StringChatComponent struct {
	Text string `json:"text"`
	*ChatComponent
}

func (component *StringChatComponent) SetText(text string) {
	component.Text = text
}
func (component *StringChatComponent) JSON() (string, error) {
	b, err := json.Marshal(component)
	return string(b), err
}

func ChatMessage(text string) *StringChatComponent {
	return &StringChatComponent{
		text,
		&ChatComponent{
			false,
			false,
			false,
			false,
			false,
			nil,
			nil,
			nil,
			nil,
			nil,
		}}
}
