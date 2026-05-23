package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ndemco/gojira/jira"
)

func main() {
	items := []list.Item{
		jira.Issue{
			Key: "GOJ-1", Summary: "Set up CI pipeline",
			Status: "Done", Assignee: "alice",
			Body: "Configure GitHub Actions to run tests and lint on every PR. Include build caching to keep runs fast.",
		},
		jira.Issue{
			Key: "GOJ-2", Summary: "Implement auth token refresh",
			Status: "In Progress", Assignee: "bob",
			Body: "Access tokens expire after 1h. Add background refresh logic so the user isn't booted mid-session.",
		},
		jira.Issue{
			Key: "GOJ-3", Summary: "Fix pagination on issue list",
			Status: "To Do", Assignee: "alice",
			Body: "The list only loads the first 50 issues. Wire up cursor-based pagination to load more as the user scrolls.",
		},
		jira.Issue{
			Key: "GOJ-4", Summary: "Add label filtering",
			Status: "To Do", Assignee: "carol",
			Body: "Users want to filter the list by label. Add a filter bar at the top with multi-select label chips.",
		},
		jira.Issue{
			Key: "GOJ-5", Summary: "Dark mode support",
			Status: "In Progress", Assignee: "bob",
			Body: "Detect terminal background color and switch between light/dark palettes automatically.",
		},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Issues"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	p := tea.NewProgram(model{list: l}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
