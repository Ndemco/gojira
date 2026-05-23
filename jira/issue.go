package jira

import "fmt"

type Issue struct {
	Key      string
	Summary  string
	Status   string
	Assignee string
	Body     string
	Comments []Comment
}

type Comment struct {
	Author string
	Body   string
}

func (i Issue) Title() string       { return fmt.Sprintf("[%s] %s", i.Key, i.Summary) }
func (i Issue) Description() string { return fmt.Sprintf("%s · %s", i.Status, i.Assignee) }
func (i Issue) FilterValue() string { return i.Summary }
