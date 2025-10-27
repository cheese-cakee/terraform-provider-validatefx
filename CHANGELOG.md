# Changelog

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
