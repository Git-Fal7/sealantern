package command

import (
	"fmt"
	"strings"
)

type Manager struct {
	commands map[string]SimpleCommand
}

func (m *Manager) Register(commandName string, simpleCommand SimpleCommand, aliases ...string) error {
	err := m.register(commandName, simpleCommand)
	if err != nil {
		return err
	}
	for _, alias := range aliases {
		err = m.register(alias, simpleCommand)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) register(commandName string, simpleCommand SimpleCommand) error {
	cmdName := strings.ToLower(commandName)
	if _, ok := m.commands[cmdName]; ok {
		return fmt.Errorf("already registered a command called %s", cmdName)
	}
	m.commands[cmdName] = simpleCommand
	return nil
}

func (m *Manager) GetCommand(commandName string) (SimpleCommand, error) {
	v, ok := m.commands[strings.ToLower(commandName)]
	if !ok {
		return nil, fmt.Errorf("command not found")
	}
	return v, nil
}

func NewManager() *Manager {
	return &Manager{
		commands: map[string]SimpleCommand{},
	}
}
