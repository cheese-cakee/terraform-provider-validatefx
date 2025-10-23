package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewSemVerFunction exposes the semver validator as a Terraform function.
func NewSemVerFunction() function.Function {
	return newStringValidationFunction(
		"semver",
		"Validate that a string follows Semantic Versioning (SemVer 2.0.0).",
		"Returns true when the input string matches the SemVer 2.0.0 specification.",
		validators.SemVer(),
	)
}
