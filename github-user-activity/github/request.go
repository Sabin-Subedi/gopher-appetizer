package github

import (
	"io"
	"net/http"
)

func requestGithub(url string) []byte {
	resp, error := http.Get(
		"https://api.github.com/" + url,
	)

	if error != nil {
		panic(error)
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(error)
	}

	return body
}
