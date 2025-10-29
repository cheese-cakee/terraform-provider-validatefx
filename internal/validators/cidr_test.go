package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestCIDRValidatorValid(t *testing.T) {
	t.Parallel()

	validator := CIDR()

	tests := map[string]string{
		"ipv4": "192.168.0.0/24",
		"ipv6": "2001:db8::/32",
	}

	for name, value := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := frameworkvalidator.StringRequest{
				Path:        path.Root("cidr"),
				ConfigValue: types.StringValue(value),
			}
			resp := &frameworkvalidator.StringResponse{}

			validator.ValidateString(context.Background(), req, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("expected no error diagnostics, got: %v", resp.Diagnostics)
			}
		})
	}
}

func TestCIDRValidatorInvalid(t *testing.T) {
	t.Parallel()

	validator := CIDR()

	tests := map[string]string{
		"bad_format":             "not-a-cidr",
		"bad_prefix":             "300.1.1.1/24",
		"bad_mask":               "192.168.0.0/99",
		"ipv6_bad_mask":          "2001:db8::/129",
		"missing_mask_separator": "192.168.0.0",
	}

	for name, value := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			req := frameworkvalidator.StringRequest{
				Path:        path.Root("cidr"),
				ConfigValue: types.StringValue(value),
			}
			resp := &frameworkvalidator.StringResponse{}

			validator.ValidateString(context.Background(), req, resp)

			if !resp.Diagnostics.HasError() {
				t.Fatalf("expected error diagnostics for %q", value)
			}

			diag := resp.Diagnostics[0]
			if diag.Severity() != frameworkdiag.SeverityError {
				t.Fatalf("expected error severity, got %s", diag.Severity())
			}

			if diag.Summary() != "Invalid CIDR" && diag.Summary() != "Invalid CIDR Mask" {
				t.Fatalf("unexpected diagnostic summary: %s", diag.Summary())
			}
		})
	}
}

func TestCIDRValidatorHandlesNullUnknown(t *testing.T) {
	t.Parallel()

	validator := CIDR()

	testCases := map[string]frameworkvalidator.StringRequest{
		"null": {
			Path:        path.Root("cidr"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("cidr"),
			ConfigValue: types.StringUnknown(),
		},
		"empty": {
			Path:        path.Root("cidr"),
			ConfigValue: types.StringValue(""),
		},
	}

	for name, req := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resp := &frameworkvalidator.StringResponse{}
			validator.ValidateString(context.Background(), req, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("expected no diagnostics for %s, got %v", name, resp.Diagnostics)
			}
		})
	}
}
