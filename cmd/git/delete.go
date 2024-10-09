package git

import (
	"fmt"
	"net/http"
	"os"
)

func deleteRepoUrl(name string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s", os.Getenv("GITHUB_ORGANISATION"), name)
}

func DeleteRepo(name string) error {
	url := deleteRepoUrl(name)
	res, err := sendRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}
	return nil
}
