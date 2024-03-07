# Treblle CLI

```md
The Treblle CLI tool.

Usage:
  treblle [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  insights    Generate an API Insights report.

Flags:
  -h, --help      help for treblle
  -v, --version   version for treblle

Use "treblle [command] --help" for more information about a command.
```

## Installing

Using homebrew on Mac or Linux, you can install the Treblle CLI using:

```bash
brew install treblle
```

## API Insights

You can use the `insights` command on the Treblle CLI to upload your OpenAPI Specification to API Insights, and get your report back in just a few seconds.

```bash
treblle insights path/to/openapi-spec.json
```
