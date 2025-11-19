package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	input := make(chan string)
	search := make(chan []Package)

	input_chan = input
	search_chan = search

	p := tea.NewProgram(default_model())
	if _, err := p.Run(); err != nil {
		log.Fatal("Bubble tea error")
	}

	for {
		select {
		case packages := <-search:
			p.Send(packages)
		case input_str := <-input:
			if len(input_str) > 2 {
				go Search(input_str)
			}
		}
	}
}
