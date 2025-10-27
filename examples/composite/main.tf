terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = ">= 0.1.0"
    }
  }
}

provider "validatefx" {}

locals {
  checks = [
    provider::validatefx::email("user@example.com"),
    provider::validatefx::uuid("d9428888-122b-11e1-b85c-61cd3cbb3210"),
  ]

  all_pass = provider::validatefx::all_valid(local.checks)
  any_pass = provider::validatefx::any_valid(local.checks)
}

output "composite_results" {
  value = {
    all = local.all_pass
    any = local.any_pass
  }
}
