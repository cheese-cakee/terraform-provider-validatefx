package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &validateFXProvider{}

// validateFXProvider defines the ValidateFX Terraform provider implementation.
type validateFXProvider struct {
	version string
}

// New returns a new instance of the ValidateFX provider factory function.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &validateFXProvider{
			version: version,
		}
	}
}

func (p *validateFXProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "validatefx"
	resp.Version = p.version
}

func (p *validateFXProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *validateFXProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *validateFXProvider) Resources(ctx context.Context) []func() resource.Resource {
	return nil
}

func (p *validateFXProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return nil
}
