package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewCIDRFunction exposes the CIDR validator as a Terraform function.
func NewCIDRFunction() function.Function {
	return newStringValidationFunction(
		"cidr",
		"Validate that a string is an IPv4 or IPv6 CIDR block.",
		"Returns true when the input is a valid CIDR block and false otherwise.",
		validators.CIDR(),
	)
}
