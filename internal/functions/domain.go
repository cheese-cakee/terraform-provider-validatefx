package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewDomainFunction exposes the domain validator as a Terraform function.
func NewDomainFunction() function.Function {
	return newStringValidationFunction(
		"domain",
		"Validate that a string is a compliant domain name.",
		"Returns true when the input is a valid domain per RFC 1123/952 rules.",
		validators.Domain(),
	)
}
