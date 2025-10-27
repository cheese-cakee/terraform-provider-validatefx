---
title: "any_valid"
description: "Return true when at least one validation result passes"
---

# any_valid Function

Evaluate a list of boolean validation results and report success as soon as one element is `true`.

## Definition

```
provider::validatefx::any_valid(checks)
```

## Arguments

- `checks` (`list(bool)`) – Collection of boolean validation results. `null` entries are treated as failing; `unknown` values propagate unless a `true` is found.

## Return Value

- `bool`

Returns `true` when any element is `true`, `false` when all are definitively `false`/`null`, and `unknown` when no `true` value is present but at least one element is `unknown`.

## Examples

```hcl
locals {
  validation_checks = [
    provider::validatefx::email("invalid"),
    provider::validatefx::uuid("d9428888-122b-11e1-b85c-61cd3cbb3210"),
  ]

  any_check_passes = provider::validatefx::any_valid(local.validation_checks)
}
```
