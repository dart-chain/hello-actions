package utils

import "testing"

func TestIsVersionValid(t *testing.T) {
	tests := []struct {
		version string
		valid   bool
	}{
		{"1.0.0", false},
		{"v1.2.3", true},
		{"2.0.0-alpha", false},
		{"3.1.4+build.123", false},
		{"v4.5.6-beta+exp.sha.5114f85", true},
		{"1", false},
		{"1.0", false},
		{"v1", false},
		{"version1.0.0", false},
		{"1.0.0-", false},
		{"1.0.0+", false},
	}

	for _, test := range tests {
		t.Run(test.version, func(t *testing.T) {
			if got := IsVersionValid(test.version); got != test.valid {
				t.Errorf("IsVersionValid(%q) = %v; want %v", test.version, got, test.valid)
			}
		})
	}
}

func TestExtractVersion(t *testing.T) {
	tests := []struct {
		version  string
		expected string
	}{
		{"1.0.0", ""},
		{"v1.2.3", "v1.2.3"},
		{"2.0.0-alpha", ""},
		{"3.1.4+build.123", ""},
		{"v4.5.6-beta+exp.sha.5114f85", "v4.5.6"},
		{"1", ""},
		{"1.0", ""},
		{"v1", ""},
		{"version1.0.0", ""},
	}

	for _, test := range tests {
		t.Run(test.version, func(t *testing.T) {
			if got := ExtractVersion(test.version); got != test.expected {
				t.Errorf("ExtractVersion(%q) = %v; want %v", test.version, got, test.expected)
			}
		})
	}
}
