package main

import (
	"os"
	"testing"
)

func TestEnvironmentVariables(t *testing.T) {
	os.Setenv("PLUGIN_REGISTRY_URL", "")
	// Set other required env variables to non-empty values

	if os.Getenv("PLUGIN_REGISTRY_URL") == "" {
		t.Error("Expected an error due to missing PLUGIN_REGISTRY_URL")
	}
	// Repeat for other variables
}
