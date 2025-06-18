package util

import "net/url"

// Checks if a given URL is valid
func IsValidURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return false
	}

	u, err := url.Parse(urlString)
	return err == nil && u.Scheme != "" && u.Host != ""
}
