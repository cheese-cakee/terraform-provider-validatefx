package functions

import (
	"github.com/hashicorp/terraform-plugin-framework/function"

	"github.com/The-DevOps-Daily/terraform-provider-validatefx/internal/validators"
)

// NewCreditCardFunction exposes the credit card validator as a Terraform function.
func NewCreditCardFunction() function.Function {
	return newStringValidationFunction(
		"validatefx_credit_card",
		"Validate that a string is a credit card number using the Luhn algorithm.",
		"Returns true when the input passes Luhn validation and false otherwise.",
		validators.CreditCard(),
	)
}
