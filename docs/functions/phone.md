---
title: "phone"
description: "Validate E.164 compliant phone numbers"
---

# Phone Function

Validates a string using the same semantics as the `validatefx::phone` Terraform function.

## Definition

```
provider::validatefx::phone(value)
```

## Arguments

- `value` (`string`) – Phone number to validate. Accepts null or unknown.

## Return Value

- `bool`

Returns `true` when the value is a valid E.164 phone number, `false` otherwise.

## Examples

```hcl
locals {
  phone_is_valid = provider::validatefx::phone("+14155552671")
}
```
