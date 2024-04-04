package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	Template
}

func (i item) Title() string {
	return fmt.Sprintf("%s - %s", i.Template.Name, i.Template.Source)
}
func (i item) Description() string { return i.Template.Description }
func (i item) FilterValue() string { return i.Template.Name }

func initList(templates []Template) list.Model {
	var items []list.Item
	for _, template := range templates {
		items = append(items, item{Template: template})
	}

	return list.New(items, list.NewDefaultDelegate(), 0, 0)
}

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		f, _ := tea.LogToFile("./debug.log", "debug")
		defer f.Close()

		if msg.String() == "enter" {
			// selectedTemplate := m.list.SelectedItem().(item)
			// return m, SelectFlake(selectedTemplate)
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func SelectFlake(selected item) tea.Cmd {
	_, err := flakeInit(selected.Template)
	if err != nil {
		log.Fatal(err)
	}

	// print(output)
	return tea.Quit
}