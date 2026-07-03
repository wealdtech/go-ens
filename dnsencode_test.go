// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	ens "github.com/wealdtech/go-ens/v4"
)

// ---------------------------------------------------------------------------
// Happy paths.
// ---------------------------------------------------------------------------

func TestDNSEncode_Empty(t *testing.T) {
	out, err := ens.DNSEncode("")
	require.NoError(t, err)
	assert.Equal(t, []byte{0}, out)
}

func TestDNSEncode_SingleLabel(t *testing.T) {
	out, err := ens.DNSEncode("eth")
	require.NoError(t, err)
	assert.Equal(t, []byte{3, 'e', 't', 'h', 0}, out)
}

func TestDNSEncode_TwoLabels(t *testing.T) {
	out, err := ens.DNSEncode("foo.eth")
	require.NoError(t, err)
	assert.Equal(t, []byte{3, 'f', 'o', 'o', 3, 'e', 't', 'h', 0}, out)
}

func TestDNSEncode_DeepLabels(t *testing.T) {
	out, err := ens.DNSEncode("a.b.c.eth")
	require.NoError(t, err)
	assert.Equal(t, []byte{1, 'a', 1, 'b', 1, 'c', 3, 'e', 't', 'h', 0}, out)
}

func TestDNSEncode_SingleCharLabels(t *testing.T) {
	out, err := ens.DNSEncode("a.b.c")
	require.NoError(t, err)
	assert.Equal(t, []byte{1, 'a', 1, 'b', 1, 'c', 0}, out)
}

func TestDNSEncode_NumericLabel(t *testing.T) {
	out, err := ens.DNSEncode("123.eth")
	require.NoError(t, err)
	assert.Equal(t, []byte{3, '1', '2', '3', 3, 'e', 't', 'h', 0}, out)
}

func TestDNSEncode_HyphenatedLabel(t *testing.T) {
	out, err := ens.DNSEncode("foo-bar.eth")
	require.NoError(t, err)
	assert.Equal(t, []byte{7, 'f', 'o', 'o', '-', 'b', 'a', 'r', 3, 'e', 't', 'h', 0}, out)
}

// ---------------------------------------------------------------------------
// Normalization (IDNA MapForLookup should lowercase).
// ---------------------------------------------------------------------------

func TestDNSEncode_NormalizesToLowercase(t *testing.T) {
	out, err := ens.DNSEncode("FOO.ETH")
	require.NoError(t, err)
	assert.Equal(t, []byte{3, 'f', 'o', 'o', 3, 'e', 't', 'h', 0}, out)
}

func TestDNSEncode_NormalizesMixedCase(t *testing.T) {
	out, err := ens.DNSEncode("Foo.ETH")
	require.NoError(t, err)
	assert.Equal(t, []byte{3, 'f', 'o', 'o', 3, 'e', 't', 'h', 0}, out)
}

// ---------------------------------------------------------------------------
// Trailing-dot handling.
// ---------------------------------------------------------------------------

func TestDNSEncode_SingleTrailingDot(t *testing.T) {
	out, err := ens.DNSEncode("foo.eth.")
	require.NoError(t, err)
	assert.Equal(t, []byte{3, 'f', 'o', 'o', 3, 'e', 't', 'h', 0}, out)
}

// The doc comment states only ONE trailing dot is trimmed; a second produces
// an empty label and must be rejected.
func TestDNSEncode_DoubleTrailingDot(t *testing.T) {
	_, err := ens.DNSEncode("foo.eth..")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty label")
}

// ---------------------------------------------------------------------------
// Empty-label rejection (leading dot, only-dot, double-dot in middle).
// ---------------------------------------------------------------------------

func TestDNSEncode_LeadingDot(t *testing.T) {
	_, err := ens.DNSEncode(".eth")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty label")
}

func TestDNSEncode_OnlyDot(t *testing.T) {
	_, err := ens.DNSEncode(".")
	require.Error(t, err)
}

func TestDNSEncode_DoubleDotInMiddle(t *testing.T) {
	_, err := ens.DNSEncode("foo..eth")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "empty label")
}

// ---------------------------------------------------------------------------
// Label-length boundaries.
//
// RFC 1035 §3.1 caps a DNS label at 63 bytes. The top 2 bits of the length
// byte are reserved for compression-pointer flags, so a 64+ byte label cannot
// be encoded validly.
//
// The current implementation uses a 63-byte cap.
// ---------------------------------------------------------------------------

