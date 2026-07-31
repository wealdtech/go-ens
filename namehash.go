// Copyright 2017 - 2023 Weald Technology Trading.
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
	"strings"

	"github.com/adraffy/go-ens-normalize/ensip15"
	"github.com/pkg/errors"
	"golang.org/x/net/idna"

	"golang.org/x/crypto/sha3"
)

// pStrict is the legacy IDNA "strict" profile, used only by
// NormaliseDomainStrict for DNS-strict checks (no underscores, etc.). It is
// NOT an ENSIP-15 normalizer — see Normalize / NormaliseDomain for that.
var pStrict = idna.New(idna.MapForLookup(), idna.ValidateLabels(false), idna.CheckHyphens(false), idna.StrictDomainName(true), idna.Transitional(false))

// Normalize normalizes a name according to ENSIP-15. Empty labels, disallowed
// codepoints, mixed-script confusables, and other ENSIP-15 violations return
// an error. The implementation delegates to github.com/adraffy/go-ens-normalize,
// the canonical Go port (Unicode 17.0.0, passes the full ENSIP-15 validation
// suite).
func Normalize(input string) (string, error) {
	output, err := ensip15.Shared().Normalize(input)
	if err != nil {
		return "", errors.Wrap(err, "failed to ensip-15 normalize")
	}
	return output, nil
}

// LabelHash generates a simple hash for a piece of a name.
func LabelHash(label string) ([32]byte, error) {
	var hash [32]byte

	normalizedLabel, err := Normalize(label)
	if err != nil {
		return [32]byte{}, err
	}

	sha := sha3.NewLegacyKeccak256()
	if _, err = sha.Write([]byte(normalizedLabel)); err != nil {
		return [32]byte{}, errors.Wrap(err, "failed to write hash")
	}
	sha.Sum(hash[:0])

	return hash, nil
}

// NameHash generates a hash from a name that can be used to
// look up the name in ENS.
func NameHash(name string) ([32]byte, error) {
	var hash [32]byte

	if name == "" {
		return hash, nil
	}

	normalizedName, err := Normalize(name)
	if err != nil {
		return [32]byte{}, err
	}
	parts := strings.Split(normalizedName, ".")
	for i := len(parts) - 1; i >= 0; i-- {
		if hash, err = nameHashPart(hash, parts[i]); err != nil {
			return [32]byte{}, err
		}
	}

	return hash, nil
}

func nameHashPart(currentHash [32]byte, name string) ([32]byte, error) {
	var hash [32]byte

	sha := sha3.NewLegacyKeccak256()
	if _, err := sha.Write(currentHash[:]); err != nil {
		return [32]byte{}, errors.Wrap(err, "failed to write hash")
	}
	nameSha := sha3.NewLegacyKeccak256()
	if _, err := nameSha.Write([]byte(name)); err != nil {
		return [32]byte{}, errors.Wrap(err, "failed to write hash")
	}
	nameHash := nameSha.Sum(nil)
	if _, err := sha.Write(nameHash); err != nil {
		return [32]byte{}, errors.Wrap(err, "failed to write hash")
	}
	sha.Sum(hash[:0])

	return hash, nil
}
