package validators

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	ErrInvalidURL = errors.New("invalid URL format")
)

type urlValidator struct{}

func URL() validator.String {
	return urlValidator{}
}

func (v urlValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	parsed, err := url.ParseRequestURI(value)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Invalid URL", ErrInvalidURL.Error()))
	}
}

func (v urlValidator) Description(ctx context.Context) string {
	return "Ensures that the string is a valid URL including scheme and host"
}

func (v urlValidator) MarkdownDescription(ctx context.Context) string {
	return "Ensures that the string is a **valid URL** including scheme (`http`, `https`, etc.) and host."
}
