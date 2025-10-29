## [2025-10-29]

- Merge pull request #135 from The-DevOps-Daily/fix/string-length-integration (75379a4)
  * Fix failing tests (be9d1db)
  * Refactor string length function logic (abecf39)
  * Integrate string length validator (6b504be)

- Merge pull request #132 from Madhur-Dodake/feat/string-length-validator-issue-105 (9b98b71) #105
  * chore: run gofmt to fix lint formatting issue (c77e1e6)
  * fix: rename min/max and remove unused ctx to satisfy linter (7bc3f30)
  * feat(validators): add string length validator (#105) (986d9c4)

- Merge pull request #133 from The-DevOps-Daily/feature/cidr-validator (f1fc4b7) #100
  * Fix integration CIDR formatting (0809e49)
  * Refine CIDR docs and integration (53cfc96)
  * Add CIDR validation function (#100) (9a3037c)

- Merge pull request #128 from The-DevOps-Daily/dependabot/github_actions/actions/setup-go-6 (cb51c3e)
  * Bump actions/setup-go from 5 to 6 (0e44d89)

- Merge pull request #130 from Rachitb0611/docs/add-contributors-badge (9e90793)
  * docs: Add contributors badge to README (e5b1c5a)

- Merge pull request #129 from The-DevOps-Daily/dependabot/github_actions/goreleaser/goreleaser-action-6 (a69aba5)
  * Bump goreleaser/goreleaser-action from 5 to 6 (526fb34)

- Merge pull request #127 from The-DevOps-Daily/dependabot/github_actions/actions/upload-artifact-5 (e74f0ee)
  * Bump actions/upload-artifact from 4 to 5 (b9a95d6)

- Merge pull request #126 from The-DevOps-Daily/dependabot/github_actions/actions/checkout-5 (6babd66)
  * Bump actions/checkout from 4 to 5 (1e87f4d)

- Merge pull request #124 from cheese-cakee/main (501493a)
  * Create dependabot.yml (a094c56)
  * Delete .github/workflows/dependabot.yml (bf3180c)
  * Create dependabot.yml (d30096c)

- Merge pull request #125 from The-DevOps-Daily/auto-issue-creation (7a2eb13)
  * Add auto-issue creation script for future reference (b47e6a0)

- Merge pull request #1 from The-DevOps-Daily/main (7f165d6)
  * docs: add OS-specific installation and troubleshooting guide for ValidateFX provider (e3c639e)

- Merge pull request #123 from Madhur-Dodake/docs/os-installation-111 (3165060)
  * docs: add OS-specific installation and troubleshooting guide for ValidateFX provider (e3c639e)

- Merge pull request #120 from cheese-cakee/main (4e37890)
  * Update CONTRIBUTING.md (1f05187)
  * Create .pre-commit-config.yaml (e9bb56f)




# Changelog

## [0.1.3] - 2025-10-28

### Features

- Add an HTTP/HTTPS URL validator exposed as `provider::validatefx::url`, including schema tests and Terraform coverage (`faf98d4`, `6a545cf`, `51bef43`).
- Expose provider metadata through the new `provider::validatefx::version` function with integration coverage and documentation updates (`9cdba92`, `84ba24d`, `18dd815`, `81e29af`).

### Improvements

- Expand Terraform integration scenarios to exercise additional validators and the provider version endpoint (`211d656`, `bec4e33`, `c6a6c4f`).
- Add defensive tests ensuring string validation functions surface diagnostics for non-string inputs (`c386eb0`, `e61d50b`).
- Restructure examples and documentation to streamline generation and add a provider quick-start snippet (`7027ef8`, `86db796`, `f67b9b2`, `9472110`).

### Bug Fixes

- Harden URL validation behavior and align imports and formatting (`860cb71`, `6a545cf`).
- Stabilize integration expectations by correcting email/base64 fixtures and handling null inputs (`6992130`, `5a01c2c`).
- Resolve intermittent test failures surfaced during integration expansion (`bfdba96`, `5676adc`).

---

## [0.1.2] - 2025-10-27

### Features

- Add composite validation helpers `all_valid` and `any_valid` for aggregating multiple checks (`a3e1c9a`, `8574455`).
- Expose the existing phone E.164 validator as a Terraform function with docs and examples (`5f62599`).
- Introduce the `matches_regex` Terraform function for pattern validation (`f825340`).

### Bug Fixes

- Cache compiled regular expressions in the `matches_regex` validator to avoid repeated compilation (`db161f7`).

### Misc

- Preserve the provider docs index template during documentation generation (`4171e03`).
- Publish a custom provider index document to improve docs navigation (`337b172`).

---

## [0.1.1] - 2025-10-26

### Features

- Add Terraform functions for JSON structure validation, Semantic Versioning checks, and IP address validation (`1ed7d28`, `ee2e5f3`, `19140c2`).
- Automate regeneration of the README “Available Functions” table to keep documentation in sync (`3bf9caa`, `3c8133a`).

### Bug Fixes

- Correct integration test Docker plugin path, README build badge, and Terraform Registry URLs (`13b6573`, `162c267`, `e3d40a6`).

### Misc

- Remove unused internal function helpers discovered during review (`d397f4d`).

---

## [0.1.0] - 2025-09-28

### Features

- Initial release of the provider scaffold with validators for email, UUID, base64, credit card, domain, and phone numbers plus Terraform examples and unit tests (`046cb51`, `c07ff64`, `0a478c1`, `35497a3`, `211bedc`, `8ce87fd`).
- Add Terraform integration workflows and supporting Makefile targets to validate the provider end to end (`0e74156`, `6944b72`, `2d42556`, `c6845dd`).
- Introduce release automation via GitHub Actions (`58c069f`).

### Bug Fixes

- Iterate on release workflows to resolve checksum, packaging, and pipeline failures (`6d823c7`, `5980981`, `5bb84d9`, `03babef`, `1679763`).
- Fix function parameter naming issues uncovered during early CI automation (`196831d`).

### Misc

- Add contributor guidelines, AGENTS metadata, and README badges to polish the project presentation (`c6845dd`, `7892797`, `9ff1444`).
- Expand validator test coverage with comprehensive table-driven suites (`8b222a5`, `24e67c5`).
