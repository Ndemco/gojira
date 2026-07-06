package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ndemco/gojira/jira"
)

func (m model) updateHome(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.HomeState.ActivePane {
	case filterPane:
		return updateFilterPane(m, msg)
	case listPane:
		return updateListPane(m, msg)
	case detailPane:
		return updateDetailPane(m, msg)
	}

	return m, nil
}

func updateFilterPane(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "f":
		m.HomeState.ShowFilter = false
		m.HomeState.ActivePane = listPane
	case "tab":
		m.HomeState.ActivePane = listPane
	case "up", "k":
		if m.HomeState.FilterCursor > 0 {
			m.HomeState.FilterCursor--
		}
	case "down", "j":
		if m.HomeState.FilterCursor < len(quickFilterOptions)-1 {
			m.HomeState.FilterCursor++
		}
	case "enter":
		var issues, err = jira.FetchIssues(quickFilterOptions[m.HomeState.FilterCursor].JQL)
		if err != nil {
			break
		}
		items := make([]list.Item, len(issues))
		for i, issue := range issues {
			items[i] = issue
		}
		m.HomeState.Issues.SetItems(items)
	case "esc":
		m.HomeState.ShowFilter = false
		m.HomeState.ActivePane = listPane
	}

	return m, nil
}

func updateListPane(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "f":
		m.HomeState.ShowFilter = true
		m.HomeState.ActivePane = filterPane
		m.HomeState.Issues.SetSize(leftWidth(m.width)-4, m.listInnerHeight())
		return m, nil
	case "tab", "enter":
		m.HomeState.ActivePane = detailPane
		return m, nil
	}

	var cmd tea.Cmd
	m.HomeState.Issues, cmd = m.HomeState.Issues.Update(msg)
	return m, cmd
}

func updateDetailPane(m model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "tab", "esc":
		m.HomeState.ActivePane = listPane
	}
	return m, nil
}
