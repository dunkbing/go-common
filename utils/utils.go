package utils

import (
	"net/url"
	"strings"
)

// CleanURL cleans a URL by removing the query parameters and the fragment.
func CleanURL(inputURL string) (string, error) {
	// Parse the URL to ensure it's valid
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	// Remove the query parameters and the fragment
	parsedURL.RawQuery = ""
	parsedURL.Fragment = ""

	// Reconstruct the URL without the query and fragment
	cleanedURL := parsedURL.String()

	// Remove the trailing slash if it exists
	cleanedURL = strings.TrimSuffix(cleanedURL, "/")

	return cleanedURL, nil
}
