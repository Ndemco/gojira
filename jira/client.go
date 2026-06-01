package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

var client = &http.Client {
	Timeout: 10 * time.Second,
}

type searchResponse struct {
	Issues []issueJSON `json:"issues"`
}

type issueJSON struct {
	Key    string      `json:"key"`
	Fields issueFields `json:"fields"`
}

type issueFields struct {
	Summary     string      `json:"summary"`
	Status      statusJSON  `json:"status"`
	Assignee    *userJSON   `json:"assignee"`
	Description *adfNode    `json:"description"`
	Comment     commentPage `json:"comment"`
}

type statusJSON struct {
	Name string `json:"name"`
}

type userJSON struct {
	DisplayName string `json:"displayName"`
}

type commentPage struct {
	Comments []commentJSON `json:"comments"`
}

type commentJSON struct {
	Author userJSON `json:"author"`
	Body   *adfNode `json:"body"`
}

func FetchIssues(jql string) ([]Issue, error) {
	base := os.Getenv("JIRA_BASE_URL")
	url, err := url.Parse(base + "/rest/api/3/search/jql")
	if err != nil {
		return nil, err
	}
	query := url.Query()
	query.Set("jql", jql)
	query.Set("fields", "summary,status,assignee,description,comment")
	query.Set("maxResults", "50")
	url.RawQuery = query.Encode()

	req, err := buildRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("jira: %d %s", resp.StatusCode, body)
	}

	var result searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	issues := make([]Issue, len(result.Issues))
	for i, ij := range result.Issues {
		assignee := ""
		if ij.Fields.Assignee != nil {
			assignee = ij.Fields.Assignee.DisplayName
		}

		comments := make([]Comment, len(ij.Fields.Comment.Comments))
		for j, c := range ij.Fields.Comment.Comments {
			comments[j] = Comment{
				Author: c.Author.DisplayName,
				Body:   adfToText(c.Body),
			}
		}

		issues[i] = Issue{
			Key:      ij.Key,
			Summary:  ij.Fields.Summary,
			Status:   ij.Fields.Status.Name,
			Assignee: assignee,
			Body:     adfToText(ij.Fields.Description),
			Comments: comments,
		}
	}

	return issues, nil
}

func buildRequest(method, endpoint string, body any) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(os.Getenv("JIRA_EMAIL"), os.Getenv("JIRA_API_TOKEN"))

	return req, nil
}
