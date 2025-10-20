terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = "0.0.1"
    }
  }
}

provider "validatefx" {}

variable "test_email" {
  type    = string
  default = "user@example.com"
}

variable "bad_email" {
  type    = string
  default = "not-an-email"
}

locals {
  # Test 1: Valid email should pass
  test1 = provider::validatefx::assert(
    provider::validatefx::email(var.test_email),
    "Email validation failed for ${var.test_email}"
  )

  # Test 2: Simple condition that passes
  test2 = provider::validatefx::assert(
    5 > 3,
    "Math check failed!"
  )

  # Test 3: String length check
  test3 = provider::validatefx::assert(
    length("test") == 4,
    "String length check failed!"
  )
}

output "assert_tests" {
  value = {
    test1_email_validation = local.test1
    test2_math_check       = local.test2
    test3_string_length    = local.test3
  }
  description = "All assert function tests should return true"
}

# Uncomment to test failure:
# locals {
#   test_fail = provider::validatefx::assert(
#     provider::validatefx::email(var.bad_email),
#     "This email is invalid: ${var.bad_email}"
#   )
# }
