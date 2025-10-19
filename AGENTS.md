# Repository Guidelines

## Project Structure & Module Organization
- `internal/provider/` exposes the Terraform provider entry point and now registers all validatefx functions.
- `internal/functions/` wraps string validators as Terraform functions; reuse these helpers when adding new callable utilities.
- `internal/validators/` holds the core Go validators and their unit tests.
- `integration/` contains HCL scenarios that exercise every exported function end to end.
- `examples/` provides practitioner-facing usage samples grouped by validator type.

## Build, Test, and Development Commands
- `go fmt ./...` — format Go sources across the repository; run before committing.
- `go test ./...` — execute unit tests, ensuring validators and helper packages stay green.
- `make build` — compile the provider binary into `bin/terraform-provider-validatefx`.
- `make install` — install the provider into the local Terraform plugin directory for manual validation.

## Coding Style & Naming Conventions
- Go files must remain `gofmt` clean with tabs for indentation.
- Exported Go symbols follow mixed-case naming (`ValidateSomething`), while Terraform function names use the `validatefx_<name>` pattern.
- HCL examples and integration configs prefer two-space indentation and descriptive local variable names.

## Testing Guidelines
- Unit tests live beside their targets (e.g., `internal/validators/base64_test.go`) and should cover success, failure, and null/unknown cases.
- Integration tests rely on Docker (`make dev`) and the `integration/` directory; update scenarios whenever adding new provider functions.
- Use table-driven tests for new validators and ensure `go test ./...` passes prior to submission.

## Commit & Pull Request Guidelines
- Commit messages typically use the imperative mood (`Add credit-card validator`, `Refactor functions helper`). Keep them scoped and reference issues where applicable.
- Pull requests should describe the validator/function changes, list testing evidence (`go test`, Terraform run), and link to any related integration updates.
- Include screenshots or logs only when troubleshooting Terraform apply output; otherwise prefer concise bullet summaries.

## Security & Configuration Tips
- Avoid hardcoding secrets in examples; rely on Terraform variables when demonstrating configurable values.
- When adding new validators, ensure external dependencies are vetted and present in `go.mod` with minimal version bump.
