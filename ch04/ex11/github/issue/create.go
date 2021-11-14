package issue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

// Create create an issue
func Create(owner, repo string, title, body []byte) (*Issue, error) {
	body, err := json.Marshal(struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{string(title), string(body)})
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/repos/%s/%s/issues", baseURL, owner, repo)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	client := oauth2.NewClient(context.Background(),
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: os.Getenv("GH_ACCESS_TOKEN")}))
	resp, _ := client.Do(req)
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return nil, fmt.Errorf("create failed: %s", resp.Status)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
