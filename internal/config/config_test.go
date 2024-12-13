package config

import (
	"testing"
)

func TestConfigRead(t *testing.T) {
	dburl := "postgres://example"

	expectedConfig := Config{DBUrl: &dburl}

	var resultConfig Config

	err := Read(&resultConfig)

	if err != nil {
		print("Error: ", err)
	}

	if resultConfig.DBUrl == expectedConfig.DBUrl {
		t.Errorf("Structs are not equal: got %s, want %s", *resultConfig.DBUrl, *expectedConfig.DBUrl)
	}
}
