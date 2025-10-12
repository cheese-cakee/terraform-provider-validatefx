package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestUUIDValidatorValid(t *testing.T) {
	t.Parallel()

	validator := UUID()
	testCases := map[string]string{
		"uuid_v1": "d9428888-122b-11e1-b85c-61cd3cbb3210",
		"uuid_v3": "f47ac10b-58cc-3ff9-8f5d-fc45d6d6f8d7",
		"uuid_v4": "6ba7b810-9dad-11d1-80b4-00c04fd430c8",
		"uuid_v5": "6fa459ea-ee8a-5ca4-894e-db77e160355e",
	}

	for name, value := range testCases {
		value := value
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := frameworkvalidator.StringRequest{
				Path:        path.Root("id"),
				ConfigValue: types.StringValue(value),
			}
			resp := &frameworkvalidator.StringResponse{}

			validator.ValidateString(context.Background(), req, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("expected no diagnostics for valid UUID %q, got: %v", value, resp.Diagnostics)
			}
		})
	}
}

func TestUUIDValidatorInvalid(t *testing.T) {
	t.Parallel()

	validator := UUID()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("id"),
		ConfigValue: types.StringValue("not-a-uuid"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid UUID")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}

	if diagnostic.Summary() != "Invalid UUID" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}
}

func TestUUIDValidatorHandlesEmptyAndNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("id"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("id"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("id"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := UUID()

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

func TestUUIDValidatorUnsupportedVersion(t *testing.T) {
	t.Parallel()

	validator := UUID()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("id"),
		ConfigValue: types.StringValue("00000000-0000-0000-0000-000000000000"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for UUID with unsupported version")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}

	if diagnostic.Summary() != "Unsupported UUID Version" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}
}
