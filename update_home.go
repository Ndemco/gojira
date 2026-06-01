package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updateHome(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
		case "f":
			m.HomeState.ShowFilter = !m.HomeState.ShowFilter
			if m.HomeState.ShowFilter {
				m.HomeState.ActivePane = filterPane
			} else {
				m.HomeState.ActivePane = listPane
			}
			m.HomeState.Issues.SetSize(leftWidth(m.width)-4, m.listInnerHeight())
		case "tab":
			switch m.HomeState.ActivePane {
			case listPane:
				m.HomeState.ActivePane = detailPane
			case detailPane, filterPane:
				m.HomeState.ActivePane = listPane
			}
			return m, nil
		case "up", "k":
			if m.HomeState.ActivePane == filterPane {
				if m.HomeState.FilterCursor > 0 {
					m.HomeState.FilterCursor--
				}
				return m, nil
			}
		case "down", "j":
			if m.HomeState.ActivePane == filterPane {
				if m.HomeState.FilterCursor < len(filterOptions)-1 {
					m.HomeState.FilterCursor++
				}
				return m, nil
			}
		case "enter":
			if m.screen == screenHome {
				if m.HomeState.ActivePane == filterPane {
					m.HomeState.ShowFilter = false
					m.HomeState.ActivePane = listPane
					m.HomeState.Issues.SetSize(leftWidth(m.width)-4, m.listInnerHeight())
				} else {
					m.HomeState.ActivePane = detailPane
				}
			}
			return m, nil
		case "esc":
			if m.HomeState.ShowFilter {
				m.HomeState.ShowFilter = false
				m.HomeState.ActivePane = listPane
				m.HomeState.Issues.SetSize(leftWidth(m.width)-4, m.listInnerHeight())
			}
			return m, nil
	}

	if  m.HomeState.ActivePane == listPane {
		var cmd tea.Cmd
		m.HomeState.Issues, cmd = m.HomeState.Issues.Update(msg)
		return m, cmd
	}

	return m, nil
}
