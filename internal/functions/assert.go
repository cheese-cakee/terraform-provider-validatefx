package functions

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type assertFunction struct{}

var _ function.Function = (*assertFunction)(nil)

// NewAssertFunction returns a new instance of the assert function.
func NewAssertFunction() function.Function {
	return &assertFunction{}
}

func (f *assertFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "assert"
}

func (f *assertFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Assert a condition with a custom error message.",
		MarkdownDescription: "Validates a condition and raises an error with a custom message if the condition is false. Returns true if the condition is true.",
		Return:              function.BoolReturn{},
		Parameters: []function.Parameter{
			function.BoolParameter{
				Name:                "condition",
				AllowNullValue:      false,
				AllowUnknownValues:  true,
				Description:         "The boolean condition to validate.",
				MarkdownDescription: "The boolean condition to validate.",
			},
			function.StringParameter{
				Name:                "error_message",
				AllowNullValue:      false,
				AllowUnknownValues:  false,
				Description:         "The custom error message to display if the condition is false.",
				MarkdownDescription: "The custom error message to display if the condition is false.",
			},
		},
	}
}

func (f *assertFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var condition types.Bool
	var errorMessage types.String

	// Get the condition argument
	if err := req.Arguments.GetArgument(ctx, 0, &condition); err != nil {
		resp.Error = function.NewArgumentFuncError(0, "Invalid condition argument: "+err.Error())
		return
	}

	// Get the error message argument
	if err := req.Arguments.GetArgument(ctx, 1, &errorMessage); err != nil {
		resp.Error = function.NewArgumentFuncError(1, "Invalid error_message argument: "+err.Error())
		return
	}

	// If condition is unknown, return unknown
	if condition.IsUnknown() {
		resp.Result = function.NewResultData(types.BoolUnknown())
		return
	}

	// If condition is null, treat it as false
	if condition.IsNull() {
		resp.Error = function.NewFuncError("Assertion failed: " + errorMessage.ValueString())
		return
	}

	// Check the condition
	if !condition.ValueBool() {
		// Condition is false, raise an error with the custom message
		resp.Error = function.NewFuncError("Assertion failed: " + errorMessage.ValueString())
		return
	}

	// Condition is true, return true
	resp.Result = function.NewResultData(basetypes.NewBoolValue(true))
}
