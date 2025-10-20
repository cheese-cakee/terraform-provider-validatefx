package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewBase64Function exposes the base64 validator as a Terraform function.
func NewBase64Function() function.Function {
	return newStringValidationFunction(
		"base64",
		"Validate that a string is Base64 encoded.",
		"Returns true when the input string can be decoded from Base64.",
		validators.Base64Validator(),
	)
}
