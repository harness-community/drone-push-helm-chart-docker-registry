import unittest
from unittest.mock import patch
import os
from io import StringIO
from main import main_function


class TestPushOCIChartToRegistry(unittest.TestCase):

    def tearDown(self):
        os.environ["PLUGIN_CHART_NAME"] = ""
        os.environ["PLUGIN_CHART_VERSION"] = ""
        os.environ["PLUGIN_DOCKER_REGISTRY"] = ""
        os.environ["PLUGIN_DOCKER_USERNAME"] = ""
        os.environ["PLUGIN_DOCKER_PASSWORD"] = ""
        os.environ["PLUGIN_CHART_PATH"] = ""

    @patch("subprocess.run")
    def test_successful_chart_push(self, mock_subprocess_run):
        os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
        test_docker_username = os.environ["TEST_DOCKER_USERNAME"]
        test_docker_password = os.environ["TEST_DOCKER_PASSWORD"]
        os.environ["PLUGIN_DOCKER_USERNAME"] = test_docker_username
        os.environ["PLUGIN_DOCKER_PASSWORD"] = test_docker_password
        os.environ["PLUGIN_CHART_VERSION"] = "5.0.0"
        os.environ["PLUGIN_DOCKER_REGISTRY"] = "registry.hub.docker.com"

        mock_subprocess_run.return_value.returncode = 0

        with patch("sys.stdout", new_callable=StringIO) as mock_stdout:
            main_function()

        mock_subprocess_run.assert_called_with(
            ["helm", "push", "mywebapp-5.0.0.tgz", f"oci://registry.hub.docker.com/{test_docker_username}"])

        expected_output = 'Chart pushed successfully.'
        self.assertEqual(mock_stdout.getvalue().strip(), expected_output)

    @patch("subprocess.run")
    def test_failed_to_package_chart(self, mock_subprocess_run):
        os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
        test_docker_username = os.environ["TEST_DOCKER_USERNAME"]
        test_docker_password = os.environ["TEST_DOCKER_PASSWORD"]
        os.environ["PLUGIN_DOCKER_USERNAME"] = test_docker_username
        os.environ["PLUGIN_DOCKER_PASSWORD"] = test_docker_password
        os.environ["PLUGIN_CHART_PATH"] = "chart"

        mock_subprocess_run.return_value.returncode = 1

        with patch("sys.stdout", new_callable=StringIO) as mock_stdout:
            with self.assertRaises(SystemExit) as context:
                main_function()

        self.assertEqual(context.exception.code, 1)

        expected_output = 'Failed to package chart!'
        self.assertEqual(mock_stdout.getvalue().strip(), expected_output)

    @patch("subprocess.run")
    def test_chart_name_not_provided(self, mock_subprocess_run):
        test_docker_username = os.environ["TEST_DOCKER_USERNAME"]
        test_docker_password = os.environ["TEST_DOCKER_PASSWORD"]
        os.environ["PLUGIN_DOCKER_USERNAME"] = test_docker_username
        os.environ["PLUGIN_DOCKER_PASSWORD"] = test_docker_password
        os.environ["PLUGIN_CHART_PATH"] = "chart"

        mock_subprocess_run.return_value.returncode = 0

        with patch("sys.stdout", new_callable=StringIO) as mock_stdout:
            with self.assertRaises(SystemExit) as context:
                main_function()

        self.assertEqual(context.exception.code, 1)

        expected_output = 'Please provide a chart name'
        self.assertEqual(mock_stdout.getvalue().strip(), expected_output)

    @patch("subprocess.run")
    def test_username_and_password_not_provided(self, mock_subprocess_run):
        os.environ["PLUGIN_CHART_NAME"] = "mywebapp"
        os.environ["PLUGIN_CHART_PATH"] = "chart"

        mock_subprocess_run.return_value.returncode = 0

        with patch("sys.stdout", new_callable=StringIO) as mock_stdout:
            with self.assertRaises(SystemExit) as context:
                main_function()

        self.assertEqual(context.exception.code, 1)

        expected_output = 'Please provide a username and a password'
        self.assertEqual(mock_stdout.getvalue().strip(), expected_output)


if __name__ == '__main__':
    unittest.main()
