terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = ">= 0.0.1"
    }
  }
}

provider "validatefx" {}

locals {
  usernames = [
    "devops",
    "too-long-username",
  ]

  checked_lengths = [
    for name in local.usernames : {
      value = name
      valid = provider::validatefx::string_length(name, 3, 10)
    }
  ]
}

output "checked_lengths" {
  value = local.checked_lengths
}
