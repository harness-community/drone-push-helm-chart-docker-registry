# Plugin Name: Push OCI Chart to Registry
# Description: Pushes an Helm Chart to a Docker Registry

import os
import subprocess

# Environment Variables


def main_function():

    CHART_NAME = os.getenv("PLUGIN_CHART_NAME")
    CHART_VERSION = os.getenv("PLUGIN_CHART_VERSION", "1.0.0")

    DOCKER_REGISTRY = os.getenv(
        "PLUGIN_DOCKER_REGISTRY", 'registry.hub.docker.com')

    DOCKER_USERNAME = os.getenv(
        "PLUGIN_DOCKER_USERNAME")
    DOCKER_PASSWORD = os.getenv(
        "PLUGIN_DOCKER_PASSWORD")

    CHART_PATH = os.getenv("PLUGIN_CHART_PATH")

    DOCKER_NAMESPACE = os.getenv("PLUGIN_DOCKER_NAMESPACE")

    if (CHART_NAME is None or CHART_NAME == ""):
        print("Please provide a chart name")
        exit(1)

    if (DOCKER_USERNAME is None or DOCKER_PASSWORD is None or DOCKER_USERNAME == "" or DOCKER_PASSWORD == ""):
        print("Please provide a username and a password")
        exit(1)

    if (DOCKER_NAMESPACE is None or DOCKER_NAMESPACE == ""):
        print("Please provide a namespace")
        exit(1)

    if (CHART_PATH is not None):
        os.chdir(CHART_PATH)

    try:
        if (subprocess.run(["helm", "package", "--dependency-update", "."]).returncode != 0):
            raise Exception("Failed to package chart!")
    except:
        print("Failed to package chart!")
        exit(1)

    chart_filename = f"{CHART_NAME}-{CHART_VERSION}.tgz"

    try:
        login_command = ['helm', 'registry', 'login', DOCKER_REGISTRY,
                         '-u', DOCKER_USERNAME, '-p', DOCKER_PASSWORD]
        if (subprocess.run(login_command).returncode != 0):
            raise Exception("Failed to login!")
    except:
        print("Failed to login!")
        exit(1)

    try:
        docker_push_command = ["helm", "push", chart_filename,
                               f"oci://{DOCKER_REGISTRY}/{DOCKER_NAMESPACE}"]
        if (subprocess.run(docker_push_command).returncode != 0):
            raise Exception("Failed to push chart!")
        else:
            print("Chart pushed successfully.")
    except:
        print("Failed to push chart!")
        exit(1)


if __name__ == "__main__":
    main_function()
