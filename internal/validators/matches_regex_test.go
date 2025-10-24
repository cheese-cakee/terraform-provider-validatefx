package validators

import (
	"context"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestMatchesRegexValidatorValid(t *testing.T) {
	t.Parallel()

	validator := MatchesRegex(`^[a-z0-9_]+$`)
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("username"),
		ConfigValue: types.StringValue("user_123"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if resp.Diagnostics.HasError() {
		t.Fatalf("expected no diagnostics for value matching pattern, got: %v", resp.Diagnostics)
	}
}

func TestMatchesRegexValidatorMismatch(t *testing.T) {
	t.Parallel()

	validator := MatchesRegex(`^[a-z0-9_]+$`)
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("username"),
		ConfigValue: types.StringValue("User-123"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for non-matching value")
	}

	diag := resp.Diagnostics[0]

	if diag.Summary() != "Regex Mismatch" {
		t.Fatalf("unexpected diagnostic summary: %s", diag.Summary())
	}

	if diag.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error severity, got %s", diag.Severity())
	}
}

func TestMatchesRegexValidatorInvalidPattern(t *testing.T) {
	t.Parallel()

	validator := MatchesRegex(`[`)
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("username"),
		ConfigValue: types.StringValue("hello"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid regex pattern")
	}

	diag := resp.Diagnostics[0]

	if diag.Summary() != "Invalid Regex Pattern" {
		t.Fatalf("unexpected diagnostic summary: %s", diag.Summary())
	}
}

func TestMatchesRegexHandlesEmptyNullUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"empty": {
			Path:        path.Root("username"),
			ConfigValue: types.StringValue(""),
		},
		"null": {
			Path:        path.Root("username"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("username"),
			ConfigValue: types.StringUnknown(),
		},
	}

	validator := MatchesRegex(`.*`)

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
