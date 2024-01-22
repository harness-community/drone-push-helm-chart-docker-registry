// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"testing"

	"helm.sh/helm/v3/pkg/registry"
)

func TestVerifyArgs(t *testing.T) {
	var err error

	err = VerifyArgs(&Args{
		RegistryUrl: "https://registry.hub.docker.com",
	})

	if err == nil {
		t.Error(err)
	}

	err = VerifyArgs(&Args{
		Username: "octocat",
		Password: "pass",
	})

	if err == nil {
		t.Error(err)
	}

	err = VerifyArgs(&Args{
		RegistryUrl: "https://registry.hub.docker.com",
		Username:    "octocat",
		Password:    "pass",
	})

	if err == nil {
		t.Error(err)
	}
}

func TestPackageChart(t *testing.T) {
	var err error

	_, err = packageChart(&Args{
		ChartPath: "test-chart",
	})

	if err != nil {
		t.Error(err)
	}

	_, err = packageChart(&Args{
		ChartPath: "test-chart-fail",
	})

	if err == nil {
		t.Error(err)
	}
}

func TestRegistryLogin(t *testing.T) {

	err := registryLogin(&Args{
		RegistryUrl: "https://registry.hub.docker.com",
		Username:    "octocat",
		Password:    "pass",
	}, []registry.ClientOption{})

	if err == nil {
		t.Error(err)
	}
}

func TestRegistryPush(t *testing.T) {

	err := registryPush(&Args{
		RegistryUrl: "https://registry.hub.docker.com",
		Username:    "octocat",
		Password:    "pass",
	}, []registry.ClientOption{}, "test-chart")

	if err == nil {
		t.Error(err)
	}
}
