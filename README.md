# Introducing the _drone-push-helm-chart-docker-registry_ Plugin

At Harness, we are always striving to make Continuous Integration (CI) and Continuous Deployment (CD) as smooth and efficient as possible. We understand the importance of managing and deploying Helm charts with ease. That's why we are excited to introduce the drone-push-helm-chart-docker-registry plugin, designed to simplify the packaging and pushing of Helm charts to a Docker registry.

### What is the _drone-push-helm-chart-docker-registry_ Plugin?

The drone-push-helm-chart-docker-registry is a Drone plugin that streamlines the process of packaging and pushing Helm charts to a Docker registry. Helm charts are essential for deploying applications, and this plugin makes it easier than ever to integrate this step into your CI/CD pipeline.

### Build and Docker Image

Building and using the plugin is a straightforward process. You can run the script directly using the following command:

    PLUGIN_CHART_NAME=CHART_NAME \
    PLUGIN_CHART_VERSION=CHART_VERSION \
    PLUGIN_DOCKER_REGISTRY=DOCKER_REGISTRY \
    PLUGIN_CHART_PATH=CHART_PATH \
    PLUGIN_DOCKER_USERNAME=DOCKER_USERNAME \
    PLUGIN_DOCKER_PASSWORD=DOCKER_PASSWORD \
    python3 main.py

Additionally, you can build the Docker image with these commands:

    docker buildx build -t DOCKER_ORG/drone-helm-chart-docker-registry --platform linux/amd64 .

This will build the image for the Linux AMD64 platform.

### Usage in Harness CI

Integrating the drone-push-helm-chart-docker-registry plugin into your Harness CI pipeline is easy. You can use Docker to run the plugin with environment variables. Here's how:

    docker run --rm \
    -e PLUGIN_CHART_NAME=${CHART_NAME} \
    -e PLUGIN_CHART_VERSION=${CHART_VERSION} \
    -e PLUGIN_DOCKER_REGISTRY=${DOCKER_REGISTRY} \
    -e PLUGIN_CHART_PATH=${CHART_PATH} \
    -e PLUGIN_DOCKER_USERNAME=${DOCKER_USERNAME} \
    -e PLUGIN_DOCKER_PASSWORD=${DOCKER_PASSWORD} \
    -v $(pwd):$(pwd) \
    -w $(pwd) \
    harnesscommunity/drone-helm-chart-docker-registry

In your Harness CI pipeline, you can define the plugin as a step, like this:

    - step:
    type:  Plugin
    name:  helm  to  docker
    identifier:  helm_to_docker
    spec:
    connectorRef:  docker-registry-connector
    image:  harnesscommunity/drone-helm-chart-docker-registry
    settings:
    chart_name:  mywebapp
    docker_username:  <+variable.docker_username>
    docker_password:  <+secrets.getValue("pat-token")>
    chart_path:  path-to-helm-chart
    chart_version:  1.0.0
    docker_registry:  registry.hub.docker.com

#### Plugin Options

The plugin offers several options for customization:

- **_-\-chart_name_**: The name of your Helm chart. You should replace ${CHART_NAME} with the actual name of your Helm chart.

- **_-\-chart_version_**: This variable allows you to specify the version of your Helm chart (default: 1.0.0). You can replace ${CHART_VERSION} with the desired version.

- **_-\-docker_registry_**: Use this variable to specify the Docker registry where you want to push your Helm chart (default: registry.hub.docker.com). Replace ${DOCKER_REGISTRY} with your preferred registry.

- **_-\-chart_path_**: This variable is used to define the path to your Helm chart's directory (default: root directory). Replace ${CHART_PATH} with the actual path to your Helm chart.

- **_-\-docker_username_**: Here, you provide your Docker login username. Replace ${DOCKER_USERNAME} with your Docker username.

- **_-\-docker_password_**: This variable requires your Docker login password. Replace ${DOCKER_PASSWORD} with your Docker password.

These environment variables are essential for configuring and customizing the behavior of the drone-push-helm-chart-docker-registry plugin when it's executed as a Docker container. They allow you to provide specific values and credentials required for the plugin to package and push your Helm chart to the Docker registry of your choice. Make sure to set these variables according to your project's needs.

## Get Started with the drone-push-helm-chart-docker-registry Plugin

Whether you are a seasoned DevOps professional or just getting started with CI/CD, the drone-push-helm-chart-docker-registry plugin can streamline your deployment process and help you manage Helm charts effortlessly. Give it a try and see how it simplifies your CI/CD pipelines!

For more information, documentation, and updates, please visit our GitHub repository: [drone-push-helm-chart-docker-registry](https://github.com/harnesscommunity/drone-push-helm-chart-docker-registry).
