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

  json_payloads = [
    "{\"key\": \"value\"}",
    "{\"invalid\":",
    "[]",
  ]

  semver_values = [
    "1.0.0",
    "v1.0.0",
    "1.0",
  ]

  ip_values = [
    "127.0.0.1",
    "::1",
    "999.999.999.999",
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

  json_results = [
    for value in local.json_payloads : {
      value = value
      valid = provider::validatefx::json(value)
    }
  ]

  semver_results = [
    for value in local.semver_values : {
      value = value
      valid = provider::validatefx::semver(value)
    }
  ]

  ip_results = [
    for value in local.ip_values : {
      value = value
      valid = provider::validatefx::ip(value)
    }
  ]

  # Assert function tests
  assert_email_valid = provider::validatefx::assert(
    provider::validatefx::email("alice@example.com"),
    "Email validation failed!"
  )

  assert_uuid_valid = provider::validatefx::assert(
    provider::validatefx::uuid("d9428888-122b-11e1-b85c-61cd3cbb3210"),
    "UUID validation failed!"
  )

  assert_custom_condition = provider::validatefx::assert(
    length("test") == 4,
    "String length assertion failed!"
  )
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

output "validatefx_json" {
  value = local.json_results
}

output "validatefx_semver" {
  value = local.semver_results
}

output "validatefx_ip" {
  value = local.ip_results
}

output "validatefx_assert" {
  value = {
    email_check      = local.assert_email_valid
    uuid_check       = local.assert_uuid_valid
    custom_condition = local.assert_custom_condition
  }
}
