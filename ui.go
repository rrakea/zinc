package main

import (
	text "github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	packages []Package
	cursor   int
	input    text.Model
}

func default_model() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			install(m.packages[m.cursor])
		case "q", "crtl+c":
			cmd = tea.Quit
		}
	}

	return m, cmd
}

func (m Model) View() string {
	return ""
}
