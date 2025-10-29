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
  networks = [
    "10.0.0.0/24",
    "bad-cidr",
  ]

  checked_cidrs = [
    for cidr in local.networks : {
      value = cidr
      valid = provider::validatefx::cidr(cidr)
    }
  ]
}

output "checked_cidrs" {
  value = local.checked_cidrs
}
