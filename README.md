# 🧩 Terraform Provider - ValidateFX

Reusable validation functions for Terraform, built with the latest [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework).

ValidateFX lets you write cleaner, more expressive validations using functions like `is_email`, `is_uuid`, and `is_semver`.

---

## 🚀 Example

```hcl
terraform {
  required_providers {
    validatefx = {
      source  = "thedevopsdaily/validatefx"
      version = "0.1.0"
    }
  }
}

provider "validatefx" {}

variable "email" {
  type = string
  validation {
    condition     = provider::validatefx::is_email(var.email)
    error_message = "Must be a valid email address"
  }
}
````

---

## ⚙️ Development

```bash
git clone https://github.com/The-DevOps-Daily/terraform-provider-validatefx.git
cd terraform-provider-validatefx
go mod tidy
make build
make install
make dev
```

Example usage in `examples/basic/main.tf`.

---

## 🧩 Available Functions

| Function            | Description                |
| ------------------- | -------------------------- |
| `is_email(string)`  | Validates email format     |
| `is_uuid(string)`   | Validates UUID             |
| `is_semver(string)` | Validates semantic version |

---

## 💡 Contributing

Open to PRs!
Good first issues include adding new validators like `is_ip`, `is_hostname`, or `matches_regex`.

---

## 📜 License

MIT © 2025 [DevOps Daily](https://github.com/The-DevOps-Daily)
