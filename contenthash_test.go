// Copyright 2019-201 Weald Technology Trading
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
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func _hexStr(input string) []byte {
	res, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}
	return res
}

func TestContenthash(t *testing.T) {
	tests := []struct {
		name string
		repr string
		bin  []byte
		res  string
		err  error
	}{
		{
			name: "Nil",
			repr: "",
			err:  errors.New("no content hash"),
		},
		{
			name: "CodecMissing",
			repr: "://QmayQq2DWCkY3d4x3xKh4suohuRPEXe2fBqMBam5xtDj3t",
			err:  errors.New("codec missing"),
		},
		{
			name: "CodecMissing2",
			repr: "//QmayQq2DWCkY3d4x3xKh4suohuRPEXe2fBqMBam5xtDj3t",
			err:  errors.New("codec missing"),
		},
		{
			name: "Bad",
			repr: "/bad/QmayQq2DWCkY3d4x3xKh4suohuRPEXe2fBqMBam5xtDj3t",
			err:  errors.New("unknown codec bad"),
		},
		{
			name: "MissingData",
			repr: "/ipfs/",
			err:  errors.New("data missing"),
		},
		{
			name: "BadHash",
			repr: "/ipfs/invalid",
			err:  errors.New("invalid IPFS data: selected encoding not supported"),
		},
		{
			name: "IPFSCIDv0",
			repr: "/ipfs/QmRAQB6YaCyidP37UdDnjFY5vQuiBrcqdyoW1CuDgwxkD4",
			bin:  _hexStr("e3010170122029f2d17be6139079dc48696d1f582a8530eb9805b561eda517e22a892c7e3f1f"),
			res:  "/ipfs/k2jmtxseqz46solsx2rmxavgbzp6ij1t1kiq1or8a00c2g9bx1for0gv",
		},
		{
			name: "IPFSCIDv1",
			repr: "/ipfs/bafybeibj6lixxzqtsb45ysdjnupvqkufgdvzqbnvmhw2kf7cfkesy7r7d4",
			bin:  _hexStr("e3010170122029f2d17be6139079dc48696d1f582a8530eb9805b561eda517e22a892c7e3f1f"),
			res:  "/ipfs/k2jmtxseqz46solsx2rmxavgbzp6ij1t1kiq1or8a00c2g9bx1for0gv",
		},
		{
			name: "Onion",
			repr: "/onion/zqktlwi4fecvo6ri",
			bin:  _hexStr("bc037a716b746c776934666563766f367269"),
			res:  "onion://zqktlwi4fecvo6ri",
		},
		{
			name: "Onion3",
			repr: "/onion3/p53lf57qovyuvwsc6xnrppyply3vtqm7l6pcobkmyqsiofyeznfu5uqd",
			bin:  _hexStr("bd037035336c663537716f7679757677736336786e72707079706c79337674716d376c3670636f626b6d797173696f6679657a6e667535757164"),
			res:  "onion3://p53lf57qovyuvwsc6xnrppyply3vtqm7l6pcobkmyqsiofyeznfu5uqd",
		},
		{
			name: "IPFS",
			repr: "ipfs://QmRAQB6YaCyidP37UdDnjFY5vQuiBrcqdyoW1CuDgwxkD4",
			bin:  _hexStr("e3010170122029f2d17be6139079dc48696d1f582a8530eb9805b561eda517e22a892c7e3f1f"),
			res:  "/ipfs/k2jmtxseqz46solsx2rmxavgbzp6ij1t1kiq1or8a00c2g9bx1for0gv",
		},
		{
			name: "IPNSIdentity",
			repr: "/ipns/k51qzi5uqu5djwbl0zcd4g9onue26a8nq97c0m9wp6kir1gibuyjxpkqpoxwag",
			bin:  _hexStr("e5010172002408011220950b8f62b925ecc50247cc8de1084b43f854fbc452894d9a1a97d7f27d0addb8"),
			res:  "/ipns/k51qzi5uqu5djwbl0zcd4g9onue26a8nq97c0m9wp6kir1gibuyjxpkqpoxwag",
		},
		{
			name: "IPNSEIP1577",
			repr: "ipns://k51qzi5uqu5djwbl0zcd4g9onue26a8nq97c0m9wp6kir1gibuyjxpkqpoxwag",
			bin:  _hexStr("e5010172002408011220950b8f62b925ecc50247cc8de1084b43f854fbc452894d9a1a97d7f27d0addb8"),
			res:  "/ipns/k51qzi5uqu5djwbl0zcd4g9onue26a8nq97c0m9wp6kir1gibuyjxpkqpoxwag",
		},
		{
			name: "SwarmBad",
			repr: "/swarm/invalid",
			err:  errors.New("invalid hex: encoding/hex: invalid byte: U+0069 'i'"),
		},
		{
			name: "Swarm",
			repr: "/swarm/d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162",
			bin:  _hexStr("e40101fa011b20d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162"),
			res:  "bzz://d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162",
		},
		{
			name: "bzz",
			repr: "bzz://d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162",
			bin:  []byte{0xe4, 0x01, 0x01, 0xfa, 0x01, 0x1b, 0x20, 0xd1, 0xde, 0x99, 0x94, 0xb4, 0xd0, 0x39, 0xf6, 0x54, 0x8d, 0x19, 0x1e, 0xb2, 0x67, 0x86, 0x76, 0x9f, 0x58, 0x08, 0x09, 0x25, 0x6b, 0x46, 0x85, 0xef, 0x31, 0x68, 0x05, 0x26, 0x5e, 0xa1, 0x62},
			res:  "bzz://d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162",
		},
		{
			name: "Skynet",
			repr: "sia://CABAB_1Dt0FJsxqsu_J4TodNCbCGvtFf1Uys_3EgzOlTcg",
			bin:  _hexStr("90b2c60508004007fd43b74149b31aacbbf2784e874d09b086bed15fd54cacff7120cce95372"),
			res:  "sia://CABAB_1Dt0FJsxqsu_J4TodNCbCGvtFf1Uys_3EgzOlTcg",
		},
		{
			name: "SkynetBase32",
			repr: "sia://100401vt8erk2idj3ambnsjo9q3kq2dggqvd2nul9imfus90pjkl6sg",
			bin:  _hexStr("90b2c60508004007fd43b74149b31aacbbf2784e874d09b086bed15fd54cacff7120cce95372"),
			res:  "sia://CABAB_1Dt0FJsxqsu_J4TodNCbCGvtFf1Uys_3EgzOlTcg",
		},
		{
			name: "SkynetInvalidSize",
			repr: "sia://AAAACABAB_1Dt0FJsxqsu_J4TodNCbCGvtFf1Uys_3EgzOlTcg",
			err:  errors.New("skylinks should be either 46 or 55 characters, depending on whether it is base64 or base32 encoded"),
		},
		{
			name: "SkynetInvalidEncoding",
			repr: "sia://%CABAB_1Dt0FJsxqsu_J4TodNCbCGvtFf1Uys_3EgzOlTc",
			err:  errors.New("skylink not correctly encoded"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bin, err := StringToContenthash(test.repr)
			if test.err != nil {
				require.NotNil(t, err, "missing error")
				assert.Equal(t, test.err.Error(), err.Error(), "incorrect error")
			} else {
				require.Nil(t, err, "unexpected error creating content hash")
				require.Equal(t, test.bin, bin, "incorrect coding")
				res, err := ContenthashToString(bin)
				require.Nil(t, err, "unexpected error decoding content hash")
				assert.Equal(t, test.res, res, "incorrect decoding")
			}
		})
	}
}
