terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = ">= 0.0.0"
    }
  }
}

provider "validatefx" {}

variable "emails" {
  type = list(string)
  default = [
    "alice@example.com",
    "bad-email",
  ]
}

locals {
  checked_emails = [
    for addr in var.emails : {
      address = addr
      valid   = try(validatefx_email(addr), false)
    }
  ]
}

output "valid_emails" {
  value = [for item in local.checked_emails : item if item.valid]
}

output "invalid_emails" {
  value = [for item in local.checked_emails : item if !item.valid]
}
