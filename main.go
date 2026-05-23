package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/ndemco/gojira/jira"
)

func main() {
	godotenv.Load()

	issues, err := jira.FetchIssues()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	items := make([]list.Item, len(issues))
	for i, issue := range issues {
		items[i] = issue
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
