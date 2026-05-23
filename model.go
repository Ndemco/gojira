package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

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

var filterOptions = []string{
	"Assigned to me",
	"Reported by me",
	"All open issues",
	"By project",
}

const filterInnerHeight = 6 // "Filter\n\n" + 4 option rows

type model struct {
	list         list.Model
	activePane   pane
	screen       screen
	width        int
	height       int
	showFilter   bool
	filterCursor int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(leftWidth(msg.Width)-4, m.listInnerHeight())
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
		case "f":
			if m.screen == screenHome {
				m.showFilter = !m.showFilter
				if m.showFilter {
					m.activePane = filterPane
				} else {
					m.activePane = listPane
				}
				m.list.SetSize(leftWidth(m.width)-4, m.listInnerHeight())
			}
			return m, nil
		case "esc":
			if m.showFilter {
				m.showFilter = false
				m.activePane = listPane
				m.list.SetSize(leftWidth(m.width)-4, m.listInnerHeight())
			}
			return m, nil
		case "tab":
			if m.screen == screenHome {
				switch m.activePane {
				case listPane:
					m.activePane = detailPane
				case detailPane, filterPane:
					m.activePane = listPane
				}
			}
			return m, nil
		case "enter":
			if m.screen == screenHome {
				if m.activePane == filterPane {
					m.showFilter = false
					m.activePane = listPane
					m.list.SetSize(leftWidth(m.width)-4, m.listInnerHeight())
				} else {
					m.activePane = detailPane
				}
			}
			return m, nil
		case "up", "k":
			if m.activePane == filterPane {
				if m.filterCursor > 0 {
					m.filterCursor--
				}
				return m, nil
			}
		case "down", "j":
			if m.activePane == filterPane {
				if m.filterCursor < len(filterOptions)-1 {
					m.filterCursor++
				}
				return m, nil
			}
		}
	}

	if m.screen == screenHome && m.activePane == listPane {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func leftWidth(total int) int { return total * 2 / 5 }

func (m model) listInnerHeight() int {
	height := m.height - 6
	if m.showFilter {
		height -= filterInnerHeight + 2 // subtract filter box outer height (inner + 2 borders)
	}
	return height
}
