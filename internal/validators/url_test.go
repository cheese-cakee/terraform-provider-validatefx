package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func runURLValidation(v validator.String, value types.String) diag.Diagnostics {
	req := validator.StringRequest{
		Path:        path.Root("url"),
		ConfigValue: value,
	}

	resp := &validator.StringResponse{}
	v.ValidateString(context.Background(), req, resp)
	return resp.Diagnostics
}

func TestURLValidatorValid(t *testing.T) {
	t.Parallel()

	tests := []string{
		"https://example.com",
		"http://example.org/path",
		"https://sub.domain.com/path?query=1#fragment",
	}

	v := URL()

	for _, tc := range tests {
		diags := runURLValidation(v, types.StringValue(tc))
		if diags.HasError() {
			t.Fatalf("expected URL %q to be valid, got diagnostics: %v", tc, diags)
		}
	}
}

func TestURLValidatorInvalid(t *testing.T) {
	t.Parallel()

	tests := []string{
		"example.com",
		"http:/broken.com",
		"https://",
		"invalid",
		"ftp://example.com",
	}

	v := URL()

	for _, tc := range tests {
		diags := runURLValidation(v, types.StringValue(tc))
		if !diags.HasError() {
			t.Fatalf("expected URL %q to be invalid", tc)
		}
	}
}

func TestURLValidatorHandlesNullUnknown(t *testing.T) {
	t.Parallel()

	v := URL()

	tests := []struct {
		name  string
		value types.String
	}{
		{"null", types.StringNull()},
		{"unknown", types.StringUnknown()},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			diags := runURLValidation(v, tc.value)
			if diags.HasError() {
				t.Fatalf("expected no diagnostics for %s value", tc.name)
			}
		})
	}
}
