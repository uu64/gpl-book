package issue

import (
	"encoding/json"
	"fmt"
)

// Update updates an issue
func Update(owner, repo, number string, issue Issue) (*Issue, error) {
	// url := fmt.Sprintf("%s/repos/%s/%s/issues/%s", baseURL, owner, repo, number)
	body, _ := json.Marshal(issue)
	fmt.Println(body)

	// req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
	// if err != nil {
	// 	return nil, err
	// }

	// client := &http.Client{}
	// resp, err := client.Do(req)
	// if resp.StatusCode != http.StatusOK {
	// 	resp.Body.Close()
	// 	return nil, fmt.Errorf("update failed: %s", resp.Status)
	// }

	// var result Issue
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	resp.Body.Close()
	// 	return nil, err
	// }
	// resp.Body.Close()
	// return &result, nil
	return nil, nil
}
