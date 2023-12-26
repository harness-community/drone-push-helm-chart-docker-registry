import unittest
from unittest.mock import patch
import os
from io import StringIO
from main import main_function
import subprocess


class TestPushOCIChartToRegistry(unittest.TestCase):

    # Function successfully packages helm chart and pushes it to Docker Hub
    def test_successful_chart_push(self, mocker):
        # Set up environment variables
        os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
        os.environ["PLUGIN_DOCKER_USERNAME"] = "testuser"
        os.environ["PLUGIN_DOCKER_PASSWORD"] = "testpassword"
        os.environ["PLUGIN_CHART_PATH"] = "chart"

        # Mock subprocess.run to return success
        mocker.patch("subprocess.run").return_value.returncode = 0

        # Call the main function
        main_function()

        # Assert that subprocess.run was called with the correct arguments
        subprocess.run.assert_called_with(
            ["helm", "push", "mywebapp-1.0.0.tgz", "oci://registry.hub.docker.com/testuser"])

        # Assert that the expected output is printed
        expected_output = 'Chart pushed successfully.'
        assert mock_stdout.getvalue().strip() == expected_output

        # Function successfully handles optional chart path
    def test_optional_chart_path(self, mocker):
        # Set up environment variables
        os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
        os.environ["PLUGIN_DOCKER_USERNAME"] = "testuser"
        os.environ["PLUGIN_DOCKER_PASSWORD"] = "testpassword"

        # Mock subprocess.run to return success
        mocker.patch("subprocess.run").return_value.returncode = 0

        # Call the main function
        main_function()

        # Assert that subprocess.run was called with the correct arguments
        subprocess.run.assert_called_with(
            ["helm", "push", "mywebapp-1.0.0.tgz", "oci://registry.hub.docker.com/testuser"])

        # Assert that the expected output is printed
        expected_output = 'Chart pushed successfully.'
        assert mock_stdout.getvalue().strip() == expected_output

        # Function exits with code 1 if chart packaging fails
    def test_chart_packaging_failure(self, mocker):
        # Set up environment variables
        os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
        os.environ["PLUGIN_DOCKER_USERNAME"] = "testuser"
        os.environ["PLUGIN_DOCKER_PASSWORD"] = "testpassword"

        # Mock subprocess.run to raise an exception
        mocker.patch("subprocess.run").side_effect = subprocess.CalledProcessError(
            1, "helm package")

        # Call the main function and capture the output
        with pytest.raises(SystemExit) as e:
            main_function()

        # Assert that the exit code is 1
        assert e.type == SystemExit
        assert e.value.code == 1


if __name__ == '__main__':
    unittest.main()
