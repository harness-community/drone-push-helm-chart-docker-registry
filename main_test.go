package main

import (
	"os"
	"testing"
)

func TestMain_EnvVarsNotSet(t *testing.T) {
	// Save current environment variables
	originalRegistryUrl := os.Getenv("PLUGIN_REGISTRY_URL")
	originalUsername := os.Getenv("PLUGIN_REGISTRY_USERNAME")
	originalToken := os.Getenv("PLUGIN_REGISTRY_PASSWORD")
	originalChartPath := os.Getenv("PLUGIN_CHART_PATH")
	originalNamespace := os.Getenv("PLUGIN_REGISTRY_NAMESPACE")

	// Clear environment variables
	os.Setenv("PLUGIN_REGISTRY_URL", "")
	os.Setenv("PLUGIN_REGISTRY_USERNAME", "")
	os.Setenv("PLUGIN_REGISTRY_PASSWORD", "")
	os.Setenv("PLUGIN_CHART_PATH", "")
	os.Setenv("PLUGIN_REGISTRY_NAMESPACE", "")

	defer func() {
		// Restore original environment variables
		os.Setenv("PLUGIN_REGISTRY_URL", originalRegistryUrl)
		os.Setenv("PLUGIN_REGISTRY_USERNAME", originalUsername)
		os.Setenv("PLUGIN_REGISTRY_PASSWORD", originalToken)
		os.Setenv("PLUGIN_CHART_PATH", originalChartPath)
		os.Setenv("PLUGIN_REGISTRY_NAMESPACE", originalNamespace)
	}()

	err := verifyEnvVars()
	if err == nil {
		t.Error("Expected error, but got nil")
	}

}

func TestMain_EnvVarsSet(t *testing.T) {
	// Save current environment variables
	originalRegistryUrl := os.Getenv("PLUGIN_REGISTRY_URL")
	originalUsername := os.Getenv("PLUGIN_REGISTRY_USERNAME")
	originalToken := os.Getenv("PLUGIN_REGISTRY_PASSWORD")
	originalChartPath := os.Getenv("PLUGIN_CHART_PATH")
	originalNamespace := os.Getenv("PLUGIN_REGISTRY_NAMESPACE")

	// Set environment variables
	os.Setenv("PLUGIN_REGISTRY_URL", "https://registry.example.com")
	os.Setenv("PLUGIN_REGISTRY_USERNAME", "username")
	os.Setenv("PLUGIN_REGISTRY_PASSWORD", "token")
	os.Setenv("PLUGIN_CHART_PATH", "charts")
	os.Setenv("PLUGIN_REGISTRY_NAMESPACE", "namespace")

	defer func() {
		// Restore original environment variables
		os.Setenv("PLUGIN_REGISTRY_URL", originalRegistryUrl)
		os.Setenv("PLUGIN_REGISTRY_USERNAME", originalUsername)
		os.Setenv("PLUGIN_REGISTRY_PASSWORD", originalToken)
		os.Setenv("PLUGIN_CHART_PATH", originalChartPath)
		os.Setenv("PLUGIN_REGISTRY_NAMESPACE", originalNamespace)
	}()

	err := verifyEnvVars()
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}
}
