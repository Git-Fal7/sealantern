package component

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ChatColor struct {
	char byte
	name string
}

func (color *ChatColor) GetChar() byte {
	return color.char
}
func (color *ChatColor) ChatFormat() string {
	return "&" + fmt.Sprintf("%c", color.char)
}
func (color *ChatColor) GetName() string {
	return color.name
}

type ChatStyle struct {
	*ChatColor
}

var (
	ChatColorBlack         = ChatColor{'0', "black"}
	ChatColorDarkBlue      = ChatColor{'1', "dark_blue"}
	ChatColorDarkGreen     = ChatColor{'2', "dark_green"}
	ChatColorDarkAqua      = ChatColor{'3', "dark_aqua"}
	ChatColorDarkRed       = ChatColor{'4', "dark_red"}
	ChatColorDarkPurple    = ChatColor{'5', "dark_purple"}
	ChatColorGold          = ChatColor{'6', "gold"}
	ChatColorGray          = ChatColor{'7', "gray"}
	ChatColorDarkGray      = ChatColor{'8', "dark_gray"}
	ChatColorIndigo        = ChatColor{'9', "blue"}
	ChatColorGreen         = ChatColor{'a', "green"}
	ChatColorAqua          = ChatColor{'b', "aqua"}
	ChatColorRed           = ChatColor{'c', "red"}
	ChatColorPink          = ChatColor{'d', "light_purple"}
	ChatColorYellow        = ChatColor{'e', "yellow"}
	ChatColorWhite         = ChatColor{'f', "white"}
	ChatStyleObfuscated    = ChatStyle{&ChatColor{'k', "obfuscated"}}
	ChatStyleBold          = ChatStyle{&ChatColor{'l', "bold"}}
	ChatStyleStrikeThrough = ChatStyle{&ChatColor{'m', "strikethrough"}}
	ChatStyleUnderlined    = ChatStyle{&ChatColor{'n', "underlined"}}
	ChatStyleItalic        = ChatStyle{&ChatColor{'o', "italic"}}
	ChatColorIds           = []ChatColor{
		ChatColorBlack,
		ChatColorDarkBlue,
		ChatColorDarkGreen,
		ChatColorDarkAqua,
		ChatColorDarkRed,
		ChatColorDarkPurple,
		ChatColorGold,
		ChatColorGray,
		ChatColorDarkGray,
		ChatColorIndigo,
		ChatColorGreen,
		ChatColorAqua,
		ChatColorRed,
		ChatColorPink,
		ChatColorYellow,
		ChatColorWhite,
	}
)

type ChatAction struct {
	Action string      `json:"action"`
	Value  interface{} `json:"value"`
}

const (
	chatClickTypeOpenUrl        = "open_url"
	chatClickTypeRunCommand     = "run_command"
	chatClickTypeSuggestCommand = "suggest_command"
	chatClickTypeChangePage     = "change_page"
	chatHoverTypeShowText       = "show_text"
	chatHoverTypeShowItem       = "show_item"
	chatHoverTypeShowEntity     = "show_entity"
)

func ChatClickRunCommand(command string) *ChatAction {
	return &ChatAction{
		chatClickTypeRunCommand,
		command,
	}
}

func ChatHoverText(component string) *ChatAction {
	return &ChatAction{
		chatHoverTypeShowText,
		component,
	}
}
func ChatHoverMessage(components []IChatComponent) *ChatAction {
	messages := make([]interface{}, 0)
	for _, c := range components {
		j, err := c.JSON()
		if err == nil {
			messages = append(messages, json.RawMessage(j))
		}
	}
	return &ChatAction{
		chatHoverTypeShowText,
		messages,
	}
}

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
	ClickEvent    *ChatAction      `json:"clickEvent,omitempty"`
	HoverEvent    *ChatAction      `json:"hoverEvent,omitempty"`
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
func (component *ChatComponent) SetColor(color *ChatColor) {
	component.Color = &color.name
}
func (component *ChatComponent) SetInsertion(insertion string) {
	component.Insertion = &insertion
}
func (component *ChatComponent) SetClickEvent(action *ChatAction) {
	component.ClickEvent = action
}
func (component *ChatComponent) SetHoverEvent(action *ChatAction) {
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

func BukkitMessageConvert(message string) IChatComponent {
	base := ChatMessage("")

	current := base
	buff := bytes.NewBufferString("")
	for i := 0; i < len(message); i++ {
		if message[i] == '&' && len(message) > i+1 {
			if (message[i+1] >= '0' && message[i+1] <= '9') ||
				(message[i+1] >= 'a' && message[i+1] <= 'f') ||
				message[i+1] == 'r' {
				// Create new node
				current.Text = buff.String()
				n := ChatMessage("")
				current.AddExtra(n)
				current = n
				buff.Reset()
			}

			if message[i+1] >= '0' && message[i+1] <= '9' {
				current.SetColor(&ChatColorIds[message[i+1]-'0'])
			} else if message[i+1] >= 'a' && message[i+1] <= 'f' {
				current.SetColor(&ChatColorIds[message[i+1]-'a'+10])
			} else {
				switch message[i+1] {
				case 'k':
					current.SetObfuscated(true)
				case 'l':
					current.SetBold(true)
				case 'm':
					current.SetStrikeThrough(true)
				case 'n':
					current.SetUnderlined(true)
				case 'o':
					current.SetItalic(true)
				case 'r':
					current.SetColor(&ChatColorWhite)
				}
			}
			i += 1
		} else {
			buff.WriteByte(message[i])
		}
	}

	return base
}
