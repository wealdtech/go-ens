// Copyright 2026 Weald Technology Trading.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0

package ens

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

// These tests lock in the ENSIP-15 behaviours that distinguish the
// adraffy/go-ens-normalize implementation from the pre-ENSIP-15 IDNA-only
// normaliser previously used in this package. They are intentionally narrow:
// the upstream library has its own 5000+ vector suite against the canonical
// ENSIP-15 validation tests — replicating those here would duplicate work
// without adding value. These tests instead exercise the wiring (our
// public Normalize function, error surface, and the categories of input
// the report's R15 recommendation called out).

func TestNormalize_AcceptsCanonicalInputs(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{"already-normalised ASCII", "vitalik.eth", "vitalik.eth"},
		{"case-folds ASCII", "VITALIK.ETH", "vitalik.eth"},
		{"mixed case", "ViTaLiK.eTh", "vitalik.eth"},
		{"non-ASCII passes through", "öbb.eth", "öbb.eth"},
		{"non-ASCII case-folds", "ÖBB.eth", "öbb.eth"},
		{"underscore at label start", "_foo.eth", "_foo.eth"},
		{"emoji label", "💷pound.eth", "💷pound.eth"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := Normalize(c.in)
			require.NoError(t, err)
			require.Equal(t, c.out, got)
		})
	}
}

func TestNormalize_RejectsENSIP15Violations(t *testing.T) {
	cases := []struct {
		name string
		in   string
	}{
		{"empty label / leading dot", ".eth"},
		{"empty label in middle", "foo..eth"},
		{"empty label / trailing dot", "foo."},
		{"bare dot", "."},
		{"XX-- label extension", "te--st.eth"},
		{"underscore mid-label", "a_b.eth"},
		{"whole-script confusable", "apple.дррӏе.аррӏе.aррӏе"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := Normalize(c.in)
			require.Error(t, err, "expected ENSIP-15 to reject %q", c.in)
		})
	}
}

// ENSIP-15 differs from IDNA on Arabic-Indic digits: IDNA mapped them to
// ASCII digits, ENSIP-15 maps Extended Arabic-Indic (U+06F0-U+06F9) to
// Arabic-Indic (U+0660-U+0669) but never folds to ASCII. The resulting
// NameHash therefore differs from a pre-ENSIP-15 IDNA hash. This test pins
// the new behaviour so future library bumps don't silently regress.
func TestNormalize_ArabicIndicDigitsNotFoldedToASCII(t *testing.T) {
	got, err := Normalize("۰۱۲۳۷۸۹.eth")
	require.NoError(t, err)
	require.NotEqual(t, "0123789.eth", got, "digits must NOT be ASCII-folded")
	require.Equal(t, "٠١٢٣٧٨٩.eth", got)
}

// Sanity check that Normalize errors are still wrapped with the
// "failed to ensip-15 normalize" prefix so the existing error-string
// behaviour at the API boundary is stable.
func TestNormalize_ErrorIsWrapped(t *testing.T) {
	_, err := Normalize(".eth")
	require.Error(t, err)
	require.Contains(t, err.Error(), "ensip-15")
	// Underlying ensip15 error is preserved via errors.Unwrap.
	require.NotNil(t, errors.Unwrap(err))
}
