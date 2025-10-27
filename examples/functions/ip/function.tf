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
  addresses = {
    loopback_v4 = "127.0.0.1"
    loopback_v6 = "::1"
    bad_address = "999.999.999.999"
  }

  validation_results = {
    for name, addr in local.addresses : name => provider::validatefx::ip(addr)
  }

  assert_loopback_v6 = provider::validatefx::assert(
    provider::validatefx::ip(local.addresses.loopback_v6),
    "loopback must be a valid IP"
  )
}

output "ip_validation" {
  value = local.validation_results
}

output "assert_ip" {
  value = local.assert_loopback_v6
}
