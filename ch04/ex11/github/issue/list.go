package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// List gets the list of issues from the specified repository
func List(owner, repo, state string) (*[]Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues?state=%s", baseURL, owner, repo, url.QueryEscape(state))
	resp, err := http.Get(url)
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
