// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"context"
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/downloader"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/registry"
)

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Level defines the plugin log level.
	Level string `envconfig:"PLUGIN_LOG_LEVEL"`

	RegistryUrl string `envconfig:"PLUGIN_REGISTRY_URL"`
	Username    string `envconfig:"PLUGIN_REGISTRY_USERNAME"`
	Password    string `envconfig:"PLUGIN_REGISTRY_PASSWORD"`
	ChartPath   string `envconfig:"PLUGIN_CHART_PATH"`
	Namespace   string `envconfig:"PLUGIN_REGISTRY_NAMESPACE"`
	ProjectId   string `envconfig:"PLUGIN_GCLOUD_PROJECT_ID"`
}

// Exec executes the plugin.
func Exec(ctx context.Context, args Args) error {
	if err := VerifyArgs(&args); err != nil {
		return err
	}

	packageRun, err := packageChart(&args)
	if err != nil {
		return err
	}

	opts := []registry.ClientOption{
		registry.ClientOptWriter(os.Stdout),
	}

	err = registryLogin(&args, opts)
	if err != nil {
		return err
	}

	err = registryPush(&args, opts, packageRun)
	if err != nil {
		return err
	}

	return nil
}

func VerifyArgs(args *Args) error {
	if args.RegistryUrl == "" {
		return fmt.Errorf("registry url is required")
	}
	if args.Username == "" {
		return fmt.Errorf("username is required")
	}
	if args.Password == "" {
		return fmt.Errorf("password is required")
	}
	if args.ChartPath == "" {
		return fmt.Errorf("chart path is required")
	}
	if args.Namespace == "" {
		return fmt.Errorf("namespace is required")
	}
	return nil
}

func packageChart(args *Args) (string, error) {
	helmClient := action.NewPackage()
	helmClient.DependencyUpdate = true
	helmClient.Destination = args.ChartPath

	settings := cli.New()
	getters := getter.All(settings)

	registryClient, err := registry.NewClient()

	if err != nil {
		return "", fmt.Errorf("failed to create registry client")
	}

	downloadManager := &downloader.Manager{
		Out:              os.Stdout,
		ChartPath:        args.ChartPath,
		Debug:            false,
		Getters:          getters,
		RepositoryConfig: settings.RepositoryConfig,
		RepositoryCache:  settings.RepositoryCache,
		RegistryClient:   registryClient,
	}

	if err := downloadManager.Build(); err != nil {
		return "", fmt.Errorf("failed to retrieve chart")
	}

	packageRun, err := helmClient.Run(args.ChartPath, nil)
	if err != nil {
		return "", fmt.Errorf("failed to package chart")
	}

	fmt.Print("Successfully packaged chart\n")

	return packageRun, nil
}

func registryLogin(args *Args, opts []registry.ClientOption) error {
	registryClient, err := registry.NewClient(opts...)
	if err != nil {
		return fmt.Errorf("failed to create registry client")
	}

	cfg := new(action.Configuration)
	cfg.RegistryClient = registryClient

	err = action.NewRegistryLogin(cfg).Run(
		os.Stdout,
		args.RegistryUrl,
		args.Username,
		args.Password,
	)

	if err != nil {
		return fmt.Errorf("failed to login to registry")
	}

	return nil
}

func registryPush(args *Args, opts []registry.ClientOption, packageRun string) error {
	registryClient, err := registry.NewClient(opts...)
	if err != nil {
		return fmt.Errorf("failed to create registry client")
	}

	cfg := new(action.Configuration)
	cfg.RegistryClient = registryClient

	client := action.NewPushWithOpts(action.WithPushConfig(cfg))

	settings := new(cli.EnvSettings)
	client.Settings = settings

	var remoteURL string

	if args.ProjectId != "" {
		remoteURL = fmt.Sprintf("oci://%s/%s/%s", args.RegistryUrl, args.ProjectId, args.Namespace)
	} else {
		remoteURL = fmt.Sprintf("oci://%s/%s", args.RegistryUrl, args.Namespace)
	}

	_, err = client.Run(packageRun, remoteURL)
	if err != nil {
		return fmt.Errorf("failed to push chart")
	}

	fmt.Print("Successfully pushed chart\n")

	return nil
}
