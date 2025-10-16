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
  domains = [
    "example.com",
    "www.example.com",
    "api.v1.example.com",
    "my-domain.example-site.com",
    "invalid..domain",
    "example-.com",
    "-example.com",
    "example.123",
    "example@.com",
  ]

  checked_domains = [
    for domain in local.domains : {
      domain = domain
      valid  = validatefx_domain(domain)
    }
  ]
}

output "checked_domains" {
  value = local.checked_domains
}