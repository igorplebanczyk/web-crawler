package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove scheme with port",
			inputURL: "https://blog.boot.dev:8080/path",
			expected: "blog.boot.dev:8080/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://example.com",
			expected: "example.com",
		},
		{
			name:     "remove https scheme",
			inputURL: "https://example.com",
			expected: "example.com",
		},
		{
			name:     "no scheme",
			inputURL: "example.com",
			expected: "example.com",
		},
		{
			name:     "empty URL",
			inputURL: "",
			expected: "",
		},
		{
			name:     "URL with query parameters",
			inputURL: "https://example.com/path?query=1",
			expected: "example.com/path?query=1",
		},
		{
			name:     "URL with fragment",
			inputURL: "https://example.com/path#fragment",
			expected: "example.com/path#fragment",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
