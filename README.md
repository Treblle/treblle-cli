# Treblle CLI

Get insights into your OpenAPI specifications and API using the Treblle CLI.

<div align="center">
  <img src="https://assets.treblle.com/github/insights-CLI.png"/>
</div>

## Overview

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

### `--technology` Technology Profile

Add the `--technology` flag to see the technology profile discovered on your API.

### `--details` Details

The `--details` flag requires one of the following options:

- `--details=performance` to show the performance tests pass/fail for your API.
- `--details=security` to show the security test pass/fail for your API.
- `--details=quality` to show the quality test pass/fail for your API.
- `--details=all` to show the performance, security, and quality test pass/fail for your API.

