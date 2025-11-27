package main

import (
	"github.com/charmbracelet/bubbles/key"
)

type Keymap struct {
	keys map[string]key.Binding
}

func (k Keymap) ShortHelp() []key.Binding {
	return []key.Binding{k.keys["up"], k.keys["down"], k.keys["install"], k.keys["quit"]}
}

func (k Keymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}
func Default_keymap() Keymap {
	km := map[string]key.Binding{
		"up":      key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "Up")),
		"down":    key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "Down")),
		"install": key.NewBinding(key.WithKeys("enter"), key.WithHelp("↲", "Install")),
		"quit":    key.NewBinding(key.WithKeys("esc", "crtl+c"), key.WithHelp("Esc", "Quit")),
	}

	return Keymap{
		keys: km,
	}
}
