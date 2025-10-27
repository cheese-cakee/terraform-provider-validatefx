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
    " user@example.com ",
    "john.doe+test@sub.domain.io",
    "user@localhost",
    "UPPER@EXAMPLE.COM",
    "",
    null,
  ]

  uuids = [
    "d9428888-122b-11e1-b85c-61cd3cbb3210",
    "D9428888-122B-11E1-B85C-61CDCBB3210",
    "00000000-0000-0000-0000-000000000000",
    "d9428888-122b-11e1-b85c-61cd3cbb3210?query=name",
    "{d9428888-122b-11e1-b85c-61cd3cbb3210}",
    "not-a-uuid",
    "",
  ]

  base64_values = [
    "U29sdmVkIQ==",
    "c3VjY2VzczEyMw==",
    " U29sdmVkIQ== ",
    "Zm9vYmFy",
    "invalid base64",
    "U29sdmVkIQ",
    "",
  ]

  credit_cards = [
    "4532015112830366",
    "4539 1488 0343 6467",
    "4111-1111-1111-1111",
    "6011111111111117",
    "1234567890123456",
    "0000-0000-0000-0000",
    "",
  ]

  phone_numbers = [
    "+14155552671",
    "+918888888888",
    "+447911123456",
    "+33123456789",
    "14155552671",
    "+00123456789",
    "+123456789012345",
    "+1234567890123456",
    "",
  ]

  url_values = [
    "https://example.com",
    "http://example.org/path?query=1#frag",
    "ftp://example.com",
    "HTTP://legacy.example.com",
    "https://example.com:8443/sub",
    "https://",
    "relative/path",
    "https:// spaced.com",
    "",
  ]

  domains = [
    "example.com",
    "invalid..domain",
    "sub.domain.io",
    "with-hyphen-domain.com",
    "-leadinghyphen.com",
    "xn--bücher.de",
    "example.com.",
    "localhost",
  ]

  json_payloads = [
    "{\"key\": \"value\"}",
    "{\"nested\":{\"foo\":\"bar\"}}",
    "{\"array\":[1,2,3]}",
    "{\"invalid\":",
    "[]",
    "null",
    "",
  ]

  semver_values = [
    "1.0.0",
    "v1.0.0",
    "1.0",
    "1.0.0-alpha+001",
    "0.0.0",
    "latest",
    "1.2.3-rc.1",
    "1.2.3+build.5",
    "",
  ]

  ip_values = [
    "127.0.0.1",
    "::1",
    "999.999.999.999",
    "2001:db8::1",
    "10.0.0.256",
    "0.0.0.0",
    "abc",
  ]

  regex_samples = [
    {
      value   = "user_123"
      pattern = "^[a-z0-9_]+$"
    },
    {
      value   = "Invalid-User"
      pattern = "^[a-z0-9_]+$"
    },
    {
      value   = "abc123"
      pattern = "^abc\\d+$"
    },
    {
      value   = "ABC"
      pattern = "(?i)^abc$"
    },
    {
      value   = "abc"
      pattern = "("
    },
  ]

  all_valid_cases = [
    [true, true, true],
    [true, false],
    [true, null],
    [true, true, null],
    [],
  ]

  any_valid_cases = [
    [false, false],
    [false, true],
    [false, null, false],
    [null, null],
    [],
  ]

  email_results = [
    for value in local.emails : {
      value   = value
      trimmed = value != null ? trimspace(value) : null
      valid   = trimmed != null && trimmed != "" ? provider::validatefx::email(trimmed) : false
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
      value   = value
      trimmed = trimspace(value)
      valid   = trimmed != "" ? provider::validatefx::base64(trimmed) : provider::validatefx::base64(value)
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

  matches_regex_results = [
    for item in local.regex_samples : {
      value   = item.value
      pattern = item.pattern
      valid   = provider::validatefx::matches_regex(item.value, item.pattern)
    }
  ]

  phone_results = [
    for value in local.phone_numbers : {
      value = value
      valid = provider::validatefx::phone(value)
    }
  ]

  url_results = [
    for value in local.url_values : {
      value = value
      valid = provider::validatefx::url(value)
    }
  ]

  all_valid_results = [
    for checks in local.all_valid_cases : {
      checks = checks
      result = provider::validatefx::all_valid(checks)
    }
  ]

  any_valid_results = [
    for checks in local.any_valid_cases : {
      checks = checks
      result = provider::validatefx::any_valid(checks)
    }
  ]
  # Assert function tests
  assert_email_valid = provider::validatefx::assert(
    provider::validatefx::email("alice@example.com"),
    "Email validation failed!",
  )

  assert_uuid_valid = provider::validatefx::assert(
    provider::validatefx::uuid("d9428888-122b-11e1-b85c-61cd3cbb3210"),
    "UUID validation failed!",
  )

  assert_custom_condition = provider::validatefx::assert(
    length("test") == 4,
    "String length assertion failed!",
  )

}

# Outputs

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

output "validatefx_matches_regex" {
  value = local.matches_regex_results
}

output "validatefx_phone" {
  value = local.phone_results
}

output "validatefx_url" {
  value = local.url_results
}

output "validatefx_all_valid" {
  value = local.all_valid_results
}

output "validatefx_any_valid" {
  value = local.any_valid_results
}

output "validatefx_assert" {
  value = {
    email_check      = local.assert_email_valid
    uuid_check       = local.assert_uuid_valid
    custom_condition = local.assert_custom_condition
  }
}
