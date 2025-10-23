package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestIPValidatorValid(t *testing.T) {
	t.Parallel()

	validator := IP()

	tests := []string{
		"127.0.0.1",
		"192.168.1.10",
		"::1",
		"2001:db8::1",
	}

	for _, value := range tests {
		t.Run(value, func(t *testing.T) {
			t.Parallel()

			req := frameworkvalidator.StringRequest{
				Path:        path.Root("ip"),
				ConfigValue: types.StringValue(value),
			}
			resp := &frameworkvalidator.StringResponse{}

			validator.ValidateString(context.Background(), req, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("expected no diagnostics for valid IP %q, got: %v", value, resp.Diagnostics)
			}
		})
	}
}

func TestIPValidatorInvalid(t *testing.T) {
	t.Parallel()

	validator := IP()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("ip"),
		ConfigValue: types.StringValue("999.999.999.999"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid IP")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}

	if diagnostic.Summary() != "Invalid IP Address" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}
}

func TestIPValidatorHandlesEmptyNullUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("ip"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("ip"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("ip"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := IP()

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
