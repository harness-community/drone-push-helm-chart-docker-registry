# environment variables

if [ -z "$PLUGIN_CHART_NAME" ]; then
    echo "chart name is required"
    exit 1
fi

if [ -z "$PLUGIN_CHART_VERSION" ]; then
    echo "chart version is required"
    exit 1
fi

if [ -z "$PLUGIN_DOCKER_USERNAME" ]; then
    echo "docker username is required"
    exit 1
fi

if [ -z "$PLUGIN_DOCKER_PASSWORD" ]; then
    echo "docker password is required"
    exit 1
fi

# Set docker registry to docker hub if not set
if [ -z "$PLUGIN_DOCKER_REGISTRY" ]; then
    PLUGIN_DOCKER_REGISTRY="registry.hub.docker.com"
fi

if [ -z "$PLUGIN_CHART_PATH" ]; then
    PLUGIN_PATH="."
fi

# plugin script

echo $PLUGIN_CHART_PATH
cd "$PLUGIN_CHART_PATH"
helm package --dependency-update . || exit 1

CHART_FILENAME="$PLUGIN_CHART_NAME-$PLUGIN_CHART_VERSION.tgz" || exit 1

helm registry login $PLUGIN_DOCKER_REGISTRY -u $PLUGIN_DOCKER_USERNAME -p $PLUGIN_DOCKER_PASSWORD || exit 1

helm push $CHART_FILENAME oci://$PLUGIN_DOCKER_REGISTRY/$PLUGIN_DOCKER_USERNAME || exit 1
echo "Chart pushed to $PLUGIN_DOCKER_REGISTRY/$PLUGIN_DOCKER_USERNAME"