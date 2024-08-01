package chat

import "fmt"

type Decoration struct {
	char byte
}

func (color Decoration) String() string {
	return fmt.Sprintf("§%c", color.char)
}

var (
	Obfuscated    = Decoration{'k'}
	Bold          = Decoration{'l'}
	Strikethrough = Decoration{'m'}
	Underline     = Decoration{'n'}
	Italic        = Decoration{'o'}
	Reset         = Decoration{'r'}
)
