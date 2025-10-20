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
  emails = [
    "alice@example.com",
    "bad-email",
  ]

  uuids = [
    "d9428888-122b-11e1-b85c-61cd3cbb3210",
    "not-a-uuid",
  ]

  base64_values = [
    "U29sdmVkIQ==",
    "invalid base64",
  ]

  credit_cards = [
    "4532015112830366",
    "4532015112830367",
  ]

  domains = [
    "example.com",
    "invalid..domain",
  ]

  email_results = [
    for value in local.emails : {
      value = value
      valid = provider::validatefx::email(value)
    }
  ]

  uuid_results = [
    for value in local.uuids : {
      value = value
      valid = provider::validatefx::uuid(value)
    }
  ]

  base64_results = [
    for value in local.base64_values : {
      value = value
      valid = provider::validatefx::base64(value)
    }
  ]

  credit_card_results = [
    for value in local.credit_cards : {
      value = value
      valid = provider::validatefx::credit_card(value)
    }
  ]

  domain_results = [
    for value in local.domains : {
      value = value
      valid = provider::validatefx::domain(value)
    }
  ]
}

output "validatefx_email" {
  value = local.email_results
}

output "validatefx_uuid" {
  value = local.uuid_results
}

output "validatefx_base64" {
  value = local.base64_results
}

output "validatefx_credit_card" {
  value = local.credit_card_results
}

output "validatefx_domain" {
  value = local.domain_results
}
