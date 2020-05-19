# flexpool-cli
[![Build Status](https://travis-ci.org/flexpool/flexpool-cli.svg?branch=master)](https://travis-ci.org/flexpool/flexpool-cli)

CLI tool that wraps **flexpool** API.
> WARNING: This tool is in development and does not cover the the flexpool API fully.

![flexpool-cli screenshot](https://i.imgur.com/vXLZ6ca.png)

# Requirements
* go (v1.13 or higher)

# Install
Install the tool:
```
git clone https://github.com/flexpool/flexpool-cli.git
cd flexpool-cli
make install
```
After that, you would be able to access `flexpool-cli`.

# How to get started?
Add your address(es) to the watchlist by using `flexpool-cli addr add`
```
flexpool-cli add 0x85A20253f8ff8374f65D4F4797F065f8aCe6f136
```
Now you can use `flexpool-cli stat`, `flexpool-cli sum` and other commands

# Usage
```
Usage:
  flexpool-cli [command]

Available Commands:
  addr        Address management
  config      Configuration management
  help        Help about any command
  stat        View the address' stats
  sum         View the summary on watched addresses
  version     Print version

Flags:
  -h, --help   help for flexpool-cli

Use "flexpool-cli [command] --help" for more information about a command.
```

# Uninstall
Just do `make uninstall`

# Troubleshooting
## - Command not found
```
bash: command not found: flexpool-cli
```
Ensure that you have GOPATH & PATH variables set correctly.

# FAQ
## - Where configuration is located?
Linux: `$HOME/.config/flexpool/flexpool-cli/`

macOS: `$HOME/Library/Application Support/flexpool/flexpool-cli`

Windows: `%APPDATA%/flexpool/flexpool-cli`

Also you can use `flexpool-cli config path` to view the active config path. To clean the configuration, use `flexpool-cli config clean`.

# License
GNU General Public License v3
