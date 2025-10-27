package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewURLFunction exposes the URL validator as a Terraform function.
func NewURLFunction() function.Function {
	return newStringValidationFunction(
		"url",
		"Validate that a string is an HTTP(S) URL.",
		"Returns true when the input string is a valid HTTP or HTTPS URL including scheme and host.",
		validators.URL(),
	)
}
