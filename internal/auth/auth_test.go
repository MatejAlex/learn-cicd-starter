package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	cases := []struct {
		input    http.Header
		expected string
	}{
		{
			input:    http.Header{"Authorization": []string{"ApiKey test1"}},
			expected: "test1",
		},
		{
			input: http.Header{},
		},
		{
			input:    http.Header{"User-Agent": []string{"matejalex"}},
			expected: "",
		},
		{
			input:    http.Header{"User-Agent": []string{"matejalex"}, "Authorization": []string{"ApiKey test1"}},
			expected: "test1",
		},
	}

	for _, c := range cases {
		actual, _ := GetAPIKey(c.input)
		if actual != c.expected {
			t.Errorf("Results do not match, input: %+v [returned:expected] %s:%s", c.input, actual, c.expected)
		}
	}

}
