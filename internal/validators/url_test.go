func TestIsURL(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		wantOk bool
	}{
		{"https basic", "https://example.com", true},
		{"http with path and query", "http://example.com/path?x=1", true},
		{"subdomain", "https://sub.example.co.uk", true},
		{"localhost with port", "https://localhost:8080", true},
		{"ip address", "https://127.0.0.1", true},
		{"fragment", "https://example.com#frag", true},

		{"missing scheme", "example.com", false},
		{"unsupported scheme", "ftp://example.com", false},
		{"missing host", "http://", false},
		{"spaces", "http://not a url", false},
		{"empty", "", false},
		{"bad prefix", "://example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsURL(tt.input)
			if (got == nil) != tt.wantOk {
				t.Errorf("IsURL(%q) = %v, wantOk %v", tt.input, got, tt.wantOk)
			}
		})
	}
}