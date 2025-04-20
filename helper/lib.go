package helper

import (
	"net/url"
)

func IsValidURL(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	return err == nil && u.Scheme != "" && u.Host != ""
}
