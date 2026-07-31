package ens

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormaliseDomain(t *testing.T) {
	// ENSIP-15 rejects empty labels, so inputs that resolve to one (".", ".eth",
	// "foo.eth.", ".wealdtech.eth", etc.) now error rather than silently
	// round-trip the way the old IDNA-only normaliser allowed.
	tests := []struct {
		input   string
		output  string
		wantErr bool
	}{
		{input: "", output: ""},
		{input: ".", wantErr: true},
		{input: "eth", output: "eth"},
		{input: "ETH", output: "eth"},
		{input: ".eth", wantErr: true},
		{input: ".eth.", wantErr: true},
		{input: "wealdtech.eth", output: "wealdtech.eth"},
		{input: ".wealdtech.eth", wantErr: true},
		{input: "subdomain.wealdtech.eth", output: "subdomain.wealdtech.eth"},
		{input: "*.wealdtech.eth", output: "*.wealdtech.eth"},
		{input: "omg.thetoken.eth", output: "omg.thetoken.eth"},
		// Underscore-leading labels are still valid under ENSIP-15.
		{input: "_underscore.thetoken.eth", output: "_underscore.thetoken.eth"},
		{input: "點看.eth", output: "點看.eth"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := NormaliseDomain(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			if tt.output != result {
				t.Errorf("%v => %v (expected %v)", tt.input, result, tt.output)
			}
		})
	}
}

func TestNormaliseDomainStrict(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{"", "", nil},
		{".", ".", nil},
		{"eth", "eth", nil},
		{"ETH", "eth", nil},
		{".eth", ".eth", nil},
		{".eth.", ".eth.", nil},
		{"wealdtech.eth", "wealdtech.eth", nil},
		{".wealdtech.eth", ".wealdtech.eth", nil},
		{"subdomain.wealdtech.eth", "subdomain.wealdtech.eth", nil},
		{"*.wealdtech.eth", "*.wealdtech.eth", nil},
		{"omg.thetoken.eth", "omg.thetoken.eth", nil},
		{"_underscore.thetoken.eth", "", errors.New("idna: disallowed rune U+005F")},
		{"點看.eth", "點看.eth", nil},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := NormaliseDomainStrict(tt.input)
			if tt.err != nil {
				if err == nil {
					t.Fatalf("missing expected error")
				}
				if tt.err.Error() != err.Error() {
					t.Errorf("unexpected error value %v", err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error %v", err)
				}
				if tt.output != result {
					t.Errorf("%v => %v (expected %v)", tt.input, result, tt.output)
				}
			}
		})
	}
}

func TestTld(t *testing.T) {
	// Under ENSIP-15, inputs with leading dots (".eth", ".wealdtech.eth", ".")
	// produce an empty label and NormaliseDomain errors. Tld currently returns
	// "" when normalisation fails (NormaliseDomain returns "" on error). Pre-
	// ENSIP-15 behaviour returned "eth" for ".eth" and ".wealdtech.eth"; that
	// has been corrected.
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{".", ""},
		{"eth", "eth"},
		{"ETH", "eth"},
		{".eth", ""},
		{"wealdtech.eth", "eth"},
		{".wealdtech.eth", ""},
		{"subdomain.wealdtech.eth", "eth"},
	}

	for _, tt := range tests {
		result := Tld(tt.input)
		if tt.output != result {
			t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, result, tt.output)
		}
	}
}

func TestDomainPart(t *testing.T) {
	tests := []struct {
		input  string
		part   int
		output string
		err    bool
	}{
		{"", 1, "", false},
		{"", 2, "", true},
		{"", -1, "", false},
		{"", -2, "", true},
		// ENSIP-15: a bare "." is an empty label and now errors.
		{".", 1, "", true},
		{".", 2, "", true},
		{".", 3, "", true},
		{".", -1, "", true},
		{".", -2, "", true},
		{".", -3, "", true},
		{"ETH", 1, "eth", false},
		{"ETH", 2, "", true},
		{"ETH", -1, "eth", false},
		{"ETH", -2, "", true},
		// ENSIP-15: a leading dot produces an empty label and now errors for
		// every part index — including the negative indices that previously
		// happened to land on a non-empty label.
		{".ETH", 1, "", true},
		{".ETH", 2, "", true},
		{".ETH", 3, "", true},
		{".ETH", -1, "", true},
		{".ETH", -2, "", true},
		{".ETH", -3, "", true},
		{"wealdtech.eth", 1, "wealdtech", false},
		{"wealdtech.eth", 2, "eth", false},
		{"wealdtech.eth", 3, "", true},
		{"wealdtech.eth", -1, "eth", false},
		{"wealdtech.eth", -2, "wealdtech", false},
		{"wealdtech.eth", -3, "", true},
		{".wealdtech.eth", 1, "", true},
		{".wealdtech.eth", 2, "", true},
		{".wealdtech.eth", 3, "", true},
		{".wealdtech.eth", 4, "", true},
		{".wealdtech.eth", -1, "", true},
		{".wealdtech.eth", -2, "", true},
		{".wealdtech.eth", -3, "", true},
		{".wealdtech.eth", -4, "", true},
		{"subdomain.wealdtech.eth", 1, "subdomain", false},
		{"subdomain.wealdtech.eth", 2, "wealdtech", false},
		{"subdomain.wealdtech.eth", 3, "eth", false},
		{"subdomain.wealdtech.eth", 4, "", true},
		{"subdomain.wealdtech.eth", -1, "eth", false},
		{"subdomain.wealdtech.eth", -2, "wealdtech", false},
		{"subdomain.wealdtech.eth", -3, "subdomain", false},
		{"subdomain.wealdtech.eth", -4, "", true},
		{"a.b.c", 1, "a", false},
		{"a.b.c", 2, "b", false},
		{"a.b.c", 3, "c", false},
		{"a.b.c", 4, "", true},
		{"a.b.c", -1, "c", false},
		{"a.b.c", -2, "b", false},
		{"a.b.c", -3, "a", false},
		{"a.b.c", -4, "", true},
	}

	for _, tt := range tests {
		result, err := DomainPart(tt.input, tt.part)
		if err != nil && !tt.err {
			t.Errorf("Failure: %v, %v => error (unexpected)\n", tt.input, tt.part)
		}
		if err == nil && tt.err {
			t.Errorf("Failure: %v, %v => no error (unexpected)\n", tt.input, tt.part)
		}
		if tt.output != result {
			t.Errorf("Failure: %v, %v => %v (expected %v)\n", tt.input, tt.part, result, tt.output)
		}
	}
}

func TestUnqualifiedName(t *testing.T) {
	tests := []struct {
		domain string
		root   string
		name   string
		err    error
	}{
		{
			domain: "",
			root:   "",
			name:   "",
		},
		{
			domain: "wealdtech.eth",
			root:   "eth",
			name:   "wealdtech",
		},
	}

	for i, test := range tests {
		name, err := UnqualifiedName(test.domain, test.root)
		if test.err != nil {
			assert.Equal(t, test.err, err, fmt.Sprintf("Incorrect error at test %d", i))
		} else {
			require.Nil(t, err, fmt.Sprintf("Unexpected error at test %d", i))
			assert.Equal(t, test.name, name, fmt.Sprintf("Incorrect result at test %d", i))
		}
	}
}
