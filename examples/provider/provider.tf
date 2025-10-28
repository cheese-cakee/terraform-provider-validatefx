terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = "~> 0.1"
    }
  }
}

provider "validatefx" {}

locals {
  email_is_valid = provider::validatefx::email("user@example.com")
}
