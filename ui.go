package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	input   textinput.Model
	table   table.Model
	row_map map[string]Package
}

var input_chan chan string

func default_model() Model {
	input := textinput.New()
	input.Placeholder = "Type to search"
	input.Focus()
	input.CharLimit = 200
	input.Width = 20

	columns := []table.Column{
		{Title: "Name", Width: 10},
		{Title: "Description", Width: 30},
	}
	rows := []table.Row{}
	table := table.New(table.WithColumns(columns), table.WithRows(rows), table.WithFocused(false))

	return Model{
		input: input,
		table: table,
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd_input tea.Cmd
		cmd_table tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			row := m.table.SelectedRow()
			if len(row) > 0 {
				install(m.row_map[row[0]])
			}
			return m, tea.Quit
		case "esc", "crtl+c":
			return m, tea.Quit
		}

	case []Package:
		m.row_map = make(map[string]Package)
		rows := []table.Row{}
		for _, p := range msg {
			rows = append(rows, table.Row{p.Name, p.Desc})
			m.row_map[p.Name] = p
		}

		m.table.SetRows(rows)
	}

	go func() { input_chan <- m.input.Value() }()

	m.input, cmd_input = m.input.Update(msg)
	m.table, cmd_table = m.table.Update(msg)

	return m, tea.Batch(cmd_input, cmd_table)
}

func (m Model) View() string {
	style := lipgloss.NewStyle()
	style.BorderStyle(lipgloss.DoubleBorder())
	style.BorderForeground(lipgloss.Color("0"))

	return style.Render(m.table.View()) + m.input.View()
}
