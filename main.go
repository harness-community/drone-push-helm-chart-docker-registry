package main

import (
	"drone/plugin/helm-chart-docker-registry/env"
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/registry"
)

func main() {
	registryUrl := os.Getenv("PLUGIN_REGISTRY_URL")
	username := os.Getenv("PLUGIN_REGISTRY_USERNAME")
	token := os.Getenv("PLUGIN_REGISTRY_PASSWORD")
	chartPath := os.Getenv("PLUGIN_CHART_PATH")
	namespace := os.Getenv("PLUGIN_REGISTRY_NAMESPACE")
	projectId := os.Getenv("PLUGIN_GCLOUD_PROJECT_ID")

	err := env.VerifyEnvVars()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// package chart
	helmClient := action.NewPackage()
	helmClient.DependencyUpdate = true
	helmClient.Destination = chartPath

	downloadManager := &downloader.Manager{
		Out:       os.Stdout,
		ChartPath: chartPath,
		Debug:     true,
	}

	if err := downloadManager.Build(); err != nil {
		fmt.Printf("Failed to retrieve chart in %s", chartPath)
		os.Exit(1)
	}

	packageRun, err := helmClient.Run(chartPath, nil)
	if err != nil {
		fmt.Printf("Failed to package chart in %s", chartPath)
		os.Exit(1)
	}

	fmt.Printf("Successfully packaged chart in %s\n", chartPath)

	opts := []registry.ClientOption{
		registry.ClientOptWriter(os.Stdout),
	}

	registryClient, err := registry.NewClient(opts...)
	if err != nil {
		fmt.Println("Failed to create registry client")
		os.Exit(1)
	}

	cfg := new(action.Configuration)
	cfg.RegistryClient = registryClient

	action.NewRegistryLogin(cfg).Run(
		os.Stdout,
		registryUrl,
		username,
		token,
	)

	// Push the chart
	client := action.NewPushWithOpts(action.WithPushConfig(cfg))

	settings := new(cli.EnvSettings)
	client.Settings = settings

	var ociURL string

	if projectId != "" {
		ociURL = "oci://" + registryUrl + "/" + projectId + "/" + namespace
	} else {
		ociURL = "oci://" + registryUrl + "/" + namespace
	}

	_, err = client.Run(packageRun, ociURL)
	if err != nil {
		fmt.Println("Failed to push chart")
		os.Exit(1)
	}

	fmt.Printf("Successfully pushed chart to %s", ociURL)
}
