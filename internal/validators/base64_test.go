package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestBase64ValidatorValid(t *testing.T) {
	t.Parallel()
	validator := Base64Validator()

	req := frameworkvalidator.StringRequest{
		Path:        path.Root("base64"),
		ConfigValue: types.StringValue("dGVzdGluZw=="), // base64 string for testing
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no diagnostics for valid base64, got: %v", resp.Diagnostics)
	}
}

func TestBase64ValidatorInvalid(t *testing.T) {
	t.Parallel()

	validator := Base64Validator()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("base64"),
		ConfigValue: types.StringValue("not_a_valid_base64_string"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid base64 string")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}

	if diagnostic.Summary() != "Invalid base64 string" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}
}

func TestBase64ValidatorInvalidHandlesEmptyAndNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("base64"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("base64"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("base64"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := Base64Validator()

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

