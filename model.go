package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type homeState struct {
	Issues       list.Model
	ActivePane   pane
	ShowFilter   bool
	FilterCursor int
}

type pane int

const (
	listPane pane = iota
	filterPane
	detailPane
)

type screen int

const (
	screenHome screen = iota
	screenSettings
)

type quickFilter struct {
	Name string
	JQL  string
}

var quickFilterOptions = []quickFilter{
	{"Assigned to me", "assignee = currentUser() ORDER BY updated DESC"},
	{"Reported by me", "reporter = currentUser() ORDER BY updated DESC"},
	{"All open issues", "resolution = Unresolved ORDER BY updated DESC"},
}

const filterInnerHeight = 6 // "Filter\n\n" + 4 option rows

type model struct {
	HomeState homeState
	screen    screen
	width     int
	height    int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.HomeState.Issues.SetSize(leftWidth(msg.Width)-4, m.listInnerHeight())
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m.screen = screenHome
			return m, nil
		case "2":
			m.screen = screenSettings
			return m, nil
		}

		switch m.screen {
		case screenHome:
			return m.updateHome(msg)
		}
	}

	return m, nil
}

func leftWidth(total int) int { return total * 2 / 5 }

func (m model) listInnerHeight() int {
	height := m.height - 6
	if m.HomeState.ShowFilter {
		height -= filterInnerHeight + 2 // subtract filter box outer height (inner + 2 borders)
	}
	return height
}
