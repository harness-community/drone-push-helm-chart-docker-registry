# drone-push-helm-chart-docker-registry

Drone plugin to package and push a helm chart to a docker registry.

## Build

Run the script directly using the command:

```python
python3 main.py
```

## Docker

Build the Docker image with the following commands:

```
docker buildx build -t DOCKER_ORG/drone-helm-docker-registry --platform linux/amd64 .
```

Please build the image for Linux AMD64 platform

## Usage

```bash
docker run --rm \
  -e PLUGIN_CHART_NAME=${CHART_NAME} \
  -e PLUGIN_CHART_VERSION=${CHART_VERSION} \
  -e PLUGIN_DOCKER_REGISTRY=${DOCKER_REGISTRY} \
  -e PLUGIN_CHART_PATH=${CHART_PATH} \
  -e PLUGIN_DOCKER_USERNAME=${DOCKER_USERNAME} \
  -e PLUGIN_DOCKER_PASSWORD=${DOCKER_PAT} \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
harnesscommunity/drone-helm-docker-registry
```

In Harness CI,

```yaml
- step:
    type: Plugin
    name: helm to docker
    identifier: helm_to_docker
    spec:
      connectorRef: akshitnodeserverconnector
      image: harnesscommunity/drone-helm-docker-registry
      settings:
        chart_name: mywebapp
        docker_username: <+variable.docker_username>
        docker_password: <+secrets.getValue("helmpluginpat")>
        chart_path: test
        chart_version: 1.0.0
        docker_registry: registry.hub.docker.com
```

```
GLOBAL OPTIONS:
   --chart_name         value   required     Helm Chart Name [$PLUGIN_CHART_NAME]
   --chart_version      value   optional     Helm Chart Version (Default: 1.0.0) [$PLUGIN_CHART_VERSION]
   --docker_registry    value   optional     Docker Registry for pushing Helm Chart (Default: registry.hub.docker.com) [$PLUGIN_DOCKER_REGISTRY]
   --chart_path         value   optional     Path to Helm Chart's Directory (Default: Root Directory) [$PLUGIN_CHART_PATH]
   --docker_username    value   required     Docker Login Username [$PLUGIN_DOCKER_USERNAME]
   --docker_password    value   required     Docker Login Password [$PLUGIN_DOCKER_PASSWORD]
```

Please make sure the Chart name and version match the Chart.yaml, otherwise chart packaging will fail.
