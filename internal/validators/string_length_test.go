package validators

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func TestStringLengthValidator(t *testing.T) {
	ctx := context.Background()
	min := 2
	max := 5

	tests := []struct {
		name    string
		val     types.String
		wantErr bool
	}{
		{"null value", types.StringNull(), false},
		{"unknown value", types.StringUnknown(), false},
		{"too short", types.StringValue("a"), true},
		{"within range", types.StringValue("abcd"), false},
		{"too long", types.StringValue("abcdefg"), true},
		{"multi-byte emoji", types.StringValue("😊😊"), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := validator.StringRequest{
				ConfigValue: tt.val,
				Path:        path.Root("test_attr"),
			}
			var resp validator.StringResponse
			v := NewStringLengthValidator(&min, &max)
			v.ValidateString(ctx, req, &resp)

			if resp.Diagnostics.HasError() != tt.wantErr {
				t.Errorf("Test %s failed: expected error=%v, got %v", tt.name, tt.wantErr, resp.Diagnostics.HasError())
			}
		})
	}
}