func TestDNSEncode_Label63Bytes(t *testing.T) {
	label := strings.Repeat("a", 63)
	out, err := ens.DNSEncode(label + ".eth")
	require.NoError(t, err)
	assert.Equal(t, byte(63), out[0])
	assert.Equal(t, label, string(out[1:64]))
}

func TestDNSEncode_Label64Bytes(t *testing.T) {
	label := strings.Repeat("a", 64)
	_, err := ens.DNSEncode(label + ".eth")
	require.Error(t, err, "labels above 63 bytes are invalid per RFC 1035 §3.1")
}

func TestDNSEncode_Label255Bytes(t *testing.T) {
	label := strings.Repeat("a", 255)
	_, err := ens.DNSEncode(label + ".eth")
	require.Error(t, err, "labels above 63 bytes are invalid per RFC 1035 §3.1")
}

// 256-byte label is rejected by the implementation's `> 63` check.
func TestDNSEncode_Label256Bytes(t *testing.T) {
	label := strings.Repeat("a", 256)
	_, err := ens.DNSEncode(label + ".eth")
	require.Error(t, err)
}

// ---------------------------------------------------------------------------
// Total-name-length boundaries.
//
// ENSIP-10 explicitly removes RFC 1035's 255-octet cap on the total encoded
// name, so DNSEncode must accept arbitrarily long encodings as long as each
// label individually conforms.
// ---------------------------------------------------------------------------

// 255 octets total: 63 + 63 + 63 + 61 data bytes; +4 length bytes +1 null.
func TestDNSEncode_TotalLength255(t *testing.T) {
	parts := []string{
		strings.Repeat("a", 63),
		strings.Repeat("b", 63),
		strings.Repeat("c", 63),
		strings.Repeat("d", 61),
	}
	out, err := ens.DNSEncode(strings.Join(parts, "."))
	require.NoError(t, err)
	assert.Len(t, out, 255)
}

// 256 octets — past the RFC 1035 total-length cap, but ENSIP-10 removes that
// cap, so this must still succeed.
func TestDNSEncode_TotalLength256(t *testing.T) {
	parts := []string{
		strings.Repeat("a", 63),
		strings.Repeat("b", 63),
		strings.Repeat("c", 63),
		strings.Repeat("d", 62),
	}
	out, err := ens.DNSEncode(strings.Join(parts, "."))
	require.NoError(t, err, "ENSIP-10 explicitly removes the RFC 1035 255-octet cap")
	assert.Len(t, out, 256)
}

// Far beyond RFC 1035's cap — ENSIP-10 still requires this to succeed.
func TestDNSEncode_TotalLengthFarAbove255(t *testing.T) {
	// 10 labels of 63 bytes each = 630 data + 10 length + 1 null = 641 octets.
	parts := make([]string, 10)
	for i := range parts {
		parts[i] = strings.Repeat("a", 63)
	}
	out, err := ens.DNSEncode(strings.Join(parts, "."))
	require.NoError(t, err)
	assert.Len(t, out, 641)
}

// ---------------------------------------------------------------------------
// Non-ASCII / IDN labels.
//
// Normalize() uses idna.ToUnicode + MapForLookup, so the returned form is
// Unicode (NOT Punycode). DNSEncode then takes UTF-8 bytes verbatim.
// "münchen" UTF-8 = m(0x6D) ü(0xC3 0xBC) n(0x6E) c(0x63) h(0x68) e(0x65) n(0x6E)
// = 8 bytes.
// ---------------------------------------------------------------------------

func TestDNSEncode_NonASCII(t *testing.T) {
	out, err := ens.DNSEncode("münchen.eth")
	require.NoError(t, err)
	expected := []byte{
		8, 'm', 0xC3, 0xBC, 'n', 'c', 'h', 'e', 'n',
		3, 'e', 't', 'h',
		0,
	}
	assert.Equal(t, expected, out)
}

// IDNA MapForLookup should case-fold uppercase non-ASCII to lowercase.
func TestDNSEncode_NonASCIIUppercase(t *testing.T) {
	out, err := ens.DNSEncode("MÜNCHEN.eth")
	require.NoError(t, err)
	expected := []byte{
		8, 'm', 0xC3, 0xBC, 'n', 'c', 'h', 'e', 'n',
		3, 'e', 't', 'h',
		0,
	}
	assert.Equal(t, expected, out)
}
