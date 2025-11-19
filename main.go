package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	input := make(chan string, 16)
	search := make(chan []Package, 16)

	input_chan = input
	search_chan = search

	p := tea.NewProgram(default_model())

	go func() {
		if _, err := p.Run(); err != nil {
			log.Fatal("Bubble tea error")
		}
	}()

	for {
		select {
		case packages := <-search:
			p.Send(packages)
		case input_str := <-input:
			if len(input_str) > 5 {
				go Search(input_str)
			}
		}
	}
}
