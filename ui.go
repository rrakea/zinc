package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	input    textinput.Model
	packages table.Model
	row_map  map[string]Package
	info     string
	help     help.Model
	keys     Keymap
}

var input_chan chan string

const SEARCH_BOX_HEIGHT = 10

func default_model() Model {
	input := textinput.New()
	input.Placeholder = "Type to search"
	input.Focus()
	input.CharLimit = 200
	input.Width = 100

	columns := []table.Column{
		{Title: "Name", Width: 40},
	}
	rows := []table.Row{}
	table := table.New(table.WithColumns(columns), table.WithRows(rows), table.WithFocused(false), table.WithHeight(SEARCH_BOX_HEIGHT))

	return Model{
		input:    input,
		packages: table,
		info:     info(Package{}),
		help:     help.New(),
		keys:     Default_keymap(),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd_input tea.Cmd
		cmd_table tea.Cmd
		cmd       tea.Cmd
	)

	m.input, cmd_input = m.input.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.keys["install"]):
			row := m.packages.SelectedRow()
			if len(row) > 0 {
				to_install = m.row_map[row[0]].Name
			}
			cmd = tea.Quit
		case key.Matches(msg, m.keys.keys["down"]):
			if len(m.packages.Rows()) != 0 {
				m.packages.MoveDown(1)
				m.info = info(m.row_map[m.packages.SelectedRow()[0]])
			}
		case key.Matches(msg, m.keys.keys["up"]):
			if len(m.packages.Rows()) != 0 {
				m.packages.MoveUp(1)
				m.info = info(m.row_map[m.packages.SelectedRow()[0]])
			}
		case key.Matches(msg, m.keys.keys["quit"]):
			cmd = tea.Quit
		default:
			input_chan <- m.input.Value()
		}

	case []Package:
		m.row_map = make(map[string]Package)
		rows := []table.Row{}
		for _, p := range msg {
			rows = append(rows, table.Row{p.Name})
			m.row_map[p.Name] = p
		}

		m.packages.SetRows(rows)
		row := m.packages.SelectedRow()
		if len(row) > 0 {
			m.info = info(m.row_map[row[0]])
		}
	}

	m.packages, cmd_table = m.packages.Update(msg)

	return m, tea.Batch(cmd, cmd_input, cmd_table)
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(lipgloss.Color("#cba6f7"))

	help := m.help.View(m.keys)
	pack := style.Render(m.packages.View())
	info := style.Width(40).Render(m.info)
	input := style.Width(84).Render(m.input.View())

	top := lipgloss.JoinHorizontal(0, pack, info)
	full := lipgloss.JoinVertical(0, top, help, input)

	return full
}

func info(p Package) string {
	var out_of_date string
	if p.Out_of_date == 0 {
		out_of_date = "No"
	} else {
		out_of_date = "Yes"
	}

	var maintainer string
	if p.Maintainer == "" {
		maintainer = "Abandoned"
	} else {
		maintainer = p.Maintainer
	}

	str := "Name: " + p.Name + "\n" +
		"Version: " + p.Version + "\n" +
		"URL: " + p.Url + "\n" +
		"Out of Date: " + out_of_date + "\n" +
		"Maintainer: " + maintainer + "\n" +
		"Description: " + p.Desc
	return str
}
