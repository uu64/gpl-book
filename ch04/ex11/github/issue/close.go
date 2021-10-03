package issue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Close closes an issue
func Close(owner, repo, number string) (*Issue, error) {
	body, err := json.Marshal(struct {
		State string `json:"state"`
	}{State: "closed"})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", baseURL, owner, repo, number)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("close failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
