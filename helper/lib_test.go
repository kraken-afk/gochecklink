package helper_test

import (
	"gochecklink/helper"
	"testing"
)

func TestIsURLValid(t *testing.T) {
	endpoints := map[string]bool{
		"https://www.example.com":               true,
		"http://localhost:8080":                 true,
		"ftp://user:password@host.com/file.txt": true,
		"mailto:john.doe@example.com":           false,
		"/api/users":                            false,
		"":                                      false,
		"ht tp://oops.com":                      false,
		"://missing-scheme":                     false,
		"news://news.example.com":               true,
	}

	for endpoint, isValid := range endpoints {
		if helper.IsValidURL(endpoint) != isValid {
			t.Errorf("[ERROR]: '%s' should be %t", endpoint, isValid)
		}
	}
}
