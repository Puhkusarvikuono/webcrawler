package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s%s", strings.ToLower(parsedURL.Host), parsedURL.Path), nil
}
