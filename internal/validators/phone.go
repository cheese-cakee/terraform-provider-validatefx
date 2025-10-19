package validators

import (
	"context"
	"fmt"
	"regexp"

	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure Phone implements frameworkvalidator.String
var _ frameworkvalidator.String = Phone()

// Phone returns a schema.String validator which enforces E.164 phone number format.
func Phone() frameworkvalidator.String {
	return phoneValidator{}
}

type phoneValidator struct{}

func (phoneValidator) Description(_ context.Context) string {
	return "value must be a valid phone number in E.164 format"
}

func (phoneValidator) MarkdownDescription(_ context.Context) string {
	return "value must be a valid phone number in E.164 format"
}

func (phoneValidator) ValidateString(_ context.Context, req frameworkvalidator.StringRequest, resp *frameworkvalidator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if value == "" {
		return
	}

	re := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	if !re.MatchString(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Phone Number",
			fmt.Sprintf("Value %q is not a valid E.164 phone number", value),
		)
	}
}
