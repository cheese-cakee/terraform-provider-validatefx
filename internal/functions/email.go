package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewEmailFunction exposes the email validator as a Terraform function.
func NewEmailFunction() function.Function {
	return newStringValidationFunction(
		"email",
		"Validate that a string is an RFC 5322 compliant email address.",
		"Returns true when the input is a valid email address and false otherwise.",
		validators.Email(),
	)
}
