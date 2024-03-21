# Treblle CLI

API Insights scores your API using over 30 tests taken from standards and industry best practices across three categories:

- **Design**: how well your API is structured. Is it a well-crafted ship, ready to withstand the rough seas of user demands and scalability?

- **Performance**: does your API respond quickly? It's like checking the wind in your sails – are you moving swiftly and smoothly?

- **Security**: how safe and secure is your API? It's akin to having a strong hull to protect against the stormy seas of cyber threats.


<div align="center">
  <img src="https://assets.apiinsights.io/insights-CLI.png"/>
</div>

## Overview

Get insights into your OpenAPI specifications and API using the Treblle CLI.

```md
The Treblle CLI tool.

Usage:
  treblle-cli [command]

Available Commands:
  help        Help about any command
  insights    Generate an API Insights report.

Flags:
  -h, --help      help for treblle
  -v, --version   version for treblle

Use "treblle-cli [command] --help" for more information about a command.
```

## Installing

Using homebrew on Mac or Linux, you can install the Treblle CLI using:

```bash
brew tap treblle/treblle
brew install treblle
```

## API Insights

You can use the `insights` command on the Treblle CLI to upload your OpenAPI Specification to API Insights, and get your report back in just a few seconds.

```bash
treblle-cli insights path/to/openapi-spec.json
```

You then have various options available to you. By default you will get an overview of the report, which will allow you to have a high-level overview of your API Insights score. You can also use any of the following options:

### Technology Profile

Add the `--technology` flag to see the technology profile discovered on your API.

### Details

The `--details` flag requires one of the following options:

- `--details=performance` to show the performance tests pass/fail for your API.
- `--details=security` to show the security test pass/fail for your API.
- `--details=quality` to show the quality test pass/fail for your API.
- `--details=all` to show the performance, security, and quality test pass/fail for your API.

### Miniumum score limit

The `--m` flag gives you a possibility to set a minimum score in CI/CD enviroments:

- `--m 90` to limit any of test to be at least 90 to pass



