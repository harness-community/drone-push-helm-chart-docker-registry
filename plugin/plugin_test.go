// Copyright 2020 the Drone Authors. All rights reserved.
// Use of this source code is governed by the Blue Oak Model License
// that can be found in the LICENSE file.

package plugin

import (
	"path/filepath"
	"testing"

	"helm.sh/helm/v3/pkg/registry"
)

func TestVerifyArgs(t *testing.T) {
	var err error

	err = verifyArgs(&Args{
		RegistryUrl: "https://registry.hub.docker.com",
	})

	if err == nil {
		t.Error(err)
	}

	err = verifyArgs(&Args{
		Username: "octocat",
		Password: "pass",
	})

	if err == nil {
		t.Error(err)
	}

	err = verifyArgs(&Args{
		RegistryUrl: "https://registry.hub.docker.com",
		Username:    "octocat",
		Password:    "pass",
	})

	if err == nil {
		t.Error(err)
	}

	err = verifyArgs(&Args{
		RegistryUrl: "https://registry.hub.docker.com",
		Username:    "octocat",
		Password:    "pass",
		ChartPath:   "test-chart",
		Namespace:   "test",
	})

	if err != nil {
		t.Error(err)
	}
}

func TestPackageChart(t *testing.T) {
	tests := []struct {
		name      string
		chartPath string
		chartDest string
		wantErr   bool
	}{
		{
			name:      "test-chart-pass",
			chartPath: "test-chart/chart-pass",
			chartDest: "mywebapp-5.0.0.tgz",
			wantErr:   false,
		},
		{
			name:      "test-chart-bad",
			chartPath: "test-chart/bad-chart",
			wantErr:   true,
		},
		{
			name:      "test-chart-bad-dependency",
			chartPath: "test-chart/bad-dependency",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		tempDir := t.TempDir()

		args := &Args{
			ChartPath:        tt.chartPath,
			ChartDestination: tt.chartDest,
		}

		got, err := packageChart(args)

		got = filepath.Base(got)

		if err != nil {
			if tt.wantErr {
				return
			}
			t.Errorf("packageChart() error = %v, wantErr %v", err, tt.wantErr)
		}

		want := filepath.Join(tempDir, tt.chartDest)

		want = filepath.Base(want)

		if got != want {
			t.Errorf("packageChart() = %v, want %v", got, want)
		}
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
