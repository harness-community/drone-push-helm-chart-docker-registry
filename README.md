# drone-helm-chart-docker-registry

- [Synopsis](#Synopsis)
- [Parameters](#Paramaters)
- [Notes](#Notes)
- [Plugin Image](#Plugin-Image)
- [Examples](#Examples)

## Synopsis

This plugin is designed to streamline the packaging and distribution of Helm charts to a Docker registry.

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

If you're using GAR, use `oauth2accesstoken` as username and `access-token` as token. Refer to this [documentation](https://cloud.google.com/artifact-registry/docs/helm/authentication#token) for generating an access token.

## Plugin Image

The plugin `harnesscommunity/drone-helm-chart-docker-registry` is available for the following architectures:

| OS            | Tag             |
| ------------- | --------------- |
| linux/amd64   | `linux-amd64`   |
| linux/arm64   | `linux-arm64`   |
| windows/amd64 | `windows-amd64` |

## Examples

```
# Plugin YAML
- step:
    type: Plugin
    name: Push Helm to Docker
    identifier: Push_Helm_to_Docker
    spec:
        connectorRef: harness-docker-connector
        image: harnesscommunity/drone-helm-chart-docker-registry
        settings:
            registry_url: registry.hub.docker.com
            registry_username: <+variable.docker_username>
            registry_password: <+secrets.getValue("docker_pat")>
            chart_path: chart
            docker_namespace: <+variable.namespace>

# Using GAR
- step:
    type: Plugin
    name: Push Helm to GAR
    identifier: Push_Helm_to_GAR
    spec:
        connectorRef: harness-docker-connector
        image: harnesscommunity/drone-helm-chart-docker-registry
        settings:
            registry_url: https://LOCATION-docker.pkg.dev
            registry_username: oauth2accesstoken
            registry_password: <+secrets.getValue("access_token")>
            chart_path: chart
            docker_namespace: <+variable.namespace>
```

> <span style="font-size: 14px; margin-left:5px; background-color: #d3d3d3; padding: 4px; border-radius: 4px;">ℹ️ If you notice any issues in this documentation, you can [edit this document](https://github.com/harness-community/drone-push-helm-chart-docker-registry/blob/main/README.md) to improve it.</span>
