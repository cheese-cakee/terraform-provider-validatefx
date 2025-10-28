package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var providerVersion = basetypes.NewStringValue("dev")

// SetProviderVersion updates the version string used by provider diagnostic functions.
//
// When the provider is built without an explicit version, Terraform sets the
// version string to an empty value. Normalize that case to "dev" so the
// function remains useful during local development.
func SetProviderVersion(version string) {
	if version == "" {
		version = "dev"
	}

	providerVersion = basetypes.NewStringValue(version)
}

// ProviderVersion returns the current provider version value.
func ProviderVersion() basetypes.StringValue {
	return providerVersion
}

// NewVersionFunction exposes the provider version as a Terraform function.
func NewVersionFunction() function.Function {
	return &versionFunction{}
}

var _ function.Function = (*versionFunction)(nil)

type versionFunction struct{}

func (versionFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "version"
}

func (versionFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Return the provider version string.",
		MarkdownDescription: "Returns the provider build version as a string for diagnostics and debugging.",
		Return:              function.StringReturn{},
	}
}

func (versionFunction) Run(_ context.Context, _ function.RunRequest, resp *function.RunResponse) {
	resp.Result = function.NewResultData(providerVersion)
}
