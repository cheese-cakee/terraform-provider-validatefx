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
  payloads = {
    valid_object   = "{\"hello\": \"world\"}"
    invalid_syntax = "{\"missing\": \"quote}"
    array_payload  = "[1, 2, 3]"
  }

  validation_results = {
    for name, payload in local.payloads : name => provider::validatefx::json(payload)
  }

  only_object = provider::validatefx::assert(
    provider::validatefx::json(local.payloads.valid_object),
    "payload must be a JSON object"
  )
}

output "json_validation" {
  value = local.validation_results
}

output "assert_json" {
  value = local.only_object
}
