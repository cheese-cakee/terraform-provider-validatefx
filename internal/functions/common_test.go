package functions

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

func TestStringValidationFunctionNonStringArguments(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	fn := newStringValidationFunction(
		"test",
		"Test",
		"Test",
		validators.Email(),
	)

	run := func(value attr.Value) *function.RunResponse {
		args := function.NewArgumentsData([]attr.Value{value})
		resp := &function.RunResponse{}
		fn.Run(ctx, function.RunRequest{Arguments: args}, resp)
		return resp
	}

	t.Run("number", func(t *testing.T) {
		t.Parallel()

		resp := run(basetypes.NewInt64Value(42))
		if resp.Error == nil {
			t.Fatalf("expected error for non-string argument")
		}

		if resp.Error.FunctionArgument != nil {
			t.Fatalf("expected missing argument index, got: %v", *resp.Error.FunctionArgument)
		}
	})

	t.Run("list", func(t *testing.T) {
		t.Parallel()

		listVal := basetypes.NewListValueMust(basetypes.StringType{}, []attr.Value{basetypes.NewStringValue("test")})
		resp := run(listVal)
		if resp.Error == nil {
			t.Fatalf("expected error for list argument")
		}
	})

	t.Run("null", func(t *testing.T) {
		t.Parallel()

		resp := run(basetypes.NewStringNull())
		if resp.Error != nil {
			t.Fatalf("unexpected error for null string: %s", resp.Error)
		}

		value := resp.Result.Value()
		boolVal, ok := value.(basetypes.BoolValue)
		if !ok {
			t.Fatalf("unexpected result value type: %T", value)
		}

		if !boolVal.IsUnknown() {
			t.Fatalf("expected unknown result for null input")
		}
	})
}
