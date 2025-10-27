package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewPhoneFunction exposes the phone validator as a Terraform function.
func NewPhoneFunction() function.Function {
	return newStringValidationFunction(
		"phone",
		"Validate that a string is an E.164 compliant phone number.",
		"Returns true when the input matches the E.164 phone number format and false otherwise.",
		validators.Phone(),
	)
}
