package git

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func createRepoUrl() string {
	return fmt.Sprintf("https://api.github.com/orgs/%s/repos", os.Getenv("GITHUB_ORGANISATION"))
}

func CreateRepo(name string, is_private bool) error {
	payload, err := json.Marshal(map[string]interface{}{
		"name":    name,
		"private": is_private,
	})
	if err != nil {
		return err
	}

	var url string = createRepoUrl()
	res, err := sendRequest(http.MethodPost, url, payload)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	return nil
}
