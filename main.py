# z Plugin Name: Push OCI Chart to Registry
# Description: Pushes an OCI Chart to a Docker Registry

import os
import subprocess

# Environment Variables

# Helm Chart
CHART_NAME = os.getenv("PLUGIN_CHART_NAME")  # ! required
CHART_VERSION = os.getenv("PLUGIN_CHART_VERSION", "1.0.0")

# Docker Repository Details
DOCKER_REGISTRY = os.getenv(
    "PLUGIN_DOCKER_REGISTRY", 'registry.hub.docker.com')  # ? optional

# Docker Hub Credentials
DOCKER_USERNAME = os.getenv(
    "PLUGIN_DOCKER_USERNAME")  # ! required
DOCKER_PASSWORD = os.getenv(
    "PLUGIN_DOCKER_PASSWORD")  # ! required

# Path to Chart
CHART_PATH = os.getenv("PLUGIN_CHART_PATH")  # ? optional


# ? validate environment variables

if (CHART_NAME is None):
    print("Please provide a chart name")
    exit(1)

if (DOCKER_USERNAME is None or DOCKER_PASSWORD is None):
    print("Please provide a username and a password")
    exit(1)


# cd into the chart directory
if (CHART_PATH is not None):
    os.chdir(CHART_PATH)

# Package the helm chart
subprocess.run(["helm", "package", "--dependency-update", "."])

# Construct the chart name
chart_filename = f"{CHART_NAME}-{CHART_VERSION}.tgz"

# Login to Docker Registry
try:
    login_command = ['helm', 'registry', 'login', DOCKER_REGISTRY,
                     '-u', DOCKER_USERNAME, '-p', DOCKER_PASSWORD]
    subprocess.run(login_command)
except:
    print("Failed to login!")
    exit(1)

# Push the chart to Docker Hub
try:
    docker_push_command = ["helm", "push", chart_filename,
                           f"oci://{DOCKER_REGISTRY}/{DOCKER_USERNAME}"]
    subprocess.run(docker_push_command)
except:
    print("Failed to push chart!")
    exit(1)
