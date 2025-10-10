terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = "0.0.0"
    }
  }
}

provider "validatefx" {}

output "integration_smoke" {
  value = "provider configured"
}
