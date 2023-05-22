# URL Monitoring Tool

This is a simple URL monitoring tool written in GoLang that periodically checks the availability and response time of a list of URLs. It sends notifications when a URL is down or experiencing issues. The tool is designed to ensure high availability and scalability.

## Features

- Periodically checks the availability and response time of URLs
- Sends notifications when a URL is down or experiencing issues
- Configurable interval for checking URLs
- Supports multiple notification methods

## Installation

1. Clone the repository: `git clone https://github.com/your-username/url-monitoring-tool.git`
2. Change to the project directory: `cd url-monitoring-tool`
3. Build the application: `go build`

## Usage

1. Edit the `urls/urls.txt` file to specify the URLs you want to monitor. Each URL should be on a separate line.
2. Configure the monitoring settings in the `config/config.go` file, such as the monitoring interval and notification options.
3. Run the application: `./url-monitoring-tool`

## Configuration

The tool can be configured by modifying the `config/config.go` file. The following settings are available:

- `MonitoringInterval`: The interval at which the URLs will be checked.
- `NotificationEmail`: The email address to which email notifications will be sent.
- Add more configuration options as needed for your specific notification methods.

## License

This project is licensed under the GNU General Public License v3.0. See the LICENSE file for details.
