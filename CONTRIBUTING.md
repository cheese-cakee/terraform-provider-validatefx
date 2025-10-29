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


### Pre-commit Hooks
This repository uses [pre-commit](https://pre-commit.com/) to automatically format and lint code before it is committed. This helps maintain code quality and consistency.
#### Installation
To use the hooks, you must install them locally:
1.  Install the `pre-commit` tool. A common way is with Python's package manager:
    ```bash
    pip install pre-commit
    ```
2.  Install the hooks in this repository. From the root directory, run:
    ```bash
    pre-commit install
    ```
After this, the hooks (including `terraform fmt`, `go fmt`, and `make lint`) will run automatically on every `git commit`. If they find an issue, they may fix it for you or stop the commit so you can fix it manually.
