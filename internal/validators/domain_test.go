package validators

import (
	"context"
	"strings"
	"testing"

	frameworkdiag "github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestDomainValidator(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		val         types.String
		expectError bool
	}{
		// Valid domains
		"valid simple domain": {
			val:         types.StringValue("example.com"),
			expectError: false,
		},
		"valid subdomain": {
			val:         types.StringValue("www.example.com"),
			expectError: false,
		},
		"valid deep subdomain": {
			val:         types.StringValue("api.v1.example.com"),
			expectError: false,
		},
		"valid domain with numbers": {
			val:         types.StringValue("test123.example.com"),
			expectError: false,
		},
		"valid domain with hyphens": {
			val:         types.StringValue("my-domain.example-site.com"),
			expectError: false,
		},
		"valid FQDN with trailing dot": {
			val:         types.StringValue("example.com."),
			expectError: false,
		},
		"valid single character labels": {
			val:         types.StringValue("a.b.c"),
			expectError: false,
		},
		"valid long domain": {
			val:         types.StringValue("very-long-subdomain-name.another-long-subdomain.example.com"),
			expectError: false,
		},

		// Invalid domains
		"empty string": {
			val:         types.StringValue(""),
			expectError: false, // Empty strings are allowed (handled by required attribute)
		},
		"just a dot": {
			val:         types.StringValue("."),
			expectError: true,
		},
		"starts with dot": {
			val:         types.StringValue(".example.com"),
			expectError: true,
		},
		"ends with dot dot": {
			val:         types.StringValue("example.com.."),
			expectError: true,
		},
		"double dots": {
			val:         types.StringValue("example..com"),
			expectError: true,
		},
		"starts with hyphen": {
			val:         types.StringValue("-example.com"),
			expectError: true,
		},
		"ends with hyphen": {
			val:         types.StringValue("example-.com"),
			expectError: true,
		},
		"label starts with hyphen": {
			val:         types.StringValue("sub.-example.com"),
			expectError: true,
		},
		"label ends with hyphen": {
			val:         types.StringValue("sub.example-.com"),
			expectError: true,
		},
		"no TLD": {
			val:         types.StringValue("example"),
			expectError: true, // Single label domains are typically not valid
		},
		"numeric TLD": {
			val:         types.StringValue("example.123"),
			expectError: true,
		},
		"contains spaces": {
			val:         types.StringValue("example .com"),
			expectError: true,
		},
		"contains invalid characters": {
			val:         types.StringValue("example@.com"),
			expectError: true,
		},
		"contains underscore": {
			val:         types.StringValue("example_test.com"),
			expectError: true,
		},
		"too long domain": {
			val:         types.StringValue("a" + strings.Repeat(".very-long-label-that-exceeds-normal-limits", 10) + ".com"),
			expectError: true,
		},
		"label too long": {
			val:         types.StringValue(strings.Repeat("a", 64) + ".com"),
			expectError: true,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			request := frameworkvalidator.StringRequest{
				Path:        path.Root("test"),
				ConfigValue: testCase.val,
			}
			response := &frameworkvalidator.StringResponse{}

			Domain().ValidateString(context.Background(), request, response)

			if !response.Diagnostics.HasError() && testCase.expectError {
				t.Fatalf("expected error, but got none")
			}

			if response.Diagnostics.HasError() && !testCase.expectError {
				t.Fatalf("expected no error, but got: %v", response.Diagnostics)
			}
		})
	}
}

func TestDomainValidatorHandlesNullAndUnknown(t *testing.T) {
	t.Parallel()

	testCases := map[string]frameworkvalidator.StringRequest{
		"null": {
			Path:        path.Root("test"),
			ConfigValue: types.StringNull(),
		},
		"unknown": {
			Path:        path.Root("test"),
			ConfigValue: types.StringUnknown(),
		},
	}

	for name, req := range testCases {
		name, req := name, req
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			resp := &frameworkvalidator.StringResponse{}

			Domain().ValidateString(context.Background(), req, resp)

			if resp.Diagnostics.HasError() {
				t.Fatalf("expected no diagnostics for %s, got: %v", name, resp.Diagnostics)
			}
		})
	}
}

func TestDomainValidatorErrorMessages(t *testing.T) {
	t.Parallel()

	validator := Domain()
	req := frameworkvalidator.StringRequest{
		Path:        path.Root("domain"),
		ConfigValue: types.StringValue("invalid..domain"),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(context.Background(), req, resp)

	if !resp.Diagnostics.HasError() {
		t.Fatalf("expected diagnostics for invalid domain")
	}

	diagnostic := resp.Diagnostics[0]

	if diagnostic.Severity() != frameworkdiag.SeverityError {
		t.Fatalf("expected error diagnostic, got severity: %s", diagnostic.Severity())
	}

	if diagnostic.Summary() != "Invalid Domain" {
		t.Fatalf("unexpected diagnostic summary: %s", diagnostic.Summary())
	}

	expectedDetail := `Value "invalid..domain" is not a valid domain name`
	if diagnostic.Detail() != expectedDetail {
		t.Fatalf("unexpected diagnostic detail: %s", diagnostic.Detail())
	}
}

func TestDomainValidatorDescription(t *testing.T) {
	t.Parallel()

	validator := Domain()
	ctx := context.Background()

	expectedDescription := "value must be a valid domain name"

	if validator.Description(ctx) != expectedDescription {
		t.Fatalf("unexpected description: %s", validator.Description(ctx))
	}

	if validator.MarkdownDescription(ctx) != expectedDescription {
		t.Fatalf("unexpected markdown description: %s", validator.MarkdownDescription(ctx))
	}
}
