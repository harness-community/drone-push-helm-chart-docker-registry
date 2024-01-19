# drone-helm-chart-container-registry

- [Synopsis](#Synopsis)
- [Parameters](#Paramaters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin is designed to streamline the packaging and distribution of Helm charts to a Container registry.

Currently, supports pushing Helm charts to Docker Hub and Google Artifact Registry.

To learn how to utilize Drone plugins in Harness CI, please consult the provided [documentation](https://developer.harness.io/docs/continuous-integration/use-ci/use-drone-plugins/run-a-drone-plugin-in-ci).

## Parameters

| Parameter                                                                                                                        | Choices/<span style="color:blue;">Defaults</span> | Comments                                                   |
| :------------------------------------------------------------------------------------------------------------------------------- | :------------------------------------------------ | :--------------------------------------------------------- |
| registry_url <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>       |                                                   | Docker registry where the packaged chart will be published |
| chart_path <span style="font-size: 10px"><br/>`string`</span>                                                                    | Defaults: <span style="color:blue;">`./`</span>   | Directory containing the helm chart                        |
| registry_username <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>  |                                                   | Username to login to the above registry.                   |
| registry_password <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span>  |                                                   | PAT / access token to authenticate                         |
| registry_namespace <span style="font-size: 10px"><br/>`string`</span> <span style="color:red; font-size: 10px">`required`</span> |                                                   | Namespace under which the chart will be published          |

## Notes

If you aim to push Helm Charts to Google Artifact Registry (GAR):

- Set the registry URL to `LOCATION-docker.pkg.dev`
- use `oauth2accesstoken` as username and `access-token` as token. Refer to this documentation for generating an access token.
- use `REPO_ID` as `registry_namespace` and `PROJECT_ID` as `gcloud_project_id`

In case of Docker Hub:

- Set the registry URL to registry.hub.docker.com

Review the "[plugin image](#plugin-image)" section to identify the available tags for supported architectures, and then use these tags in the Docker Image during the plugin step.

For more details check the [examples](#Examples) section.

## Plugin Image

The plugin `harnesscommunity/drone-helm-chart-container-registry` is available for the following architectures:

| OS            | Tag             |
| ------------- | --------------- |
| linux/amd64   | `linux-amd64`   |
| linux/arm64   | `linux-arm64`   |
| windows/amd64 | `windows-amd64` |

## Examples

```
# Plugin YAML
# DockerHub Example
- step:
    type: Plugin
    name: Push Helm Chart to DockerHub
    identifier: helm_chart_docker
    spec:
        connectorRef: harness-docker-connector
        image: harnesscommunity/drone-helm-chart-container-registry:linux-amd64
        settings:
            registry_url: registry.hub.docker.com
            registry_username: <+variable.docker_username>
            registry_password: <+secrets.getValue("docker_pat")>
            chart_path: chart
            registry_namespace: <+variable.namespace>

- step:
    type: Plugin
    name: Push Helm Chart to GAR
    identifier: helm_chart_gar
    spec:
        connectorRef: harness-docker-connector
        image: harnesscommunity/drone-helm-chart-container-registry:linux-amd64
        settings:
            registry_url: LOCATION-docker.pkg.dev
            registry_username: oauth2accesstoken
            registry_password: <+secrets.getValue("access_token")>
            chart_path: chart
            gcloud_project_id: PROJECT_ID
            registry_namespace: REPO_ID

```

> <span style="font-size: 14px; margin-left:5px; background-color: #d3d3d3; padding: 4px; border-radius: 4px;">ℹ️ If you notice any issues in this documentation, you can [edit this document](https://github.com/harness-community/drone-push-helm-chart-docker-registry/blob/main/README.md) to improve it.</span>
