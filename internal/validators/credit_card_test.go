package validators

import (
	"context"
	"testing"

	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestCreditCardValidator(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name        string
		input       string
		expectError bool
	}{
		// Valid credit card numbers (using Luhn algorithm)
		{
			name:        "Valid Visa",
			input:       "4532015112830366", // Valid Visa test number
			expectError: false,
		},
		{
			name:        "Valid MasterCard",
			input:       "5555555555554444", // Valid MasterCard test number
			expectError: false,
		},
		{
			name:        "Valid American Express",
			input:       "378282246310005", // Valid Amex test number
			expectError: false,
		},
		{
			name:        "Valid with spaces",
			input:       "4532 0151 1283 0366", // Valid Visa with spaces
			expectError: false,
		},
		{
			name:        "Valid with hyphens",
			input:       "4532-0151-1283-0366", // Valid Visa with hyphens
			expectError: false,
		},

		// Invalid credit card numbers
		{
			name:        "Invalid Luhn checksum",
			input:       "4532015112830367", // Invalid checksum (last digit changed)
			expectError: true,
		},
		{
			name:        "Too short",
			input:       "123456789012", // Only 12 digits
			expectError: true,
		},
		{
			name:        "Too long",
			input:       "12345678901234567890", // 20 digits
			expectError: true,
		},
		{
			name:        "Contains letters",
			input:       "4532015112830abc",
			expectError: true,
		},
		{
			name:        "Contains special characters",
			input:       "4532015112830366!",
			expectError: true,
		},
		{
			name:        "All zeros",
			input:       "0000000000000000",
			expectError: true,
		},

		// Edge cases
		{
			name:        "Empty string",
			input:       "",
			expectError: false, // Empty strings are allowed (handled upstream)
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			validator := CreditCard()

			req := frameworkvalidator.StringRequest{
				ConfigValue: types.StringValue(tc.input),
			}
			resp := &frameworkvalidator.StringResponse{}

			validator.ValidateString(ctx, req, resp)

			hasErrors := resp.Diagnostics.HasError()
			if hasErrors != tc.expectError {
				t.Errorf("Expected error: %v, got error: %v. Diagnostics: %v",
					tc.expectError, hasErrors, resp.Diagnostics.Errors())
			}
		})
	}
}

func TestCreditCardValidator_NullAndUnknown(t *testing.T) {
	ctx := context.Background()
	validator := CreditCard()

	// Test null value
	req := frameworkvalidator.StringRequest{
		ConfigValue: types.StringNull(),
	}
	resp := &frameworkvalidator.StringResponse{}

	validator.ValidateString(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("Expected no error for null value, got: %v", resp.Diagnostics.Errors())
	}

	// Test unknown value
	req = frameworkvalidator.StringRequest{
		ConfigValue: types.StringUnknown(),
	}
	resp = &frameworkvalidator.StringResponse{}

	validator.ValidateString(ctx, req, resp)

	if resp.Diagnostics.HasError() {
		t.Errorf("Expected no error for unknown value, got: %v", resp.Diagnostics.Errors())
	}
}

func TestLuhnCheck(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid Luhn - Visa",
			input:    "4532015112830366",
			expected: true,
		},
		{
			name:     "Valid Luhn - MasterCard",
			input:    "5555555555554444",
			expected: true,
		},
		{
			name:     "Invalid Luhn",
			input:    "4532015112830367",
			expected: false,
		},
		{
			name:     "Single digit valid",
			input:    "0",
			expected: true,
		},
		{
			name:     "Single digit invalid",
			input:    "1",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := luhnCheck(tc.input)
			if result != tc.expected {
				t.Errorf("luhnCheck(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
