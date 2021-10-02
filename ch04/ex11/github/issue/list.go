package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// List gets the list of issues from the specified repository
func List(owner, repo string) (*[]Issue, error) {
	resp, err := http.Get(strings.Join([]string{baseURL, "repos", owner, repo, "issues"}, "/"))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("list failed: %s", resp.Status)
	}

	var result []Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}
