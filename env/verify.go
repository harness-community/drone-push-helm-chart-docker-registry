package env

import (
	"fmt"
	"os"
)

func VerifyEnvVars() error {
	registryUrl := os.Getenv("PLUGIN_REGISTRY_URL")
	username := os.Getenv("PLUGIN_REGISTRY_USERNAME")
	token := os.Getenv("PLUGIN_REGISTRY_PASSWORD")
	chartPath := os.Getenv("PLUGIN_CHART_PATH")
	namespace := os.Getenv("PLUGIN_REGISTRY_NAMESPACE")

	if (registryUrl == "") || (username == "") || (token == "") || (namespace == "") || (chartPath == "") {
		return fmt.Errorf("required environment variables not set")
	}

	return nil
}
