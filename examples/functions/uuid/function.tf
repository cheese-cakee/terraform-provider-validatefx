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
  uuids = [
    "d9428888-122b-11e1-b85c-61cd3cbb3210",
    "not-a-uuid",
  ]

  checked = [
    for id in local.uuids : {
      value = id
      valid = provider::validatefx::uuid(id)
    }
  ]
}

output "uuid_check" {
  value = local.checked
}
