package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestPhoneValidator(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
	}{
		{"Valid US number", "+14155552671", false},
		{"Valid India number", "+919876543210", false},
		{"Valid UK number", "+442071838750", false},
		{"Missing +", "14155552671", true},
		{"Invalid country code", "+0123456789", true},
		{"Too long", "+1234567890123456", true},
		{"Letters", "abcd12345", true},
		{"Invalid characters", "+-123456789", true},
	}

	validator := Phone()
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &frameworkvalidator.StringResponse{
				Diagnostics: diag.Diagnostics{},
			}
			req := frameworkvalidator.StringRequest{
				ConfigValue: types.StringValue(tt.input),
			}
			validator.ValidateString(ctx, req, resp)
			if (len(resp.Diagnostics) > 0) != tt.wantErr {
				t.Errorf("Phone validator error = %v, wantErr %v", resp.Diagnostics, tt.wantErr)
			}
		})
	}
}
