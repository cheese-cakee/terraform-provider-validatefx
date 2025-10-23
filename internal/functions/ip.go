package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewIPFunction exposes the IP validator as a Terraform function.
func NewIPFunction() function.Function {
	return newStringValidationFunction(
		"ip",
		"Validate that a string is a valid IPv4 or IPv6 address.",
		"Returns true when the input string parses as a valid IPv4 or IPv6 address.",
		validators.IP(),
	)
}
