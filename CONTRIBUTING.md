# Contributing Guide

Thanks for your interest in contributing to **terraform-provider-validatefx**! This document outlines the recommended local development workflow and expectations for pull requests.

## Prerequisites

- Go 1.25.2 or newer (matching the version configured in `go.mod`)
- Docker and Docker Compose (required for integration tests)
- [`golangci-lint`](https://golangci-lint.run/) v1.61 or newer available on your `PATH`
- [`tfplugindocs`](https://github.com/hashicorp/terraform-plugin-docs) v0.19.x for documentation generation (`go install github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs@v0.19.2`)

## Getting Started

```bash
git clone https://github.com/The-DevOps-Daily/terraform-provider-validatefx.git
cd terraform-provider-validatefx
make deps
```

## Common Tasks

| Command | Description |
| --- | --- |
| `make fmt` | Format Go source files using `go fmt`. |
| `make build` | Compile the provider to verify it builds. |
| `make test` | Run unit tests. |
| `make lint` | Execute `golangci-lint` using the local installation. |
| `make docs` | Regenerate the Markdown docs under `docs/` via `tfplugindocs`. |
| `make integration` | Build the Docker image and run the Terraform integration scenario end-to-end. |
| `make clean` | Remove build artifacts and reset local integration state. |

> **Tip:** The `make help` command lists all available targets and descriptions.

## Pull Request Checklist

Before opening a PR, please ensure:

- Code is formatted (`make fmt`).
- Unit tests pass (`make test`).
- Linting reports no issues (`make lint`).
- Integration tests succeed (`make integration`) when applicable.
- Documentation is regenerated (`make docs`) when schema changes affect the published docs.
- Commits are descriptive and scoped to a logical change.
- PR includes a brief summary explaining motivation and testing performed.

## Reporting Issues

If you encounter bugs or have feature ideas, please open an issue with details about the problem, reproduction steps, and environment information (Go version, OS, etc.).

We appreciate your contributions—thank you for helping improve the provider!
