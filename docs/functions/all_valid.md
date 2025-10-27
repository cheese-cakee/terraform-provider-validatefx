---
title: "all_valid"
description: "Require every validation result to be true"
---

# all_valid Function

Combine multiple validation results and return `true` only when every entry in the list evaluates to `true`.

## Definition

```
provider::validatefx::all_valid(checks)
```

## Arguments

- `checks` (`list(bool)`) – Collection of boolean validation results. `null` or `unknown` elements are treated as failing results and may propagate unknowns.

## Return Value

- `bool`

Returns `true` when all checks are `true`, `false` when any element is `false` (or `null`), and `unknown` if at least one element is `unknown` while the remainder are not deterministically failing.

## Examples

```hcl
locals {
  validation_checks = [
    provider::validatefx::email("user@example.com"),
    provider::validatefx::uuid("d9428888-122b-11e1-b85c-61cd3cbb3210"),
  ]

  all_checks_pass = provider::validatefx::all_valid(local.validation_checks)
}
```
