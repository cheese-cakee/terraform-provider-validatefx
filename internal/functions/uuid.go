package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewUUIDFunction exposes the UUID validator as a Terraform function.
func NewUUIDFunction() function.Function {
	return newStringValidationFunction(
		"validatefx_uuid",
		"Validate that a string is an RFC 4122 UUID (versions 1-5).",
		"Returns true when the input is a UUID version 1 through 5; false otherwise.",
		validators.UUID(),
	)
}
