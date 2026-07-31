// Copyright 2017 Weald Technology Trading
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ens

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNameHash(t *testing.T) {
	// Pre-ENSIP-15 fixtures that produced valid hashes for inputs containing
	// empty labels (".eth", "foo..eth") are now expected to error — ENSIP-15
	// rejects empty labels rather than silently round-tripping them.
	tests := []struct {
		input   string
		output  string
		wantErr bool
	}{
		{input: "", output: "0000000000000000000000000000000000000000000000000000000000000000"},
		{input: "eth", output: "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"},
		{input: "Eth", output: "93cdeb708b7545dc668eb9280176169d1c33cfd8ed6f04690a0bcc88a93fc4ae"},
		{input: ".eth", wantErr: true},
		{input: "resolver.eth", output: "fdd5d5de6dd63db72bbc2d487944ba13bf775b50a80805fe6fcaba9b0fba88f5"},
		{input: "foo.eth", output: "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"},
		{input: "Foo.eth", output: "de9b09fd7c5f901e23a3f19fecc54828e9c848539801e86591bd9801b019f84f"},
		{input: "foo..eth", wantErr: true},
		{input: "bar.foo.eth", output: "275ae88e7263cdce5ab6cf296cdd6253f5e385353fe39cfff2dd4a2b14551cf3"},
		{input: "Bar.foo.eth", output: "275ae88e7263cdce5ab6cf296cdd6253f5e385353fe39cfff2dd4a2b14551cf3"},
		{input: "addr.reverse", output: "91d1777781884d03a6757a803996e38de2a42967fb37eeaca72729271025a9e2"},
		{input: "🚣‍♂️.eth", output: "37f681401d88093779de0c910da1f9437759c2181f61b5dbc9e3ef0d9f192513"},
		{input: "vitalik.eth", output: "ee6c4522aab0003e8d14cd40a6af439055fd2577951148c14b6cea9a53475835"},
		{input: "nIcK.eTh", output: "05a67c0ee82964c4f7394cdd47fee7f4d9503a23c09c38341779ea012afe6e00"},
		{input: "brantly.cash", output: "3cb8d24c572869f7ec5ec5e882243b4790de146ac11267f58fd856d290b593db"},
		{input: "🏴‍☠.art", output: "4696ec532b8c8bb20dcfa6ae6da710422611ea03fc77aebc0284d0f12c93f768"},
		{input: "öbb.eth", output: "2774094517aaa884fc7183cf681150529b9acc7c9e7b71cf3a470200135a20b4"},
		{input: "Öbb.eth", output: "2774094517aaa884fc7183cf681150529b9acc7c9e7b71cf3a470200135a20b4"},
		{input: "💷pound.eth", output: "2df37f43de8b5ab870370770e78e4503ea1199e46c16db7a4a9320732dbb42bc"},
		{input: "firstwrappedname.eth", output: "c44eec7fb870ae46d4ef4392d33fbbbdc164e7817a86289a1fe30e5f4d98ae85"},
		{input: "A.™️.Ю", output: "f48b2941db0e23087a8922f81c0b7dd1b83fa06657080fd28986f9a19752a093"},
		{input: "Ⅷ.eth", output: "a468899f828f5d80c85c0e538018ba6516713a06aa6d92e8482ee91886630ae9"},
		{input: "⨌.eth", output: "f0b6090714f4ab523b1c574077018fe49dc934f9334d092e412b3802898b2cee"},
		// `-test`, `test-`, and `t-e--s---t` are all valid labels — only the
		// `XX--` extension prefix at positions 3-4 is invalid (see "te--st.eth"
		// below).
		{input: "-test.test-.t-e--s---t", output: "0d26fe02b1eb3fdcd7db77a432c60b20d2d8c4ef9012ef06ad3ae7bf12adb225"},
		// ENSIP-15 forbids the `XX--` label extension (RFC 5891-style guard
		// against ambiguity with Punycode).
		{input: "te--st.eth", wantErr: true},
		{input: "__ab.eth", output: "b3edf1f0b8dc6168b65d034c76b5670813514ac8867115e63b51eb496f3791a0"},
		// ENSIP-15 allows underscores only at the start of a label.
		{input: "a_b.eth", wantErr: true},
		{input: "ß.eth", output: "86262f2349fb651f652e87141e6db5856cd6e40f92d0b9dc486aaf8f9c509cff"},
		{input: "💩💩💩.eth", output: "a74feb0e5fa5606d3e650275e3bb3873b006a10d558389d3ce2abbe681fcfc8e"},
		{input: "آٍَ.ऩँं", output: "a1e3d72c39aec7ebcf62906994da7adc76b8a3ffa27e52b838abb24569610146"},
		{input: "פעילותהבינאום.eth", output: "7900040d82be5a569289b0f7b30266d796e1ded13293557ed061b39367357a62"},
		{input: "ஶ்ரீ.ஸ்ரீ", output: "ef5f28b8adff4b9b86b67c40792ef314ecd92a2123a72185b1411835cce5c110"},
		{input: "🇸🇦سلمان.eth", output: "3484f75c4dbf69503310126dbfc02eec224eab46851de2288092d58785670808"},
		{input: "0a〇.黑a8", output: "d7bc21f0cb86cafe27ebafe8e788b1c63b736467c65611110575f385a4d01e1f"},
		// ENSIP-15 explicitly rejects whole-script confusables; this is the
		// canonical "Cyrillic-mimics-Latin" example.
		{input: "apple.дррӏе.аррӏе.aррӏе", wantErr: true},
		// Hash differs from the pre-ENSIP-15 value because the old IDNA profile
		// mapped Arabic-Indic digits to ASCII digits; ENSIP-15 preserves them.
		{input: "۰۱۲۳۷۸۹.eth", output: "ff4b07401acb2d0f048c20f84a6ab57f6d66a0a16d81d13bbba87e8cfacd6594"},
		{input: "𓀀𓀁𓀂.eth", output: "36f75325f4013011b96014606224f03c77366d89caf480a1060d7453edaf5c98"},
		{input: "𓆏➡🐸️.eth", output: "f2663423c7aadf794d6f84addcfee1f877458d86c20740954482094a20985228"},
		{input: "ⒶⒷⒸⒹⒺⒻⒼⒽⒾⒿⓀⓁⓂⓃⓄⓅⓆⓇⓈⓉⓊⓋⓌⓍⓎⓏ.eth", output: "254a068876bdac64ad35425816bf444fbe9a26ec19287f47e663f928727ff69c"},
		{input: "trish🫧.eth", output: "9a6aa3aaa97ed3b694265d4b0e776b2e8fad34b8869199b1a062e3cc0ee0d41d"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := NameHash(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			if got := hex.EncodeToString(result[:]); got != tt.output {
				t.Errorf("%v => %v (expected %v)", tt.input, got, tt.output)
			}
		})
	}
}

func TestLabelHash(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    error
	}{
		{"", "c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470", nil},
		{"eth", "4f5b812789fc606be1b3b16908db13fc7a9adf7ca72641f84d75b47069d3d7f0", nil},
		{"foo", "41b1a0649752af1b28b3dc29a1556eee781e4a4c3a1f7f53f90fa834de098c4d", nil},
	}

	for _, tt := range tests {
		output, err := LabelHash(tt.input)
		if tt.err == nil {
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if tt.output != hex.EncodeToString(output[:]) {
				t.Errorf("Failure: %v => %v (expected %v)\n", tt.input, hex.EncodeToString(output[:]), tt.output)
			}
		} else {
			if err == nil {
				t.Fatalf("missing expected error")
			}
			if tt.err.Error() != err.Error() {
				t.Errorf("unexpected error value %v", err)
			}
		}
	}
}
