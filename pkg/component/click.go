package component

import "fmt"

type ClickAction struct {
	Action string      `json:"action"`
	Value  interface{} `json:"value"`
}

func ClickOpenURL(url string) *ClickAction {
	return &ClickAction{
		"open_url",
		url,
	}
}

// Wont work most likely
func ClickOpenFile(file string) *ClickAction {
	return &ClickAction{
		"open_file",
		file,
	}
}

func ClickRunCommand(command string) *ClickAction {
	return &ClickAction{
		"run_command",
		command,
	}
}

func ClickSuggestCommand(command string) *ClickAction {
	return &ClickAction{
		"suggest_command",
		command,
	}
}

// Only for Books
func ClickChangePage(page uint) *ClickAction {
	return &ClickAction{
		"change_page",
		fmt.Sprintf("%d", page),
	}
}
