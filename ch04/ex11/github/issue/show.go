package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Show shows the detail of the specified issue
func Show(owner, repo, number string) (*Issue, error) {
	resp, err := http.Get(strings.Join([]string{
		baseURL, "repos", owner, repo, "issues", number}, "/"))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("show failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
