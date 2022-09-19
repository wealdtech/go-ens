package ens

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeriveTokenIdFromENSName(t *testing.T) {
	exepected := "79233663829379634837589865448569342784712482819484549289560981379859480642508"
	tokenID, err := DeriveTokenIdFromENSDomain("vitalik.eth")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, exepected, tokenID)
}
