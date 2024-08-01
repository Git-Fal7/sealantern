package chat

import "fmt"

type Color struct {
	char byte
	name string
}

func (color Color) GetName() *string {
	return &color.name
}

func (color Color) String() string {
	return fmt.Sprintf("§%c", color.char)
}

var (
	Black       = Color{'0', "black"}
	DarkBlue    = Color{'1', "dark_blue"}
	DarkGreen   = Color{'2', "dark_green"}
	DarkAqua    = Color{'3', "dark_aqua"}
	DarkRed     = Color{'4', "dark_red"}
	DarkPurple  = Color{'5', "dark_purple"}
	Gold        = Color{'6', "gold"}
	Gray        = Color{'7', "gray"}
	DarkGray    = Color{'8', "dark_gray"}
	Blue        = Color{'9', "blue"}
	Green       = Color{'a', "green"}
	Aqua        = Color{'b', "aqua"}
	Red         = Color{'c', "red"}
	LightPurple = Color{'d', "light_purple"}
	Yellow      = Color{'e', "yellow"}
	White       = Color{'f', "white"}
)
