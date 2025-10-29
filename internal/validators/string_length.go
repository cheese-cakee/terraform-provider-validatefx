package validators

import (
	"context"
	"fmt"
	"unicode/utf8"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)


// StringLengthValidator enforces minimum and/or maximum string length.
// If min or max are nil, that side of validation is skipped.
type StringLengthValidator struct {
	Min *int
	Max *int
}

// Ensure interface compliance
var _ validator.String = (*StringLengthValidator)(nil)

// NewStringLengthValidator creates a new instance.
func NewStringLengthValidator(minLen, maxLen *int) StringLengthValidator {
	return StringLengthValidator{Min: minLen, Max: maxLen}
}

// ValidateString performs the actual validation logic.
func (v StringLengthValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// Skip validation for unknown or null values
	if req.ConfigValue.IsUnknown() || req.ConfigValue.IsNull() {
		return
	}

	value := req.ConfigValue.ValueString()
	length := utf8.RuneCountInString(value)

	if v.Min != nil && length < *v.Min {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"String Too Short",
			fmt.Sprintf("String length is %d, must be at least %d characters long.", length, *v.Min),
		)
	}

	if v.Max != nil && length > *v.Max {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"String Too Long",
			fmt.Sprintf("String length is %d, must not exceed %d characters.", length, *v.Max),
		)
	}
}

// Description returns a human-readable description of the validator.
func (v StringLengthValidator) Description(_ context.Context) string {
	switch {
	case v.Min != nil && v.Max != nil:
		return fmt.Sprintf("string length must be between %d and %d characters", *v.Min, *v.Max)
	case v.Min != nil:
		return fmt.Sprintf("string length must be at least %d characters", *v.Min)
	case v.Max != nil:
		return fmt.Sprintf("string length must be at most %d characters", *v.Max)
	default:
		return "string length validation"
	}
}

// MarkdownDescription returns a Markdown formatted description of the validator.
func (v StringLengthValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}
