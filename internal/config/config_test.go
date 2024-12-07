package config

import (
	//"reflect"
	"reflect"
	"testing"
)

func TestConfigRead(t *testing.T) {
	dburl := "postgres://example"
	expectedConfig := Config{DBUrl: &dburl}

	resultConfig, err := Read()

	if err != nil {
		print("Error: ", err)
	}

	if *resultConfig.DBUrl != *expectedConfig.DBUrl {
		t.Errorf("Structs are not equal: got %s, want %s", *resultConfig.DBUrl, *expectedConfig.DBUrl)
	}
}

func TestConfigComparison(t *testing.T) {
	// Helper to create string pointers
	strPtr := func(s string) *string {
		return &s
	}

	// Test cases
	tests := []struct {
		name     string
		config1  Config
		config2  Config
		expected bool
	}{
		{
			name:     "Equal Configs",
			config1:  Config{DBUrl: strPtr("postgres://db1"), CurrentUserName: strPtr("user1")},
			config2:  Config{DBUrl: strPtr("postgres://db1"), CurrentUserName: strPtr("user1")},
			expected: true,
		},
		{
			name:     "Different DBUrls",
			config1:  Config{DBUrl: strPtr("postgres://db1"), CurrentUserName: strPtr("user1")},
			config2:  Config{DBUrl: strPtr("postgres://db2"), CurrentUserName: strPtr("user1")},
			expected: false,
		},
		{
			name:     "Different CurrentUserName",
			config1:  Config{DBUrl: strPtr("postgres://db1"), CurrentUserName: strPtr("user1")},
			config2:  Config{DBUrl: strPtr("postgres://db1"), CurrentUserName: strPtr("user2")},
			expected: false,
		},
		{
			name:     "Nil Fields",
			config1:  Config{DBUrl: nil, CurrentUserName: strPtr("user1")},
			config2:  Config{DBUrl: nil, CurrentUserName: strPtr("user1")},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use reflect.DeepEqual for comparison
			result := reflect.DeepEqual(tt.config1, tt.config2)
			if result != tt.expected {
				t.Errorf("Config comparison failed for %s: got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}
