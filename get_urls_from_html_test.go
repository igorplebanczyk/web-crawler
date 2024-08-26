package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "only absolute URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="https://example.com/path/one">
			<span>Example</span>
		</a>
		<a href="https://other.com/path/two">
			<span>Other</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://example.com/path/one", "https://other.com/path/two"},
		},
		{
			name:     "only relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Path One</span>
		</a>
		<a href="/path/two">
			<span>Path Two</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "no URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<p>No links here!</p>
	</body>
</html>
`,
			expected: []string{},
		},
		{
			name:     "URLs with fragments and queries",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one?query=1#fragment">
			<span>Path One</span>
		</a>
		<a href="https://example.com/path/two?query=2#fragment">
			<span>Path Two</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one?query=1#fragment", "https://example.com/path/two?query=2#fragment"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
