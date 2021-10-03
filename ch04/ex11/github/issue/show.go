package issue

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Show shows the detail of the specified issue
func Show(owner, repo, number string) (*Issue, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", baseURL, owner, repo, number)
	resp, err := http.Get(url)
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
