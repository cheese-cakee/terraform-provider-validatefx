package functions

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestSetProviderVersion(t *testing.T) {
	t.Parallel()

	SetProviderVersion("1.2.3")

	if !providerVersion.Equal(basetypes.NewStringValue("1.2.3")) {
		t.Fatalf("expected version 1.2.3, got %s", providerVersion.ValueString())
	}

	SetProviderVersion("")

	if !providerVersion.Equal(basetypes.NewStringValue("dev")) {
		t.Fatalf("expected fallback dev version, got %s", providerVersion.ValueString())
	}
}

func TestVersionFunctionRun(t *testing.T) {
	t.Parallel()

	SetProviderVersion("9.9.9")

	fn := NewVersionFunction()

	resp := &function.RunResponse{}
	fn.Run(context.Background(), function.RunRequest{}, resp)

	if resp.Error != nil {
		t.Fatalf("unexpected error: %s", resp.Error)
	}

	result := resp.Result.Value()

	strVal, ok := result.(basetypes.StringValue)
	if !ok {
		t.Fatalf("unexpected result type %T", result)
	}

	if strVal.ValueString() != "9.9.9" {
		t.Fatalf("expected version 9.9.9, got %s", strVal.ValueString())
	}
}
