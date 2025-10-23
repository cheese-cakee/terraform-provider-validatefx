package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestJSONValidatorValidObject(t *testing.T) {
	t.Parallel()

	validator := JSON()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("json"),
		ConfigValue: types.StringValue(`{"key":"value"}`),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no diagnostics for valid JSON object, got: %v", resp.Diagnostics)
	}
}

func TestJSONValidatorInvalidSyntax(t *testing.T) {
	t.Parallel()

	validator := JSON()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("json"),
		ConfigValue: types.StringValue(`{"key":`),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid JSON syntax")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Summary() != "Invalid JSON" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}
}

func TestJSONValidatorNonObject(t *testing.T) {
	t.Parallel()

	validator := JSON()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("json"),
		ConfigValue: types.StringValue(`[]`),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for non-object JSON value")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Summary() != "Invalid JSON Object" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}
}

func TestJSONValidatorHandlesEmptyNullUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("json"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("json"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("json"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := JSON()

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
