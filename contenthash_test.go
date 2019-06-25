// Copyright 2019 Weald Technology Trading
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
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContenthash(t *testing.T) {
	tests := []struct {
		repr string
		bin  []byte
		err  error
	}{
		{ // 0
			repr: "",
			err:  errors.New("invalid content hash"),
		},
		{ // 1
			repr: "/bad/",
			err:  errors.New("unknown codec bad"),
		},
		{ // 2
			repr: "/ipfs/",
			err:  errors.New("failed to obtain IPFS hash"),
		},
		{ // 3
			repr: "/ipfs/bad",
			err:  errors.New("failed to obtain IPFS hash"),
		},
		{ // 4
			repr: "/ipfs/QmayQq2DWCkY3d4x3xKh4suohuRPEXe2fBqMBam5xtDj3t",
			bin:  []byte{0xe3, 0x1, 0x1, 0x70, 0x12, 0x20, 0xbb, 0xb7, 0x9, 0xbb, 0x1c, 0x1f, 0xd9, 0x10, 0x7d, 0x81, 0x8, 0x3f, 0x36, 0xe3, 0x6d, 0xd9, 0x71, 0xe4, 0xfb, 0xfe, 0x45, 0x77, 0xa8, 0x6a, 0x53, 0xd1, 0x91, 0xe, 0x3b, 0xcd, 0xbd, 0xff},
		},
		{ // 5
			repr: "/ipns/QmQ4QZh8nrsczdUEwTyfBope4THUhqxqc1fx6qYhhzZQei",
			bin:  []byte{0xe5, 0x1, 0x1, 0x70, 0x12, 0x20, 0x19, 0x8e, 0x16, 0x47, 0x97, 0x81, 0xa5, 0x73, 0xfc, 0x91, 0x42, 0xc1, 0xf1, 0x48, 0x7d, 0x67, 0xd0, 0x2c, 0xdc, 0x63, 0xd1, 0x19, 0x97, 0xb1, 0xb9, 0x9f, 0xf8, 0x85, 0x4e, 0x65, 0x22, 0xd7},
		},
		{ // 6
			repr: "/ipns/wealdtech.eth",
			bin:  []byte{0xe5, 0x1, 0x1, 0x70, 0x0, 0x0d, 0x77, 0x65, 0x61, 0x6c, 0x64, 0x74, 0x65, 0x63, 0x68, 0x2e, 0x65, 0x74, 0x68},
		},
		{ // 7
			repr: "/swarm",
			err:  errors.New("invalid content hash"),
		},
		{ // 8
			repr: "/swarm/bad",
			err:  errors.New("failed to decode swarm content hash"),
		},
		{ // 9
			repr: "/swarm/d1de9994b4d039f6548d191eb26786769f580809256b4685ef316805265ea162",
			bin:  []byte{0xe4, 0x1, 0x1, 0xfa, 0x1, 0x1b, 0x20, 0xd1, 0xde, 0x99, 0x94, 0xb4, 0xd0, 0x39, 0xf6, 0x54, 0x8d, 0x19, 0x1e, 0xb2, 0x67, 0x86, 0x76, 0x9f, 0x58, 0x8, 0x9, 0x25, 0x6b, 0x46, 0x85, 0xef, 0x31, 0x68, 0x5, 0x26, 0x5e, 0xa1, 0x62},
		},
	}

	for i, test := range tests {
		bin, err := StringToContenthash(test.repr)
		if test.err != nil {
			assert.Equal(t, test.err.Error(), err.Error(), fmt.Sprintf("incorrect error at test %d", i))
		} else {
			assert.Nil(t, err, fmt.Sprintf("unexpected error creating content hash at test %d", i))
			assert.Equal(t, test.bin, bin, fmt.Sprintf("incorrect coding at test %d", i))
			repr, err := ContenthashToString(bin)
			assert.Nil(t, err, fmt.Sprintf("unexpected error decoding content hash at test %d", i))
			assert.Equal(t, test.repr, repr, fmt.Sprintf("incorrect decoding at test %d", i))
		}
	}
}
