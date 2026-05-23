package jira

import "strings"

// adfNode is a node in Atlassian Document Format, used for description and comment bodies.
type adfNode struct {
	Type    string    `json:"type"`
	Text    string    `json:"text"`
	Content []adfNode `json:"content"`
}

func adfToText(node *adfNode) string {
	if node == nil {
		return ""
	}
	switch node.Type {
	case "text":
		return node.Text
	case "hardBreak":
		return "\n"
	}
	var parts []string
	for _, child := range node.Content {
		if t := adfToText(&child); t != "" {
			parts = append(parts, t)
		}
	}
	switch node.Type {
	case "doc", "bulletList", "orderedList":
		return strings.Join(parts, "\n\n")
	default:
		return strings.Join(parts, " ")
	}
}
