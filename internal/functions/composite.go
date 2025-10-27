package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type compositeValidationFunction struct {
	name        string
	summary     string
	description string
	allMustPass bool
}

func newCompositeValidationFunction(name, summary, description string, allMustPass bool) function.Function {
	return &compositeValidationFunction{
		name:        name,
		summary:     summary,
		description: description,
		allMustPass: allMustPass,
	}
}

func (f *compositeValidationFunction) Metadata(_ context.Context, _ function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = f.name
}

func (f *compositeValidationFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             f.summary,
		MarkdownDescription: f.description,
		Return:              function.BoolReturn{},
		Parameters: []function.Parameter{
			function.ListParameter{
				Name:                "checks",
				AllowNullValue:      true,
				AllowUnknownValues:  true,
				ElementType:         basetypes.BoolType{},
				Description:         "List of boolean validation results to evaluate.",
				MarkdownDescription: "List of boolean validation results to evaluate.",
			},
		},
	}
}

func (f *compositeValidationFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var checks types.List

	if err := req.Arguments.GetArgument(ctx, 0, &checks); err != nil {
		resp.Error = err
		return
	}

	if checks.IsNull() || checks.IsUnknown() {
		resp.Result = function.NewResultData(types.BoolUnknown())
		return
	}

	var bools []basetypes.BoolValue
	if diags := checks.ElementsAs(ctx, &bools, false); diags.HasError() {
		diags.AddAttributeError(
			path.Root("checks"),
			"Invalid Boolean",
			"List elements must be boolean validation results.",
		)
		resp.Error = function.FuncErrorFromDiags(ctx, diags)
		return
	}

	if len(bools) == 0 {
		resp.Result = function.NewResultData(basetypes.NewBoolValue(f.allMustPass))
		return
	}

	var encounteredUnknown bool
	for _, elem := range bools {
		if elem.IsUnknown() {
			encounteredUnknown = true
			continue
		}

		if elem.IsNull() {
			if f.allMustPass {
				resp.Result = function.NewResultData(basetypes.NewBoolValue(false))
				return
			}
			continue
		}

		value := elem.ValueBool()
		if f.allMustPass && !value {
			resp.Result = function.NewResultData(basetypes.NewBoolValue(false))
			return
		}
		if !f.allMustPass && value {
			resp.Result = function.NewResultData(basetypes.NewBoolValue(true))
			return
		}
	}

	if encounteredUnknown {
		resp.Result = function.NewResultData(types.BoolUnknown())
		return
	}

	if f.allMustPass {
		resp.Result = function.NewResultData(basetypes.NewBoolValue(true))
	} else {
		resp.Result = function.NewResultData(basetypes.NewBoolValue(false))
	}
}

func NewAllValidFunction() function.Function {
	return newCompositeValidationFunction(
		"all_valid",
		"Return true when all provided validation checks evaluate to true.",
		"Accepts a list of boolean validation results and returns true only when every element is true.",
		true,
	)
}

func NewAnyValidFunction() function.Function {
	return newCompositeValidationFunction(
		"any_valid",
		"Return true when any provided validation check evaluates to true.",
		"Accepts a list of boolean validation results and returns true when at least one element is true.",
		false,
	)
}
