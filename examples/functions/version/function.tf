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
  provider_version = provider::validatefx::version()
}

output "validatefx_version" {
  value = local.provider_version
}
