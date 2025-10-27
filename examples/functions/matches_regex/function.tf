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
  usernames = {
    valid   = "user_123"
    invalid = "Invalid-User"
  }

  validation_results = {
    for name, value in local.usernames : name => provider::validatefx::matches_regex(value, "^[a-z0-9_]+$")
  }

  enforce_username = provider::validatefx::assert(
    provider::validatefx::matches_regex(local.usernames.valid, "^[a-z0-9_]+$"),
    "username must contain only lowercase letters, numbers, or underscores"
  )
}

output "matches_regex_results" {
  value = local.validation_results
}

output "assert_matches_regex" {
  value = local.enforce_username
}
