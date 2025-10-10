package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestEmailValidatorValid(t *testing.T) {
	t.Parallel()

	validator := Email()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("email"),
		ConfigValue: types.StringValue("user@example.com"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no diagnostics for valid email, got: %v", resp.Diagnostics)
	}
}

func TestEmailValidatorInvalid(t *testing.T) {
	t.Parallel()

	validator := Email()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("email"),
		ConfigValue: types.StringValue("not-an-email"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid email")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}

	if diagnostic.Summary() != "Invalid Email Address" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}
}

func TestEmailValidatorHandlesEmptyAndNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("email"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("email"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("email"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := Email()

	for name, req := range testCases {
		req := req
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
