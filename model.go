package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type pane int

const (
	listPane pane = iota
	detailPane
)

type screen int

const (
	screenHome screen = iota
	screenSettings
)

type model struct {
	list       list.Model
	activePane pane
	screen     screen
	width      int
	height     int
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		lw := leftWidth(msg.Width)
		m.list.SetSize(lw-4, msg.Height-6)
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
		case "tab":
			if m.screen == screenHome {
				if m.activePane == listPane {
					m.activePane = detailPane
				} else {
					m.activePane = listPane
				}
			}
			return m, nil
		case "enter":
			if m.screen == screenHome {
				m.activePane = detailPane
			}
			return m, nil
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
