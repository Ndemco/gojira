package main

import "github.com/charmbracelet/lipgloss"

var (
	activeBox   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62"))
	inactiveBox = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("240"))
	keyStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("62"))
	labelStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	authorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("214"))
	helpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Padding(0, 1)
)
