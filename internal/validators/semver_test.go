package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestSemVerValidatorValid(t *testing.T) {
	t.Parallel()

	validator := SemVer()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("semver"),
		ConfigValue: types.StringValue("1.2.3"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no diagnostics for valid semver, got: %v", resp.Diagnostics)
	}
}

func TestSemVerValidatorInvalid(t *testing.T) {
	t.Parallel()

	validator := SemVer()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("semver"),
		ConfigValue: types.StringValue("1.2"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid semver")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}

	if diagnostic.Summary() != "Invalid Semantic Version" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}
}

func TestSemVerValidatorHandlesVariants(t *testing.T) {
	t.Parallel()

	tests := map[string]string{
		"with v prefix":   "v1.0.0",
		"with prerelease": "1.0.0-alpha.1",
		"with build":      "1.0.0+build.1",
		"with both":       "v1.0.0-beta+exp.sha.5114f85",
	}

	validator := SemVer()

	for name, value := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resp := &frameworkvalidator.StringResponse{}
			validator.ValidateString(context.Background(), frameworkvalidator.StringRequest{
				Path:        path.Root("semver"),
				ConfigValue: types.StringValue(value),
			}, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("expected no diagnostics for value %q, got: %v", value, resp.Diagnostics)
			}
		})
	}
}

func TestSemVerValidatorHandlesEmptyNullUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("semver"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("semver"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("semver"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := SemVer()

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
