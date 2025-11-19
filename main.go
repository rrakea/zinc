package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	packages := Search("yay")
	_ = packages

	p := tea.NewProgram(default_model())
	if _, err := p.Run(); err != nil {
		log.Fatal("Bubble tea error")
	}
}
