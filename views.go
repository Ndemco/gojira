package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/ndemco/gojira/jira"
)

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	switch m.screen {
	case screenSettings:
		return m.settingsView()
	default:
		return m.homeView()
	}
}

func (m model) homeView() string {
	lw := leftWidth(m.width)
	rw := m.width - lw - 6
	h := m.height - 4

	var left string
	if m.HomeState.ShowFilter {
		filterStyle := inactiveBox
		if m.HomeState.ActivePane == filterPane {
			filterStyle = activeBox
		}
		listStyle := inactiveBox
		if m.HomeState.ActivePane == listPane {
			listStyle = activeBox
		}
		filterBox := filterStyle.Width(lw).Height(filterInnerHeight).Render(m.filterView())
		listBox := listStyle.Width(lw).Height(h - filterInnerHeight - 2).Render(m.HomeState.Issues.View())
		left = lipgloss.JoinVertical(lipgloss.Left, filterBox, listBox)
	} else {
		listStyle := inactiveBox
		if m.HomeState.ActivePane == listPane {
			listStyle = activeBox
		}
		left = listStyle.Width(lw).Height(h).Render(m.HomeState.Issues.View())
	}

	rightStyle := inactiveBox
	if m.HomeState.ActivePane == detailPane {
		rightStyle = activeBox
	}
	right := rightStyle.Width(rw).Height(h).Render(m.detailView(rw))

	help := helpStyle.Render("↑/↓ navigate  enter/tab switch pane  f filter  esc close  1 home  2 settings  q quit")
	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, left, right),
		help,
	)
}

func (m model) filterView() string {
	var sb strings.Builder
	sb.WriteString(keyStyle.Render("Filter"))
	sb.WriteString("\n\n")
	for i, opt := range quickFilterOptions {
		if i == m.HomeState.FilterCursor {
			sb.WriteString(keyStyle.Render("▸ " + opt.Name))
		} else {
			sb.WriteString(labelStyle.Render("  " + opt.Name))
		}
		if i < len(quickFilterOptions)-1 {
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

func (m model) settingsView() string {
	h := m.height - 4
	content := fmt.Sprintf("%s\n\n%s\n%s\n%s",
		keyStyle.Render("Settings"),
		labelStyle.Render("Theme:    Default"),
		labelStyle.Render("API URL:  https://jira.example.com"),
		labelStyle.Render("Username: alice"),
	)
	box := activeBox.Width(m.width - 4).Height(h).Render(content)
	help := helpStyle.Render("1 home  2 settings  q quit")
	return lipgloss.JoinVertical(lipgloss.Left, box, help)
}

func (m model) detailView(width int) string {
	issue, ok := m.HomeState.Issues.SelectedItem().(jira.Issue)
	if !ok {
		return labelStyle.Render("Select an issue to view details.")
	}

	statusColor := map[string]string{
		"Done":        "34",
		"In Progress": "214",
		"To Do":       "240",
	}
	color, ok := statusColor[issue.Status]
	if !ok {
		color = "240"
	}
	badge := lipgloss.NewStyle().
		Background(lipgloss.Color(color)).
		Foreground(lipgloss.Color("0")).
		Padding(0, 1).
		Render(issue.Status)

	commentsz := buildCommentsString(issue.Comments)

	return fmt.Sprintf("%s\n\n%s %s\n\n%s %s\n\n%s\n%s\n\n%s\n\n%s",
		keyStyle.Render(issue.Key+" · "+issue.Summary),
		labelStyle.Render("Status:  "), badge,
		labelStyle.Render("Assignee:"), issue.Assignee,
		labelStyle.Render("Body:"), issue.Body,
		labelStyle.Render("Comments"), commentsz,
	)
}

func buildCommentsString(comments []jira.Comment) string {
	var cb strings.Builder
	for i, c := range comments {
		fmt.Fprintf(&cb, "%s\n%s", authorStyle.Render(c.Author+":"), c.Body)
		if i < len(comments)-1 {
			cb.WriteString("\n\n")
		}
	}

	return cb.String()
}
