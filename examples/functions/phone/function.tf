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
  numbers = [
    "+49711649000",
    "12345",
  ]

  results = [
    for number in local.numbers : {
      number = number
      valid  = provider::validatefx::phone(number)
    }
  ]
}

output "phone_validation" {
  value = local.results
}
