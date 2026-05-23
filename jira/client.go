package jira

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

type searchResponse struct {
	Issues []Issue `json:"issues"`
}

func FetchIssues() ([]Issue, error) {
	request, err := newRequest("GET", "stubbedUrl", nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result searchResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Issues, nil
}

func newRequest(method, url string, body any) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(data)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("", "")
	req.SetBasicAuth("", "")
	req.SetBasicAuth("", "")

	return req, nil
}
