package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewJSONFunction exposes the JSON validator as a Terraform function.
func NewJSONFunction() function.Function {
	return newStringValidationFunction(
		"json",
		"Validate that a string decodes to a JSON object.",
		"Returns true when the input string decodes to a JSON object.",
		validators.JSON(),
	)
}
