package functions

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wantOk bool
	}{
		{"valid email", "test@example.com", true},
		{"missing domain", "test@", false},
		{"missing @", "test.com", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEmail(tt.input)
			if (got == nil) != tt.wantOk {
				t.Errorf("IsEmail(%q) = %v, wantOk %v", tt.input, got, tt.wantOk)
			}
		})
	}
}

func TestIsUUID(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wantOk bool
	}{
		{"valid UUID", "123e4567-e89b-12d3-a456-426614174000", true},
		{"invalid UUID", "12345", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsUUID(tt.input)
			if (got == nil) != tt.wantOk {
				t.Errorf("IsUUID(%q) = %v, wantOk %v", tt.input, got, tt.wantOk)
			}
		})
	}
}

func TestIsBase64(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wantOk bool
	}{
		{"valid base64", "dGVzdA==", true},
		{"invalid base64", "test!", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsBase64(tt.input)
			if (got == nil) != tt.wantOk {
				t.Errorf("IsBase64(%q) = %v, wantOk %v", tt.input, got, tt.wantOk)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wantOk bool
	}{
		{"valid object", `{"key": "value"}`, true},
		{"invalid syntax", `{"key":`, false},
		{"array", `[]`, false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsJSON(tt.input)
			if (got == nil) != tt.wantOk {
				t.Errorf("IsJSON(%q) = %v, wantOk %v", tt.input, got, tt.wantOk)
			}
		})
	}
}

func TestIsCreditCard(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wantOk bool
	}{
		{"valid card", "4111111111111111", true},
		{"invalid card", "1234", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsCreditCard(tt.input)
			if (got == nil) != tt.wantOk {
				t.Errorf("IsCreditCard(%q) = %v, wantOk %v", tt.input, got, tt.wantOk)
			}
		})
	}
}

func TestIsSemVer(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wantOk bool
	}{
		{"valid semver", "1.0.0", true},
		{"valid with v prefix", "v2.3.4", true},
		{"valid prerelease", "1.0.0-beta", true},
		{"invalid missing patch", "1.0", false},
		{"invalid text", "version1.0.0", false},
		{"empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSemVer(tt.input)
			if (got == nil) != tt.wantOk {
				t.Errorf("IsSemVer(%q) = %v, wantOk %v", tt.input, got, tt.wantOk)
			}
		})
	}
}
