terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = "0.0.1"
    }
  }
}

provider "validatefx" {}

locals {
  emails = [
    "alice@example.com",
    "bad-email",
  ]
}

output "valid_email_count" {
  value = length([for email in local.emails : email if try(validatefx_email(email), false)])
}

output "invalid_email_count" {
  value = length([for email in local.emails : email if !try(validatefx_email(email), false)])
}
