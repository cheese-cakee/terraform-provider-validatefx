package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPhoneValidatorValid(t *testing.T) {
	t.Parallel()

	validator := Phone()

	validNumbers := []string{
		"+14155552671",  // US
		"+919876543210", // India
		"+442071838750", // UK
	}

	for _, num := range validNumbers {
		req := frameworkvalidator.StringRequest{
			Path:        path.Root("phone"),
			ConfigValue: types.StringValue(num),
		}
		resp := &frameworkvalidator.StringResponse{}

		validator.ValidateString(context.Background(), req, resp)

		if resp.Diagnostics.HasError() {
			t.Fatalf("expected no diagnostics for valid number %q, got: %v", num, resp.Diagnostics)
		}
	}
}

func TestPhoneValidatorInvalid(t *testing.T) {
	t.Parallel()

	validator := Phone()

	invalidNumbers := []string{
		"14155552671",       // Missing +
		"+0123456789",       // Invalid country code
		"+1234567890123456", // Too long
		"abcd12345",         // Letters
		"+-123456789",       // Invalid characters
	}

	for _, num := range invalidNumbers {
		req := frameworkvalidator.StringRequest{
			Path:        path.Root("phone"),
			ConfigValue: types.StringValue(num),
		}
		resp := &frameworkvalidator.StringResponse{}

		validator.ValidateString(context.Background(), req, resp)

		if !resp.Diagnostics.HasError() {
			t.Fatalf("expected diagnostics for invalid number %q", num)
		}

		diag := resp.Diagnostics[0]
		if diag.Severity() != frameworkdiag.SeverityError {
			t.Fatalf("expected error diagnostic, got severity: %s", diag.Severity())
		}
		if diag.Summary() != "Invalid Phone Number" {
			t.Fatalf("unexpected diagnostic summary: %s", diag.Summary())
		}
	}
}

func TestPhoneValidatorHandlesEmptyAndNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("phone"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("phone"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("phone"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := Phone()

	for name, req := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resp := &frameworkvalidator.StringResponse{}
			validator.ValidateString(context.Background(), req, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("expected no diagnostics for case %q, got: %v", name, resp.Diagnostics)
			}
		})
	}
}
