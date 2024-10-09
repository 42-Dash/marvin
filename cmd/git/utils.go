package git

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func sendRequest(method string, url string, payload []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("GITHUB_ACCESS")))
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Content-Type", "application/json")

	return client.Do(req)
}
