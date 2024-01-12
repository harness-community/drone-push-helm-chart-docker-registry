package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	//  get environment variables
	registry := os.Getenv("PLUGIN_REGISTRY_URL")
	username := os.Getenv("PLUGIN_REGISTRY_USERNAME")
	token := os.Getenv("PLUGIN_REGISTRY_PASSWORD")
	chartPath := os.Getenv("PLUGIN_CHART_PATH")
	namespace := os.Getenv("PLUGIN_REGISTRY_NAMESPACE")

	if (registry == "") || (username == "") || (token == "") || (namespace == "") {
		fmt.Println("Missing required environment variables")
		os.Exit(1)
	}

	if chartPath == "" {
		chartPath = "./"
	}

	// get chart name
	chartName, err := getChartName(chartPath)
	if err != nil {
		fmt.Println("Failed to get chart name")
		os.Exit(1)
	}

	chartName = strings.TrimSpace(chartName)

	// get chart version
	chartVersion, err := getChartVersion(chartPath)
	if err != nil {
		fmt.Println("Failed to get chart version")
		os.Exit(1)
	}

	chartVersion = strings.TrimSpace(chartVersion)

	// cd into chart path
	if err := os.Chdir(chartPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// package chart
	packageChartCmd := exec.Command("helm", "package", "--dependency-update", ".")
	if err := packageChartCmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// get chart tarball name
	chartTarballName := fmt.Sprintf("%s-%s.tgz", chartName, chartVersion)

	fmt.Println("Chart packaged successfully - ", chartTarballName)

	// login to registry
	loginCmd := exec.Command("helm", "registry", "login", registry, "--username", username, "--password", token)
	if err := loginCmd.Run(); err != nil {
		fmt.Println("Failed to login to registry")
		os.Exit(1)
	}

	fmt.Println("Logged in to registry successfully")

	// push chart
	pushChartCmd := exec.Command("helm", "push", chartTarballName, fmt.Sprintf("oci://%s/%s", registry, namespace))
	if err := pushChartCmd.Run(); err != nil {
		fmt.Println("Failed to push chart")
		os.Exit(1)
	}

	fmt.Println("Chart pushed successfully - ", registry, "/", namespace, "/", chartName)
}

func getChartName(chartPath string) (string, error) {
	// helm show chart chartPath | grep name | awk '{print $2}'
	cmd1 := exec.Command("helm", "show", "chart", chartPath)
	cmd2 := exec.Command("awk", "/name:/ {print $2}")

	// if os is windows
	if runtime.GOOS == "windows" {
		// cmd = helm show chart ./ | Select-String 'name: (.*)' | ForEach-Object { $_.Matches.Groups[1].Value }
		cmd := exec.Command("powershell", "-Command", "helm show chart ", chartPath, "| Select-String 'name: (.*)' | ForEach-Object { $_.Matches.Groups[1].Value }")

		var out bytes.Buffer

		cmd.Stdout = &out

		if err := cmd.Run(); err != nil {
			fmt.Println("Chart does not exist in the specified path")
			return "", err
		}

		return out.String(), nil
	}

	var out1 bytes.Buffer
	var out2 bytes.Buffer

	cmd1.Stdout = &out1
	cmd2.Stdin = &out1
	cmd2.Stdout = &out2

	if err := cmd1.Run(); err != nil {
		fmt.Println("Chart does not exist in the specified path")
		return "", err
	}
	if err := cmd2.Run(); err != nil {
		return "", err
	}

	return out2.String(), nil
}

func getChartVersion(chartPath string) (string, error) {
	// helm show chart chartPath | grep version | awk '{print $2}'
	cmd1 := exec.Command("helm", "show", "chart", chartPath)
	cmd2 := exec.Command("awk", "/version:/ {print $2}")

	// if os is windows
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell", "-Command", "helm show chart ", chartPath, "| Select-String '^version: (.*)' | ForEach-Object { $_.Matches.Groups[1].Value }")

		var out bytes.Buffer

		cmd.Stdout = &out

		if err := cmd.Run(); err != nil {
			fmt.Println("Chart does not exist in the specified path")
			return "", err
		}

		return out.String(), nil
	}

	var out1 bytes.Buffer
	var out2 bytes.Buffer

	cmd1.Stdout = &out1
	cmd2.Stdin = &out1
	cmd2.Stdout = &out2

	if err := cmd1.Run(); err != nil {
		fmt.Println("Chart does not exist in the specified path")
		return "", err
	}
	if err := cmd2.Run(); err != nil {
		return "", err
	}

	return out2.String(), nil
}
