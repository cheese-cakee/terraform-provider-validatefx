package functions

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
)

// IsEmail checks if the input is a valid email address
func IsEmail(input string) error {
	re := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)
	if re.MatchString(input) {
		return nil
	}
	return fmt.Errorf("invalid email")
}

// IsUUID checks if the input is a valid UUID (v1–v5)
func IsUUID(input string) error {
	re := regexp.MustCompile(`^[a-fA-F0-9]{8}\-[a-fA-F0-9]{4}\-[1-5][a-fA-F0-9]{3}\-[89abAB][a-fA-F0-9]{3}\-[a-fA-F0-9]{12}$`)
	if re.MatchString(input) {
		return nil
	}
	return fmt.Errorf("invalid UUID")
}

// IsBase64 checks if the input is valid base64
func IsBase64(input string) error {
	if input == "" {
		return fmt.Errorf("empty string")
	}
	_, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return fmt.Errorf("invalid base64: %v", err)
	}
	return nil
}

// IsCreditCard checks if the input is a valid 13–19 digit credit card number
func IsCreditCard(input string) error {
	re := regexp.MustCompile(`^\d{13,19}$`)
	if re.MatchString(input) {
		return nil
	}
	return fmt.Errorf("invalid credit card number")
}

// IsJSON checks if the input string represents a JSON object.
func IsJSON(input string) error {
	if input == "" {
		return fmt.Errorf("empty string")
	}

	var decoded any
	if err := json.Unmarshal([]byte(input), &decoded); err != nil {
		return fmt.Errorf("invalid json: %w", err)
	}

	if _, ok := decoded.(map[string]any); !ok {
		return fmt.Errorf("invalid json object")
	}

	return nil
}

// IsSemVer checks if the input string follows Semantic Versioning (e.g., 1.0.0, v2.3.4-beta)
func IsSemVer(input string) error {
	re := regexp.MustCompile(`^v?(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-[\da-zA-Z\-]+(?:\.[\da-zA-Z\-]+)*)?(?:\+[\da-zA-Z\-]+(?:\.[\da-zA-Z\-]+)*)?$`)
	if re.MatchString(input) {
		return nil
	}
	return fmt.Errorf("invalid semantic version")
}
