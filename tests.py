import unittest
from unittest.mock import patch
import os
from io import StringIO
from main import main_function


class TestPushOCIChartToRegistry(unittest.TestCase):

    @patch("subprocess.run")
    def test_successful_chart_push(self, mock_subprocess_run):
        os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
        os.environ["PLUGIN_DOCKER_USERNAME"] = "testuser"
        os.environ["PLUGIN_DOCKER_PASSWORD"] = "testpassword"
        os.environ["PLUGIN_CHART_PATH"] = "chart"

        mock_subprocess_run.return_value.returncode = 0

        with patch("sys.stdout", new_callable=StringIO) as mock_stdout:
            main_function()

        mock_subprocess_run.assert_called_with(
            ["helm", "push", "mywebapp-1.0.0.tgz", "oci://registry.hub.docker.com/testuser"])

        expected_output = 'Chart pushed successfully.'
        self.assertEqual(mock_stdout.getvalue().strip(), expected_output)

    # @patch("subprocess.run")
    # def test_failed_chart_push(self, mock_subprocess_run):
    #     mock_subprocess_run.return_value.returncode = 1

    #     os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
    #     os.environ["PLUGIN_DOCKER_USERNAME"] = "testuser"
    #     os.environ["PLUGIN_DOCKER_PASSWORD"] = "testpassword"

    #     with patch("sys.stdout", new_callable=StringIO) as mock_stdout:
    #         main_function()

    #     mock_subprocess_run.assert_called_with(
    #         ["helm", "push", "mywebapp-1.0.0.tgz", "oci://registry.hub.docker.com/testuser"])

    #     expected_output = "Failed to push chart!"  # Update the expected output
    #     self.assertEqual(mock_stdout.getvalue().strip(), expected_output)


if __name__ == '__main__':
    unittest.main()
