package main

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	input := make(chan string, 16)
	search := make(chan []Package, 16)

	input_chan = input
	search_chan = search

	p := tea.NewProgram(default_model() /* , tea.WithAltScreen()*/)

	go func() {
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
	}()

	if _, err := p.Run(); err != nil {
		log.Fatal("Bubble tea error")
	}

	fmt.Print("\033[H\033[2J")

	install()
}
