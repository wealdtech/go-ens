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
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	multihash "github.com/multiformats/go-multihash"
	multicodec "github.com/wealdtech/go-multicodec"
)

// StringToContenthash turns EIP-1577 text format in to EIP-1577 binary format
func StringToContenthash(text string) ([]byte, error) {
	bits := strings.Split(text, "/")
	if len(bits) != 3 {
		return nil, fmt.Errorf("invalid content hash")
	}
	data := make([]byte, 0)
	switch bits[1] {
	case "ipfs":
		// Namespace
		ipfsNum, err := multicodec.ID("ipfs-ns")
		if err != nil {
			return nil, errors.New("failed to obtain IPFS namespace value")
		}
		buf := make([]byte, binary.MaxVarintLen64)
		size := binary.PutUvarint(buf, ipfsNum)
		data = append(data, buf[0:size]...)
		// CID
		size = binary.PutUvarint(buf, 1)
		data = append(data, buf[0:size]...)
		// Codec
		dagNum, err := multicodec.ID("dag-pb")
		if err != nil {
			return nil, errors.New("failed to obtain IPFS codec value")
		}
		size = binary.PutUvarint(buf, dagNum)
		data = append(data, buf[0:size]...)
		// Hash
		hash, err := multihash.FromB58String(bits[2])
		if err != nil {
			return nil, errors.New("failed to obtain IPFS hash")
		}
		data = append(data, []byte(hash)...)
	case "ipns":
		// Namespace
		ipfsNum, err := multicodec.ID("ipns-ns")
		if err != nil {
			return nil, errors.New("failed to obtain IPNS namespace value")
		}
		buf := make([]byte, binary.MaxVarintLen64)
		size := binary.PutUvarint(buf, ipfsNum)
		data = append(data, buf[0:size]...)
		// CID
		size = binary.PutUvarint(buf, 1)
		data = append(data, buf[0:size]...)
		// Codec
		dagNum, err := multicodec.ID("dag-pb")
		if err != nil {
			return nil, errors.New("failed to obtain IPNS codec value")
		}
		size = binary.PutUvarint(buf, dagNum)
		data = append(data, buf[0:size]...)
		// Assume it's a multihash to begin with
		hash, err := multihash.FromB58String(bits[2])
		if err == nil {
			data = append(data, []byte(hash)...)
		} else {
			// Wasn't a multihash; pass along as identity
			encoded, err := multihash.Encode([]byte(bits[2]), multihash.ID)
			if err != nil {
				return nil, err
			}
			data = append(data, encoded...)
		}
	case "swarm":
		// Namespace
		swarmNum, err := multicodec.ID("swarm-ns")
		if err != nil {
			return nil, errors.New("failed to obtain swarm namespace value")
		}
		buf := make([]byte, binary.MaxVarintLen64)
		size := binary.PutUvarint(buf, swarmNum)
		data = append(data, buf[0:size]...)
		// CID
		size = binary.PutUvarint(buf, 1)
		data = append(data, buf[0:size]...)
		// Codec
		manifestNum, err := multicodec.ID("swarm-manifest")
		if err != nil {
			return nil, errors.New("failed to obtain swarm manifest codec value")
		}
		size = binary.PutUvarint(buf, manifestNum)
		data = append(data, buf[0:size]...)
		// Hash
		bit, err := hex.DecodeString(bits[2])
		if err != nil {
			return nil, errors.New("failed to decode swarm content hash")
		}
		hash, err := multihash.Encode(bit, multihash.KECCAK_256)
		if err != nil {
			return nil, errors.New("failed to obtain swarm content hash")
		}
		data = append(data, []byte(hash)...)
	default:
		return nil, fmt.Errorf("unknown codec %s", bits[1])
	}
	return data, nil
}

// ContenthashToString turns EIP-1577 binary format in to EIP-1577 text format
func ContenthashToString(bytes []byte) (string, error) {
	data, codec, err := multicodec.RemoveCodec(bytes)
	if err != nil {
		return "", err
	}
	codecName, err := multicodec.Name(codec)
	if err != nil {
		return "", err
	}
	id, offset := binary.Uvarint(data)
	if id == 0 {
		return "", fmt.Errorf("unknown CID")
	}
	data, subCodec, err := multicodec.RemoveCodec(data[offset:])
	if err != nil {
		return "", err
	}
	_, err = multicodec.Name(subCodec)
	if err != nil {
		return "", err
	}

	mHash, err := multihash.Cast(data)
	if err != nil {
		return "", err
	}
	decodedMHash, err := multihash.Decode(data)
	if err != nil {
		return "", err
	}
	switch codecName {
	case "ipfs-ns":
		return fmt.Sprintf("/ipfs/%s", mHash.B58String()), nil
	case "ipns-ns":
		switch decodedMHash.Code {
		case multihash.ID:
			return fmt.Sprintf("/ipns/%s", string(decodedMHash.Digest)), nil
		default:
			return fmt.Sprintf("/ipns/%s", mHash.B58String()), nil
		}
	case "swarm-ns":
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("/swarm/%x", decodedMHash.Digest), nil
	default:
		return "", fmt.Errorf("unknown codec %s", codecName)
	}
}
