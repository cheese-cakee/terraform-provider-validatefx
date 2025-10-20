terraform {
  required_providers {
    validatefx = {
      source  = "The-DevOps-Daily/validatefx"
      version = ">= 0.0.1"
    }
  }
}

provider "validatefx" {}

variable "email" {
  type    = string
  default = "user@example.com"
}

variable "user_age" {
  type    = number
  default = 25
}

locals {
  # Example 1: Validate email with custom error message
  email_validation = provider::validatefx::assert(
    provider::validatefx::email(var.email),
    "Invalid email address provided!"
  )

  # Example 2: Validate age with custom error message
  age_validation = provider::validatefx::assert(
    var.user_age >= 18,
    "User must be at least 18 years old!"
  )

  # Example 3: Multiple validation checks
  validations = {
    email_check = provider::validatefx::assert(
      provider::validatefx::email(var.email),
      "The email '${var.email}' is not valid. Please provide a valid email address."
    )
    age_check = provider::validatefx::assert(
      var.user_age >= 18 && var.user_age <= 120,
      "Age must be between 18 and 120. Provided: ${var.user_age}"
    )
  }
}

output "validation_results" {
  value = {
    email = local.email_validation
    age   = local.age_validation
  }
  description = "Results of validation checks"
}

output "all_validations" {
  value       = local.validations
  description = "All validation checks"
}
